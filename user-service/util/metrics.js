const client = require("../handler/prom-client");
const restApiResponseTimeHistogram = new client.Histogram({
    name: "rest_api_response_time_sec",
    help: "Rest API response time in sec",
    labelNames: ["method", "route", "status_code"],
    buckets: [0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2],
});

const dbResponseTimeHistogram = new client.Histogram({
    name: "db_response_time_sec",
    help: "DB response time in sec",
    labelNames: ["operation", "success"],
    buckets: [0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2],
});

const createUserReqCounter = new client.Counter({
    name: "create_user_req_api",
    help: "Ceate User Api request count",
});

const getUserReqCounter = new client.Counter({
    name: "get_user_req_api",
    help: "Get User Api request count",
});

module.exports = {
    restApiResponseTimeHistogram,
    dbResponseTimeHistogram,
    createUserReqCounter,
    getUserReqCounter,
};
