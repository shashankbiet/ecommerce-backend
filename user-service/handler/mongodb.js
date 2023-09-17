const mongoose = require("mongoose");
const { mongodbUrl } = require("../config");
const logger = require("./logger");

module.exports = async () => {
    mongoose.set("strictQuery", false);
    await mongoose
        .connect(mongodbUrl)
        .then(() => logger.info("MONGODB CONNECTED"))
        .catch((err) => logger.error(err));
};
