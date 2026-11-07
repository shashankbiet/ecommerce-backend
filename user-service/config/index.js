const fs = require("fs");
const apiUrl = require("../app/shared/constants/api-url");

// Secrets arrive as NAME_FILE pointers (Compose secrets or ../secrets/*); load each into NAME.
for (const [fileVar, file] of Object.entries(process.env)) {
    if (!fileVar.endsWith("_FILE")) continue;
    const name = fileVar.slice(0, -"_FILE".length);
    if (!process.env[name])
        process.env[name] = fs.readFileSync(file, "utf8").trim();
}
const env = process.env.NODE_ENV === "prod" ? "prod" : "dev";

/**
 * @type {{
 * env: string,
 * appName: string,
 * port: number,
 * mongodbUrl: string,
 * adminVerificationToken: string
 * disableUrlLogging: []
 * }}
 * */
module.exports = {
    env,
    appName: "user-service",
    port: Number(process.env.USER_SERVICE_PORT) || 3002,
    mongodbUrl: `mongodb://${encodeURIComponent(process.env.MONGODB_ROOT_USERNAME)}:${encodeURIComponent(
        process.env.MONGODB_ROOT_PASSWORD
    )}@${process.env.MONGODB_HOST || "localhost"}:27017/${process.env.USER_SERVICE_DATABASE}?authSource=admin`,
    adminVerificationToken: process.env.USER_SERVICE_ADMIN_VERIFICATION_TOKEN,
    disableUrlLogging: [apiUrl.health, apiUrl.metrics],
};
