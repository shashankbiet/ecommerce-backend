const apiStatus = require("../app/shared/constants/api-status");
const { getCorrelationId } = require("../util/request-header");

let notFoundHandler = async (req, res, next) => {
    console.error(
        `CorrelationId: ${getCorrelationId(req)}, Method: ${req.method}, Url: ${
            req.originalUrl
        },  Message: Url not found`
    );
    return res.status(apiStatus.NOT_FOUND.status).send({
        success: false,
        result: apiStatus.NOT_FOUND.result,
    });
};

module.exports = notFoundHandler;
