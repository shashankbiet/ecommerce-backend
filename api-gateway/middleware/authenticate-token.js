const jwt = require("jsonwebtoken");
const config = require("../config");
const apiStatus = require("../app/shared/constants/api-status");
const { isPublicUrl } = require("../util/request-url");

const authenticateToken = (req, res, next) => {
    if (isPublicUrl(req.originalUrl)) {
        return next();
    }
    let token = req.headers["authorization"] || null;
    if (token) {
        if (token.startsWith("Bearer ")) {
            token = token.slice(7).trim();
        }
        jwt.verify(token, config.accessToken, (err, user) => {
            if (!err && user) {
                req.headers["user-id"] = user.userId;
                req.headers["user-role"] = user.role;
                return next();
            } else {
                return res.status(apiStatus.UNAUTHORIZED.status).send();
            }
        });
    } else {
        return res.status(apiStatus.UNAUTHORIZED.status).send();
    }
};

module.exports = authenticateToken;
