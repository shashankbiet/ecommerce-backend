const healthCheck = require("../util/health-check");
const authRouter = require("../app/auth/auth-router");

const router = (app) => {
    app.get("/health", healthCheck);
    app.use("/auth", authRouter);
};

module.exports = router;
