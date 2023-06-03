const userRouter = require("../app/user/user-router");
const addressRouter = require("../app/address/address-router");
const authRouter = require("../app/auth/auth-router");
const healthCheck = require("../util/health-check");

const router = (app) => {
    app.get("/api/user/health", healthCheck);
    app.use("/api/user", userRouter);
    app.use("/api/user/address", addressRouter);
    app.use("/api/auth", authRouter);
};

module.exports = router;
