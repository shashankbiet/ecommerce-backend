const { createLogger, transports, format } = require("winston");
const LokiTransport = require("winston-loki");
const logger = createLogger({
    transports: [
        new LokiTransport({
            host: "http://loki:3100",
            labels: { app: "api-gateway" },
            json: true,
            format: format.json(),
            replaceTimestamp: true,
            onConnectionError: (err) => console.error(err),
        }),
        new transports.Console({
            format: format.combine(format.simple(), format.colorize()),
        }),
    ],
});
module.exports = logger;
