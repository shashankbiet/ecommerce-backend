const { gatewayRequestDurationHistogram } = require("../util/metrics");
const config = require("../config");

// Known mount points keep the route label low-cardinality; anything else is "unmatched".
const routePrefixes = [
    config.userService.userPath,
    "/auth",
    "/health",
    "/metrics",
];

const toRouteLabel = (url) => {
    const path = url.split("?")[0];
    return (
        routePrefixes.find(
            (prefix) => path === prefix || path.startsWith(`${prefix}/`)
        ) || "unmatched"
    );
};

const recordRequestDuration = (req, res, next) => {
    const start = process.hrtime.bigint();
    res.on("finish", () => {
        const seconds = Number(process.hrtime.bigint() - start) / 1e9;
        gatewayRequestDurationHistogram
            .labels({
                method: req.method,
                route: toRouteLabel(req.originalUrl),
                status_code: res.statusCode,
            })
            .observe(seconds);
    });
    next();
};

module.exports = { recordRequestDuration, toRouteLabel };
