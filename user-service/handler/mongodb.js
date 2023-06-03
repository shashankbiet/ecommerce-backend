const mongoose = require("mongoose");
const config = require("../config");
let dbUrl = config.mongodbUrl;

module.exports = async () => {
    mongoose.set("strictQuery", false);
    await mongoose
        .connect(dbUrl)
        .then(() => console.log("MONGODB CONNECTED"))
        .catch((err) => console.log(err));
};
