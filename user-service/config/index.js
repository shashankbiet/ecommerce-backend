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
 * postgresConfig: {
 * host: string,
 * port: number,
 * database: string,
 * user: string,
 * password: string,
 * maxClients: number
 * },
 * adminVerificationToken: string
 * disableUrlLogging: []
 * }}
 * */
module.exports = {
    env,
    appName: "user-service",
    port: Number(process.env.USER_SERVICE_PORT) || 3002,
    postgresConfig: {
        host: process.env.SQL_CONFIG_HOST || "localhost",
        port: Number(process.env.SQL_CONFIG_PORT) || 5432,
        database: process.env.SQL_CONFIG_DATABASE || "user_db",
        user: process.env.SQL_CONFIG_USERNAME || "postgres_user",
        password: process.env.SQL_CONFIG_PASSWORD || "",
        maxClients: Number(process.env.SQL_CONFIG_MAX_CLIENTS) || 10,
    },
    adminVerificationToken: process.env.USER_SERVICE_ADMIN_VERIFICATION_TOKEN,
    disableUrlLogging: [apiUrl.health, apiUrl.metrics],
};
