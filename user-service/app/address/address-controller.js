const addressService = require("./address-service");
const apiStatus = require("../shared/constants/api-status");
const ApiError = require("../shared/models/api-error");
const addressController = {};

addressController.add = async (req, res, next) => {
    try {
        let result = await addressService.add(
            parseInt(req.params.userId),
            req.body
        );
        if (result) {
            return res.status(apiStatus.CREATED.status).send();
        }
        return next(ApiError.notFound());
    } catch (err) {
        return next(ApiError.internalServerError(err.message, err.stack));
    }
};

addressController.getById = async (req, res, next) => {
    try {
        let result = await addressService.get(
            parseInt(req.params.userId),
            req.params.addressId
        );
        if (result) {
            return res.status(apiStatus.OK.status).send(result);
        }
        return next(ApiError.notFound());
    } catch (err) {
        return next(ApiError.internalServerError(err.message, err.stack));
    }
};

addressController.get = async (req, res, next) => {
    try {
        let result = await addressService.get(parseInt(req.params.userId));
        if (result) {
            return res.status(apiStatus.OK.status).send(result);
        }
        return next(ApiError.notFound());
    } catch (err) {
        return next(ApiError.internalServerError(err.message, err.stack));
    }
};

addressController.update = async (req, res, next) => {
    try {
        let result = await addressService.update(
            parseInt(req.params.userId),
            req.params.addressId,
            req.body
        );
        if (result) {
            return res.status(apiStatus.OK.status).send();
        }
        return next(ApiError.notFound());
    } catch (err) {
        return next(ApiError.internalServerError(err.message, err.stack));
    }
};

addressController.delete = async (req, res, next) => {
    try {
        let result = await addressService.delete(
            parseInt(req.params.userId),
            req.params.addressId
        );
        if (result) {
            return res.status(apiStatus.OK.status).send();
        }
        return next(ApiError.notFound());
    } catch (err) {
        return next(ApiError.internalServerError(err.message, err.stack));
    }
};

module.exports = addressController;
