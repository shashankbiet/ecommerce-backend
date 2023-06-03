const router = require("express").Router();
const authController = require("./auth-controller");
const authValidator = require("./auth-validator");
const validate = require("../../util/validator");

router.post("/login", authValidator.login(), validate, authController.login);

module.exports = router;
