const addressController = require("./address-controller");
const router = require("express").Router();
const addressValidator = require("./address-validator");
const validate = require("../../util/validator");

router.post(
    "/:userId",
    addressValidator.add(),
    validate,
    addressController.add
);
router.get("/:userId", addressValidator.get(), validate, addressController.get);
router.get(
    "/:userId/:addressId",
    addressValidator.getById(),
    validate,
    addressController.getById
);
router.put(
    "/:userId/:addressId",
    addressValidator.update(),
    validate,
    addressController.update
);
router.delete(
    "/:userId/:addressId",
    addressValidator.delete(),
    validate,
    addressController.delete
);

module.exports = router;
