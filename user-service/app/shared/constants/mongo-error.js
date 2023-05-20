const status = require("./api-status");

module.exports = {
    DUPLICATE_KEY: {
        code: 11000,
        status: status.CONFLICT.status,
        message: "Resource already exists.",
    },
};
