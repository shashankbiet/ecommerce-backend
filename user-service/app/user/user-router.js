const userController = require("./user-controller");
const userValidator = require("./user-validator");
const validate = require("../../util/validator");
const router = require("express").Router();

router.post(
    "/register",
    userValidator.register(),
    validate,
    userController.register
);

router.get("/", userController.get);

router.put("/", userValidator.update(), validate, userController.update);

router.delete("/", userController.delete);

router.put(
    "/password",
    userValidator.updatePassword(),
    validate,
    userController.updatePassword
);

module.exports = router;
