const { validationResult } = require("express-validator");
const status = require("../app/shared/constants/api-status");

const validate = (req, res, next) => {
    const errors = validationResult(req);
    if (errors.isEmpty()) {
        return next();
    }
    const extractedErrors = [];
    errors.array().map((err) => extractedErrors.push({ [err.param]: err.msg }));

    return res.status(status.BAD_REQUEST.status).json({
        errors: extractedErrors,
    });
};

module.exports = validate;
