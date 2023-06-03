const config = {};
config.env = "dev";
config.appName = "api-gateway";
config.port = 3001;
config.accessToken = "9876543210";
config.userService = {
    userPath: "/api/user",
    authPath: "/api/auth",
    target: "http://localhost:3002",
};

module.exports = config;
