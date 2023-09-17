const userService = require("./user-service");
const apiStatus = require("../shared/constants/api-status");
const ApiError = require("../shared/models/api-error");
const userRoles = require("../shared/constants/user-roles");
const { getUserId } = require("../../util/request-header");
const {
    createUserReqCounter,
    getUserReqCounter,
} = require("../../util/metrics");
const userController = {};

userController.register = async (req, res, next) => {
    createUserReqCounter.inc();
    try {
        if (
            req.body.role == userRoles.ADMIN &&
            !userService.isAdmin(req.body.token)
        ) {
            return res.status(apiStatus.UNAUTHORIZED.status).send();
        }
        let result = await userService.register(req.body);
        if (result) {
            return res.status(apiStatus.CREATED.status).send();
        }
        return res.status(apiStatus.CONFLICT.status).send();
    } catch (err) {
        next(ApiError.internalServerError(err.message, err.stack));
    }
};

userController.get = async (req, res, next) => {
    getUserReqCounter.inc();
    try {
        let result = await userService.get(parseInt(getUserId(req)));
        if (result) {
            return res.status(apiStatus.OK.status).send(result);
        }
        return res.status(apiStatus.NOT_FOUND.status).send();
    } catch (err) {
        next(ApiError.internalServerError(err.message, err.stack));
    }
};

userController.update = async (req, res, next) => {
    try {
        let result = await userService.update(
            parseInt(getUserId(req)),
            req.body
        );
        if (result === true) {
            return res.status(apiStatus.OK.status).send();
        }
        return res.status(apiStatus.NOT_FOUND.status).send();
    } catch (err) {
        next(ApiError.internalServerError(err.message, err.stack));
    }
};

userController.delete = async (req, res, next) => {
    try {
        let result = await userService.delete(parseInt(getUserId(req)));
        if (result === true) {
            return res.status(apiStatus.OK.status).send();
        }
        return res.status(apiStatus.NOT_FOUND.status).send();
    } catch (err) {
        next(ApiError.internalServerError(err.message, err.stack));
    }
};

userController.updatePassword = async (req, res, next) => {
    let { currentPassword, newPassword } = req.body;
    try {
        if (currentPassword == newPassword) {
            return res.status(apiStatus.BAD_REQUEST.status).send({
                success: false,
                result: "Current and New Password cannot be same",
            });
        }
        let result = await userService.updatePassword(
            parseInt(getUserId(req)),
            req.body.currentPassword,
            req.body.newPassword
        );
        if (result === null) {
            return res.status(apiStatus.NOT_FOUND.status).send();
        } else if (result === true) {
            return res.status(apiStatus.OK.status).send();
        }
        return res.status(apiStatus.BAD_REQUEST.status).send({
            success: false,
            result: "Current Password Incorrect",
        });
    } catch (err) {
        next(ApiError.internalServerError(err.message, err.stack));
    }
};

module.exports = userController;
