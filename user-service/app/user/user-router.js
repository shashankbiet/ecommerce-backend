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

router.get("/:userId", userValidator.get(), validate, userController.get);

router.put("/:userId", userValidator.update(), validate, userController.update);

router.delete(
    "/:userId",
    userValidator.delete(),
    validate,
    userController.delete
);

router.put(
    "/password/:userId",
    userValidator.updatePassword(),
    validate,
    userController.updatePassword
);

module.exports = router;
