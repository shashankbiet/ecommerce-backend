const apiStatus = require("../app/shared/constants/api-status");
const ApiError = require("../app/shared/models/api-error");
const { getCorrelationId } = require("../util/request-header");

let apiErrorHandler = (err, req, res, next) => {
    let status = apiStatus.INTERNAL_SERVER_ERROR.status,
        result = apiStatus.INTERNAL_SERVER_ERROR.result;
    if (err instanceof ApiError) {
        status = err.status;
        result = err.result;
        if (err.message || err.stack) {
            console.error(
                `CorrelationId: ${getCorrelationId(req)},\nstatus: ${
                    err.status
                },\nmessage: ${err.message},\nstack: ${err.stack}`
            );
        }
    } else {
        console.error(
            `CorrelationId: ${getCorrelationId(req)},\nerror: ${err}`
        );
    }

    return res.status(status).send({
        success: false,
        result,
    });
};

module.exports = apiErrorHandler;
