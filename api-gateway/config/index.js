require("dotenv").config();
const env = process.env.NODE_ENV || "dev";

/**
 * @type {{
 * env: string,
 * appName: string,
 * port: number,
 * accessToken: string,
 * userService: {userPath: string, authPath: string, target: string}
 * }}
 * */
let config;
switch (env) {
    case "prod":
        config = require("./prod.js");
        break;
    default:
        config = require("./dev.js");
        break;
}
module.exports = config;
