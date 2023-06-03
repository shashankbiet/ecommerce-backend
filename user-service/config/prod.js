require("dotenv").config();
const config = {};
config.env = "prod";
config.appName = process.env.APP_NAME;
config.port = process.env.PORT;
config.mongodbUrl = process.env.MONGODB_URL;
config.adminVerificationToken = process.env.ADMIN_VERIFICATION_TOKEN;

module.exports = config;
