const { oneOf, body } = require("express-validator");
const authValidator = {};

authValidator.login = () => {
    return [
        oneOf([
            body("email").isEmail().withMessage("Invalid email address"),
            body("phoneNumber")
                .isNumeric()
                .withMessage("Phone number must be numeric")
                .isLength({ max: 10, min: 10 })
                .withMessage("Phone number must be 10 digits long"),
        ]),
        body("password").isString().notEmpty(),
    ];
};

module.exports = authValidator;
