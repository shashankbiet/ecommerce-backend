const client = require("../handler/prom-client");
const restApiResponseTimeHistogram = new client.Histogram({
    name: "rest_api_response_time_sec",
    help: "Rest API response time in sec",
    labelNames: ["method", "route", "status_code"],
    buckets: [1, 50, 100, 200, 400, 500, 800, 1000, 2000],
});

const dbResponseTimeHistogram = new client.Histogram({
    name: "db_response_time_sec",
    help: "DB response time in sec",
    labelNames: ["operation", "success"],
    buckets: [1, 50, 100, 200, 400, 500, 800, 1000, 2000],
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
