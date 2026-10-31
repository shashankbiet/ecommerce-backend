const jwt = require("jsonwebtoken");
const config = require("../../config");
const authService = {};

authService.createToken = (user) => {
    return jwt.sign(user, config.accessToken, {
        expiresIn: 86400, // expires in 24 hours
    });
};

module.exports = authService;
