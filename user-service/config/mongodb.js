const mongoose = require("mongoose");
require("dotenv").config();
let dbUrl = process.env.MONGODB_URL;

module.exports = async () => {
    await mongoose
        .connect(dbUrl)
        .then(() => console.log("MONGODB CONNECTED"))
        .catch((err) => console.log(err));
};
