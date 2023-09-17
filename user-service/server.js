const app = require("./app");
const { port = 3002 } = require("./config");
const logger = require("./handler/logger");
//Setting up server
app.listen(port, () => logger.info(`User Service running on port ${port}`));
