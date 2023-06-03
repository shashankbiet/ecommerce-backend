const jwt = require("jsonwebtoken");
const config = require("../config");
const apiStatus = require("../app/shared/constants/api-status");
const { isPublicUrl } = require("../util/request-url");

const authenticateToken = (req, res, next) => {
    if (isPublicUrl(req.originalUrl)) {
        next();
    } else {
        const token = req.headers["authorization"]
            ? req.headers["authorization"]
            : null;
        if (token) {
            jwt.verify(token, config.accessToken, (err, user) => {
                if (!err) {
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
    }
};

module.exports = authenticateToken;
