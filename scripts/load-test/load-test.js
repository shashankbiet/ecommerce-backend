#!/usr/bin/env node
/**
 * Load test for the ecommerce backend (chatbot-service excluded).
 *
 * Run from scripts/load-test after `npm install`:
 *   USER_PASSWORD='<demo customer password>' npm run load-test
 *
 * Env overrides:
 *   DURATION_SECONDS        default 300
 *   WORKERS_PER_SCENARIO    default 3
 *   GATEWAY_URL             default http://localhost:3001
 *   INVENTORY_URL           default http://localhost:3003
 *   SEARCH_ADDRESS          default localhost:9004 (gRPC)
 *   USER_EMAIL              default demo.customer01@example.com
 *   USER_PASSWORD           required, used to obtain a bearer token via the gateway
 *
 * user-service (3002) and search-service metrics (9104) are not published to the host,
 * so they are only reached through the gateway and gRPC respectively.
 */
const path = require("path");
const { performance } = require("perf_hooks");
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

const config = {
    durationSeconds: Number(process.env.DURATION_SECONDS || 300),
    workers: Number(process.env.WORKERS_PER_SCENARIO || 3),
    gatewayUrl: process.env.GATEWAY_URL || "http://localhost:3001",
    inventoryUrl: process.env.INVENTORY_URL || "http://localhost:3003",
    searchAddress: process.env.SEARCH_ADDRESS || "localhost:9004",
    userEmail: process.env.USER_EMAIL || "demo.customer01@example.com",
    userPassword: process.env.USER_PASSWORD,
};

const REQUEST_TIMEOUT_MS = 10000;
const SEARCH_PROTO = path.resolve(
    __dirname,
    "../../search-service/proto/search/search.proto"
);

const searchDefinition = protoLoader.loadSync(SEARCH_PROTO, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});
const searchClient = new (grpc.loadPackageDefinition(
    searchDefinition
).searchpb.SearchService)(
    config.searchAddress,
    grpc.credentials.createInsecure()
);

async function httpCall(url, options = {}) {
    const res = await fetch(url, {
        ...options,
        signal: AbortSignal.timeout(REQUEST_TIMEOUT_MS),
    });
    // Drain the body so keep-alive connections are reused.
    await res.arrayBuffer();
    return res.status;
}

function searchProduct(request) {
    return new Promise((resolve, reject) => {
        searchClient.GetProduct(
            request,
            { deadline: Date.now() + REQUEST_TIMEOUT_MS },
            (err, response) => (err ? reject(err) : resolve(response))
        );
    });
}

async function login() {
    const res = await fetch(`${config.gatewayUrl}/auth/login`, {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
            email: config.userEmail,
            password: config.userPassword,
        }),
        signal: AbortSignal.timeout(REQUEST_TIMEOUT_MS),
    });
    if (!res.ok) {
        throw new Error(`login failed with HTTP ${res.status}`);
    }
    // The gateway returns the JWT as a raw string body.
    return (await res.text()).trim();
}

function buildScenarios(token) {
    const json = { "content-type": "application/json" };
    const auth = { authorization: `Bearer ${token}` };
    const loginBody = JSON.stringify({
        email: config.userEmail,
        password: config.userPassword,
    });

    return [
        {
            name: "gateway GET /health",
            run: () => httpCall(`${config.gatewayUrl}/health`),
        },
        {
            name: "gateway GET /metrics",
            run: () => httpCall(`${config.gatewayUrl}/metrics`),
        },
        {
            name: "gateway POST /auth/login",
            run: () =>
                httpCall(`${config.gatewayUrl}/auth/login`, {
                    method: "POST",
                    headers: json,
                    body: loginBody,
                }),
        },
        {
            name: "gateway GET /api/user",
            run: () =>
                httpCall(`${config.gatewayUrl}/api/user`, { headers: auth }),
        },
        {
            name: "gateway GET /api/user/address",
            run: () =>
                httpCall(`${config.gatewayUrl}/api/user/address`, {
                    headers: auth,
                }),
        },
        {
            name: "inventory GET /health",
            run: () => httpCall(`${config.inventoryUrl}/health`),
        },
        {
            name: "inventory GET /metrics",
            run: () => httpCall(`${config.inventoryUrl}/metrics`),
        },
        {
            name: "inventory GET /api/category",
            run: () => httpCall(`${config.inventoryUrl}/api/category`),
        },
        {
            name: "inventory GET /api/subcategory",
            run: () => httpCall(`${config.inventoryUrl}/api/subcategory`),
        },
        {
            name: "inventory GET /api/product",
            run: () => httpCall(`${config.inventoryUrl}/api/product`),
        },
        {
            name: "inventory GET /api/inventory",
            run: () => httpCall(`${config.inventoryUrl}/api/inventory`),
        },
        {
            name: "search gRPC GetProduct",
            run: async () => {
                await searchProduct({
                    keywords: "ab",
                    category: "Electronics",
                    sub_category: "",
                });
                // Map a successful gRPC call onto the same status scale as HTTP.
                return 200;
            },
        },
    ];
}

function newStats() {
    return { latencies: [], statusCounts: {}, errors: 0 };
}

function record(stats, latencyMs, status) {
    stats.latencies.push(latencyMs);
    const key = String(status);
    stats.statusCounts[key] = (stats.statusCounts[key] || 0) + 1;
    if (status >= 400) stats.errors++;
}

async function runWorker(scenario, deadline, stats) {
    while (Date.now() < deadline) {
        const start = performance.now();
        try {
            const status = await scenario.run();
            record(stats, performance.now() - start, status);
        } catch (err) {
            record(stats, performance.now() - start, "exception");
            stats.errors++;
        }
    }
}

function percentile(sorted, p) {
    if (sorted.length === 0) return 0;
    const index = Math.min(sorted.length - 1, Math.ceil((p / 100) * sorted.length) - 1);
    return sorted[index];
}

function summarize(scenario, stats, elapsedSeconds) {
    const sorted = [...stats.latencies].sort((a, b) => a - b);
    const total = sorted.length;
    return {
        scenario: scenario.name,
        requests: total,
        rps: Number((total / elapsedSeconds).toFixed(1)),
        errorRate: total ? `${((stats.errors / total) * 100).toFixed(2)}%` : "n/a",
        p50Ms: Number(percentile(sorted, 50).toFixed(1)),
        p95Ms: Number(percentile(sorted, 95).toFixed(1)),
        p99Ms: Number(percentile(sorted, 99).toFixed(1)),
        maxMs: Number((sorted[total - 1] || 0).toFixed(1)),
        statuses: JSON.stringify(stats.statusCounts),
    };
}

async function main() {
    if (!config.userPassword) {
        console.error("USER_PASSWORD is required to obtain a bearer token.");
        process.exit(1);
    }

    console.log("Logging in through the gateway...");
    const token = await login();
    const scenarios = buildScenarios(token);

    console.log("Warm-up check (one request per scenario)...");
    for (const scenario of scenarios) {
        try {
            const status = await scenario.run();
            console.log(`  ${scenario.name}: ${status}`);
        } catch (err) {
            console.error(`  ${scenario.name}: ${err.message}`);
            console.error("Aborting: fix the failing endpoint before load testing.");
            process.exit(1);
        }
    }

    const startedAt = Date.now();
    const deadline = startedAt + config.durationSeconds * 1000;
    console.log(
        `Running ${scenarios.length} scenarios x ${config.workers} workers for ${config.durationSeconds}s...`
    );

    const statsByScenario = scenarios.map(() => newStats());
    const workers = [];
    scenarios.forEach((scenario, index) => {
        for (let i = 0; i < config.workers; i++) {
            workers.push(runWorker(scenario, deadline, statsByScenario[index]));
        }
    });
    await Promise.all(workers);

    const elapsedSeconds = (Date.now() - startedAt) / 1000;
    const rows = scenarios.map((scenario, index) =>
        summarize(scenario, statsByScenario[index], elapsedSeconds)
    );
    console.log(`\nLoad test finished in ${elapsedSeconds.toFixed(1)}s`);
    console.table(rows);
}

main().catch((err) => {
    console.error(err);
    process.exit(1);
});
