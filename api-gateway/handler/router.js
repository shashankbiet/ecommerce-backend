const healthCheck = require("../util/health-check");
const authRouter = require("../app/auth/auth-router");
const registerMetrics = require("../util/register-metrics");

const router = (app) => {
    app.get("/health", healthCheck);
    app.get("/metrics", registerMetrics);
    app.use("/auth", authRouter);
};

module.exports = router;
