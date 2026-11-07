const fs = require("fs");

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
 * accessToken: string,
 * userService: {userPath: string, authPath: string, target: string}
 * }}
 * */
module.exports = {
    env,
    appName: "api-gateway",
    port: Number(process.env.API_GATEWAY_PORT) || 3001,
    accessToken: process.env.API_GATEWAY_ACCESS_TOKEN_SECRET,
    userService: {
        userPath: process.env.USER_SERVICE_USER_PATH || "/api/user",
        authPath: process.env.USER_SERVICE_AUTH_PATH || "/api/auth",
        target: process.env.USER_SERVICE_TARGET || "http://localhost:3002",
    },
};
