module.exports = {
    OK: {
        status: 200,
        result: "The request has succeeded.",
    },
    CREATED: {
        status: 201,
        result: "The request has succeeded and a new resource has been created.",
    },
    ACCEPTED: {
        status: 202,
        result: "The request has been received but not yet acted upon.",
    },
    BAD_REQUEST: {
        status: 400,
        result: "The server could not understand the request due to invalid syntax.",
    },
    UNAUTHORIZED: {
        status: 401,
        result: "The request sent by the client could not be authenticated.",
    },
    FORBIDDEN: {
        status: 403,
        result: "Access to the requested resource is forbidden.",
    },
    NOT_FOUND: {
        status: 404,
        result: "The server can not find the requested resource.",
    },
    INTERNAL_SERVER_ERROR: {
        status: 500,
        result: "The server has encountered a situation it doesnt know how to handle.",
    },
    BAD_GATEWAY: {
        status: 502,
        result: "The server was not able to get a valid or any response",
    },
    SERVICE_UNAVAILABLE: {
        status: 503,
        result: "The server is currently unable to handle the request.",
    },
    CONFLICT: {
        status: 409,
        result: "The request conflicts with the current state of the server",
    },
};
