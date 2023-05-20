const authService = require("./auth-service");
const apiStatus = require("../shared/constants/api-status");
const ApiError = require("../shared/models/api-error");
const authController = {};

authController.authenticate = async (req, res, next) => {
    try {
        let { email, phoneNumber, password } = req.body;
        let user = await authService.authenticate(email, phoneNumber, password);
        if (user == null) {
            return res.status(apiStatus.UNAUTHORIZED.status).send();
        }
        return res.status(apiStatus.OK.status).send(user);
    } catch (err) {
        next(ApiError.internalServerError(err.message, err.stack));
    }
};

module.exports = authController;
