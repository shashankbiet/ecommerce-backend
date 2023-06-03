const { getCorrelationId } = require("../util/request-header");
const addApiLog = (req, res, next) => {
    try {
        console.log(
            `CorrelationId: ${getCorrelationId(req)}, Method: ${
                req.method
            }, URL: ${req.originalUrl}}`
        );
        next();
    } catch (err) {
        console.error(`Error in adding request log, err: ${err}`);
        let status = apiStatus.INTERNAL_SERVER_ERROR.status,
            result = apiStatus.INTERNAL_SERVER_ERROR.result;
        return res.status(status).send({
            success: false,
            result: result,
        });
    }
};

module.exports = addApiLog;
