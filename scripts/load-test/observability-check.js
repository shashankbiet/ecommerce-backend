#!/usr/bin/env node
/**
 * Checks that the monitoring stack is scraping and collecting data.
 * Requires the `monitoring` compose profile (Prometheus, Grafana, Promtail) plus Loki.
 *
 * Run from scripts/load-test:
 *   npm run observability-check
 *
 * Env overrides:
 *   PROMETHEUS_URL  default http://localhost:9090
 *   LOKI_URL        default http://localhost:3100
 *   GRAFANA_URL     default http://localhost:3000
 */
const PROMETHEUS_URL = process.env.PROMETHEUS_URL || "http://localhost:9090";
const LOKI_URL = process.env.LOKI_URL || "http://localhost:3100";
const GRAFANA_URL = process.env.GRAFANA_URL || "http://localhost:3000";
const TIMEOUT_MS = 5000;

const results = [];

async function getJson(url) {
    const res = await fetch(url, { signal: AbortSignal.timeout(TIMEOUT_MS) });
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    return res.json();
}

async function check(name, fn) {
    try {
        const detail = await fn();
        results.push({ check: name, status: "PASS", detail });
    } catch (err) {
        results.push({ check: name, status: "FAIL", detail: err.message });
    }
}

async function prometheusQuery(query) {
    const body = await getJson(
        `${PROMETHEUS_URL}/api/v1/query?query=${encodeURIComponent(query)}`
    );
    return body.data.result;
}

async function main() {
    await check("prometheus ready", async () => {
        const res = await fetch(`${PROMETHEUS_URL}/-/ready`, {
            signal: AbortSignal.timeout(TIMEOUT_MS),
        });
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        return "ready";
    });

    await check("prometheus scrape targets", async () => {
        const body = await getJson(`${PROMETHEUS_URL}/api/v1/targets`);
        const targets = body.data.activeTargets;
        const down = targets.filter((t) => t.health !== "up");
        for (const t of targets) {
            console.log(
                `  target ${t.labels.job} (${t.scrapeUrl}): ${t.health}${
                    t.lastError ? ` - ${t.lastError}` : ""
                }`
            );
        }
        if (down.length) {
            throw new Error(
                `${down.length}/${targets.length} targets down: ${down
                    .map((t) => t.labels.job)
                    .join(", ")}`
            );
        }
        return `${targets.length} targets up`;
    });

    await check("prometheus has series for each service job", async () => {
        const jobs = ["api-gateway", "inventory-service", "search-service"];
        const missing = [];
        for (const job of jobs) {
            const series = await prometheusQuery(`count({job="${job}"})`);
            if (!series.length) missing.push(job);
        }
        if (missing.length) throw new Error(`no series for: ${missing.join(", ")}`);
        return `series present for ${jobs.join(", ")}`;
    });

    await check("prometheus HTTP request rate (last 5m)", async () => {
        const series = await prometheusQuery(
            'sum by (job) (rate({__name__=~"inventory_service_http_requests_total|gateway_http_request_duration_sec_count"}[5m]))'
        );
        if (!series.length) throw new Error("no request-rate series yet; run load test first");
        return series
            .map((s) => `${s.metric.job}=${Number(s.value[1]).toFixed(2)}/s`)
            .join(", ");
    });

    await check("loki ready", async () => {
        const res = await fetch(`${LOKI_URL}/ready`, {
            signal: AbortSignal.timeout(TIMEOUT_MS),
        });
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        return "ready";
    });

    await check("loki has log labels", async () => {
        const body = await getJson(`${LOKI_URL}/loki/api/v1/labels`);
        if (!body.data || !body.data.length) {
            throw new Error("no labels; check promtail is shipping logs");
        }
        return body.data.join(", ");
    });

    await check("grafana health", async () => {
        const body = await getJson(`${GRAFANA_URL}/api/health`);
        if (body.database !== "ok") throw new Error(`database: ${body.database}`);
        return `version ${body.version}`;
    });

    console.table(results);
    const failed = results.filter((r) => r.status === "FAIL").length;
    if (failed) {
        console.error(`${failed} check(s) failed.`);
        process.exit(1);
    }
}

main().catch((err) => {
    console.error(err);
    process.exit(1);
});
