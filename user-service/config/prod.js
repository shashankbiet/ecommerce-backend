const apiUrl = require("../app/shared/constants/api-url");

require("dotenv").config();
const config = {};
config.env = "prod";
config.appName = process.env.APP_NAME;
config.port = process.env.PORT;
config.mongodbUrl = process.env.MONGODB_URL;
config.adminVerificationToken = process.env.ADMIN_VERIFICATION_TOKEN;
config.disableUrlLogging = [apiUrl.health, apiUrl.metrics];

module.exports = config;
