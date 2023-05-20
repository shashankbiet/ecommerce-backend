const apiStatus = require("../constants/api-status");

class ApiError {
    constructor(status, result, message, stack) {
        this.status = status;
        this.result = result;
        this.message = message;
        this.stack = stack;
    }

    static internalServerError(message, stack) {
        return new ApiError(
            apiStatus.INTERNAL_SERVER_ERROR.status,
            apiStatus.INTERNAL_SERVER_ERROR.result,
            message,
            stack
        );
    }

    static notFound(message, stack) {
        return new ApiError(
            apiStatus.NOT_FOUND.status,
            apiStatus.NOT_FOUND.result,
            message,
            stack
        );
    }

    static badRequest(message, stack) {
        return new ApiError(
            apiStatus.BAD_REQUEST.status,
            apiStatus.BAD_REQUEST.result,
            message,
            stack
        );
    }
}

module.exports = ApiError;
