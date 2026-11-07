const express = require("express");
const { createProxyMiddleware } = require("http-proxy-middleware");
const addCorrelationId = require("./middleware/add-correlation-id");
const { recordRequestDuration } = require("./middleware/api-metrics");
const authenticateToken = require("./middleware/authenticate-token");
const notFoundHandler = require("./middleware/not-found-handler");
const router = require("./handler/router");
const apiStatus = require("./app/shared/constants/api-status");
const config = require("./config");
const {
    getCorrelationId,
    getUserId,
    getUserRole,
} = require("./util/request-header");
const { isPublicUrl } = require("./util/request-url");
const logger = require("./handler/logger");
const app = express();

// Add a trace ID so this request can be correlated across services.
app.use(addCorrelationId);
app.use(recordRequestDuration);

// Forward authenticated user requests to the User Service.
const setRequestHeader = (proxyReq, req) => {
    proxyReq.setHeader("correlation-id", getCorrelationId(req));
    if (!isPublicUrl(req.originalUrl)) {
        proxyReq.setHeader("user-id", getUserId(req));
        proxyReq.setHeader("user-role", getUserRole(req));
    }
    // The downstream service receives trusted identity headers, not the bearer token.
    proxyReq.removeHeader("authorization");
};

// Authenticate user endpoints before proxying them to the User Service.
app.use(
    config.userService.userPath,
    authenticateToken,
    createProxyMiddleware({
        target: config.userService.target,
        changeOrigin: true,
        timeout: 10000,
        proxyTimeout: 10000,
        onError(err, req, res) {
            // Return a consistent gateway response when the upstream service fails.
            logger.error(
                `CorrelationId: ${getCorrelationId(req)},\nstatus: ${
                    apiStatus.INTERNAL_SERVER_ERROR.status
                },\nmessage: ${err.message},\nstack: ${err.stack}`
            );
            if (!res.headersSent) {
                res.writeHead(500, { "Content-Type": "application/json" });
                res.end(
                    JSON.stringify({
                        success: false,
                        message: "Gateway error: upstream service unavailable",
                    })
                );
            }
        },
        onProxyReq: (proxyReq, req, res) => {
            setRequestHeader(proxyReq, req);
        },
    })
);

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

//Routes
router(app);

//Not Found
app.use(notFoundHandler);

module.exports = app;
