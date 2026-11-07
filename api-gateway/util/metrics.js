const client = require("../handler/prom-client");

const gatewayRequestDurationHistogram = new client.Histogram({
    name: "gateway_http_request_duration_sec",
    help: "API Gateway request duration in sec",
    labelNames: ["method", "route", "status_code"],
    buckets: [0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2, 5, 10],
});

module.exports = { gatewayRequestDurationHistogram };
