const express = require("express");
const { createProxyMiddleware } = require("http-proxy-middleware");
const addCorrelationId = require("./middleware/add-correlation-id");
const addApiLog = require("./middleware/add-api-log");
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

//Add Correlation Id
app.use(addCorrelationId);

//Add Api Log
app.use(addApiLog);

const setRequestHeader = (proxyReq, req) => {
    proxyReq.setHeader("correlation-id", getCorrelationId(req));
    if (!isPublicUrl(req.originalUrl)) {
        proxyReq.setHeader("user-id", getUserId(req));
        proxyReq.setHeader("user-role", getUserRole(req));
    }
    proxyReq.removeHeader("authorization");
};

//Proxy
app.use(
    config.userService.userPath,
    authenticateToken,
    createProxyMiddleware({
        target: config.userService.target,
        changeOrigin: true,
        proxyTimeout: 1000,
        onError(err, req, res) {
            logger.error(
                `CorrelationId: ${getCorrelationId(req)},\nstatus: ${
                    apiStatus.INTERNAL_SERVER_ERROR.status
                },\nmessage: ${err.message},\nstack: ${err.stack}`
            );
            res.writeHead(500, { "Content-Type": "text/plain" });
            res.end("Something went wrong");
        },
        onProxyReq: (proxyReq, req, res) => {
            //req is the (original) incoming request, while proxyReq is the request that will be issued to the target server.

            setRequestHeader(proxyReq, req);
        },
        onProxyRes: (proxyRes, req, res) => {},
    })
);

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

//Routes
router(app);

//Not Found
app.use(notFoundHandler);

module.exports = app;
