const userRouter = require("../app/user/user-router");
const addressRouter = require("../app/address/address-router");
const authRouter = require("../app/auth/auth-router");
const healthRouter = require("../util/health-router");

const router = (app) => {
    app.use("/health", healthRouter);
    app.use("/api/user", userRouter);
    app.use("/api/user/address", addressRouter);
    app.use("/api/auth", authRouter);
};

module.exports = router;
