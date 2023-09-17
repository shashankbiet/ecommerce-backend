const apiUrl = require("../app/shared/constants/api-url");

const config = {};
config.env = "dev";
config.appName = "user-service";
config.port = 3002;
config.mongodbUrl = "mongodb://localhost:27017/userService";
config.adminVerificationToken = "1234567890";
config.disableUrlLogging = [apiUrl.health, apiUrl.metrics];

module.exports = config;
