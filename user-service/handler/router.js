const userRouter = require("../app/user/user-router");
const addressRouter = require("../app/address/address-router");
const authRouter = require("../app/auth/auth-router");
const apiUrl = require("../app/shared/constants/api-url");
const healthCheck = require("../util/health-check");
const registerMetrics = require("../util/register-metrics");

const router = (app) => {
    app.get(apiUrl.health, healthCheck);
    app.get(apiUrl.metrics, registerMetrics);
    app.use(apiUrl.user, userRouter);
    app.use(apiUrl.userAddress, addressRouter);
    app.use(apiUrl.auth, authRouter);
};

module.exports = router;
