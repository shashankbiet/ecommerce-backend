const { v4 } = require("uuid");
const apiStatus = require("../app/shared/constants/api-status");
const logger = require("../handler/logger");

const addCorrelationId = (req, res, next) => {
    try {
        req.headers["correlation-id"] = v4();
        next();
    } catch (err) {
        logger.error(`Error in adding correlationId, err: ${err}`);
        let status = apiStatus.INTERNAL_SERVER_ERROR.status,
            result = apiStatus.INTERNAL_SERVER_ERROR.result;
        return res.status(status).send({
            success: false,
            result: result,
        });
    }
};

module.exports = addCorrelationId;
