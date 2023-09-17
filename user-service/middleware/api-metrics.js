const { restApiResponseTimeHistogram } = require("../util/metrics");
const responseTime = require("response-time");
const { disableUrlLogging } = require("../config");

let registerRestApiResponseTime = responseTime((req, res, time) => {
    if (!disableUrlLogging.includes(req.originalUrl)) {
        restApiResponseTimeHistogram
            .labels({
                method: req.method,
                route: req.originalUrl,
                status_code: res.statusCode,
            })
            .observe(time);
    }
});

module.exports = { registerRestApiResponseTime };
