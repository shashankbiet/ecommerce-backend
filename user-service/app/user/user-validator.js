const { body } = require("express-validator");
const userRoles = require("../shared/constants/user-roles");
const userValidator = {};

userValidator.register = () => {
    return [
        body("phoneNumber")
            .isInt()
            .isLength({ max: 10, min: 10 })
            .withMessage("Phone number must be 10 digits long"),
        body("email").isEmail().withMessage("Invalid email address"),
        body("password")
            .exists()
            .withMessage("Password is required")
            .isString()
            .withMessage("Password should be string")
            .isLength({ min: 5 })
            .withMessage("Password should be at least 5 characters"),
        body("role").isIn(Object.values(userRoles)),
        body("token")
            .if(body("role").equals(userRoles.ADMIN))
            .notEmpty()
            .withMessage("Invalid Token"),
        body("firstName").notEmpty().withMessage("Empty first name"),
        body("lastName").notEmpty().withMessage("Empty last name"),
    ];
};

userValidator.update = () => {
    return [
        body("firstName").notEmpty().withMessage("Empty first name"),
        body("lastName").notEmpty().withMessage("Empty last name"),
    ];
};

userValidator.updatePassword = () => {
    return [
        body("currentPassword")
            .notEmpty()
            .withMessage("Current Password empty"),
        body("newPassword")
            .exists()
            .withMessage("New Password is required")
            .isString()
            .withMessage("New Password should be string")
            .isLength({ min: 5 })
            .withMessage("New Password should be at least 5 characters"),
    ];
};

module.exports = userValidator;
