const js = require("@eslint/js");
const globals = require("globals");
const prettier = require("eslint-config-prettier");

module.exports = [
    js.configs.recommended,
    {
        files: ["**/*.js"],
        languageOptions: {
            ecmaVersion: 2022,
            sourceType: "commonjs",
            globals: { ...globals.node },
        },
        rules: {
            // Express error handlers must keep the (err, req, res, next) arity.
            "no-unused-vars": ["error", { args: "none", caughtErrors: "none" }],
        },
    },
    {
        files: ["test/**/*.js"],
        languageOptions: { globals: { ...globals.jest } },
    },
    prettier,
];
