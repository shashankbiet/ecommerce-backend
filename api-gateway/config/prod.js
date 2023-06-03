require("dotenv").config();
const config = {};
config.env = "prod";
config.appName = process.env.APP_NAME;
config.port = process.env.PORT;
config.accessToken = process.env.ACCESS_TOKEN_SECRET;
config.userService = {
    userPath: "/api/user",
    authPath: "/api/auth",
    target: process.env.USER_SERVICE_TARGET,
};

module.exports = config;
