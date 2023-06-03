const { body, param, check } = require("express-validator");
const addressTypes = require("../shared/constants/address-types");
const addressValidator = {};

addressValidator.add = () => {
    return [
        body("type")
            .notEmpty()
            .withMessage("Empty type")
            .isIn(Object.values(addressTypes))
            .withMessage("Invalid type"),
        body("line1").notEmpty().withMessage("Empty line1"),
        body("district").notEmpty().withMessage("Empty district"),
        body("landmark").notEmpty().withMessage("Empty landmark"),
        body("state").notEmpty().withMessage("Empty state"),
        body("pincode")
            .isInt()
            .isLength({ max: 6, min: 6 })
            .withMessage("Pincode must be 6 digits long"),
    ];
};

addressValidator.getById = () => {
    return [param("addressId").isMongoId()];
};

addressValidator.update = () => {
    return [
        param("addressId").isMongoId(),
        body("type")
            .notEmpty()
            .withMessage("Empty type")
            .isIn(Object.values(addressTypes))
            .withMessage("Invalid type"),
        body("line1").notEmpty().withMessage("Empty line1"),
        body("district").notEmpty().withMessage("Empty district"),
        body("landmark").notEmpty().withMessage("Empty landmark"),
        body("state").notEmpty().withMessage("Empty state"),
        body("pincode").isInt().isLength({ max: 6, min: 6 }),
    ];
};

addressValidator.delete = () => {
    return [param("addressId").isMongoId()];
};

module.exports = addressValidator;
