const router = require("express").Router();
const authController = require("./auth-controller");
const authValidator = require("./auth-validator");
const validate = require("../../util/validator");

router.post(
    "/authenticate",
    authValidator.authenticate(),
    validate,
    authController.authenticate
);

module.exports = router;
