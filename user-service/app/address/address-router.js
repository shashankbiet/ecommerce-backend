const addressController = require("./address-controller");
const router = require("express").Router();
const addressValidator = require("./address-validator");
const validate = require("../../util/validator");

router.post("/", addressValidator.add(), validate, addressController.add);
router.get("/", addressController.get);
router.get(
    "/:address_id",
    addressValidator.getById(),
    validate,
    addressController.getById
);
router.put(
    "/:address_id",
    addressValidator.update(),
    validate,
    addressController.update
);
router.delete(
    "/:address_id",
    addressValidator.delete(),
    validate,
    addressController.delete
);

module.exports = router;
