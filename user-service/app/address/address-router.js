const addressController = require("./address-controller");
const router = require("express").Router();
const addressValidator = require("./address-validator");
const validate = require("../../util/validator");

router.post("/", addressValidator.add(), validate, addressController.add);
router.get("/", addressController.get);
router.get(
    "/:addressId",
    addressValidator.getById(),
    validate,
    addressController.getById
);
router.put(
    "/:addressId",
    addressValidator.update(),
    validate,
    addressController.update
);
router.delete(
    "/:addressId",
    addressValidator.delete(),
    validate,
    addressController.delete
);

module.exports = router;
