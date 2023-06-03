module.exports = {
    getCorrelationId: (req) => req.headers["correlation-id"],
    getUserId: (req) => req.headers["user-id"],
    getUserRole: (req) => req.headers["user-role"],
};
