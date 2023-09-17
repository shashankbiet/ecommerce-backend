const client = require("prom-client"); //Metric Collection
const collectDefaultMetrics = client.collectDefaultMetrics;
collectDefaultMetrics({ register: client.register });

module.exports = client;
