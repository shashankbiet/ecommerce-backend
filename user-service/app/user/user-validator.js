const { body } = require("express-validator");
const userRoles = require("../shared/constants/user-roles");
const userValidator = {};

userValidator.register = () => {
    return [
        body("phone_number")
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
        body("first_name").notEmpty().withMessage("Empty first name"),
        body("last_name").notEmpty().withMessage("Empty last name"),
    ];
};

userValidator.update = () => {
    return [
        body("first_name").notEmpty().withMessage("Empty first name"),
        body("last_name").notEmpty().withMessage("Empty last name"),
    ];
};

userValidator.updatePassword = () => {
    return [
        body("current_password")
            .notEmpty()
            .withMessage("Current Password empty"),
        body("new_password")
            .exists()
            .withMessage("New Password is required")
            .isString()
            .withMessage("New Password should be string")
            .isLength({ min: 5 })
            .withMessage("New Password should be at least 5 characters"),
    ];
};

module.exports = userValidator;
