const { restApiResponseTimeHistogram } = require("../util/metrics");
const responseTime = require("response-time");
const { disableUrlLogging } = require("../config");

let registerRestApiResponseTime = responseTime((req, res, time) => {
    if (!disableUrlLogging.includes(req.originalUrl)) {
        const route = req.route
            ? `${req.baseUrl}${req.route.path}`
            : "unmatched";
        restApiResponseTimeHistogram
            .labels({
                method: req.method,
                route,
                status_code: res.statusCode,
            })
            .observe(time / 1000);
    }
});

module.exports = { registerRestApiResponseTime };
