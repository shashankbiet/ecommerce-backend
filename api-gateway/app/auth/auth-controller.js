const authService = require("./auth-service");
const apiStatus = require("../shared/constants/api-status");
const axios = require("axios");
const config = require("../../config");
const { getCorrelationId } = require("../../util/request-header");
const logger = require("../../handler/logger");
const authController = {};

authController.login = async (req, res, next) => {
    try {
        const { email, phoneNumber, password } = req.body;
        const user = await axios.post(
            config.userService.target +
                config.userService.authPath +
                "/authenticate",
            { email, phoneNumber, password },
            {
                headers: {
                    "correlation-id": getCorrelationId(req),
                },
                timeout: 5000,
            }
        );
        const token = await authService.createToken(user.data);
        return res.status(apiStatus.OK.status).send(token);
    } catch (err) {
        if (
            err &&
            err.response &&
            err.response.status == apiStatus.UNAUTHORIZED.status
        ) {
            return res.status(apiStatus.UNAUTHORIZED.status).send();
        }

        logger.error(
            `CorrelationId: ${getCorrelationId(req)},\nstatus: ${
                apiStatus.INTERNAL_SERVER_ERROR.status
            },\nmessage: ${err.message},\nstack: ${err.stack}`
        );
        return res.status(apiStatus.INTERNAL_SERVER_ERROR.status).send();
    }
};

module.exports = authController;
