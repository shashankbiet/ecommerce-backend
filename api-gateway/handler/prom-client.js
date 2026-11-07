const client = require("@prometheus-io/client"); //Metric Collection
const collectDefaultMetrics = client.collectDefaultMetrics;
collectDefaultMetrics({ register: client.register });

module.exports = client;
