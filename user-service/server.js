const app = require("./app");
const { port = 3002 } = require("./config");
const logger = require("./handler/logger");
const { connect } = require("./handler/postgres");

const startServer = async () => {
    await connect();
    app.listen(port, () => logger.info(`User Service running on port ${port}`));
};

startServer().catch((err) => {
    logger.error(err);
    process.exit(1);
});
