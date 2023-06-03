const jwt = require("jsonwebtoken");
const config = require("../../config");
const authService = {};

authService.createToken = async (user) => {
    const accessToken = await jwt.sign(user, config.accessToken, {
        expiresIn: 86400, // expires in 24 hours
    });
    return accessToken;
};

module.exports = authService;
