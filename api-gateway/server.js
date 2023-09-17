const logger = require("./handler/logger");
const app = require("./app");
const config = require("./config");

const port = config.port || 3001;
app.listen(port, () => logger.info(`API Gateway running on port ${port}`));
