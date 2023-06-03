const express = require("express");
const router = require("./handler/router");
const apiErrorHandler = require("./middleware/api-error-handler");
const notFoundHandler = require("./middleware/not-found-handler");
const mongodb = require("./handler/mongodb");
const addApiLog = require("./middleware/add-api-log");
const app = express();

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

//Database
mongodb();

app.use(function (req, res, next) {
    //Remove headers
    res.removeHeader("X-Powered-By");

    //Enabling CORS
    res.header("Access-Control-Allow-Origin", "*");

    res.header(
        "Access-Control-Allow-Methods",
        "GET, POST, PUT, DELETE, HEAD, OPTIONS"
    );
    res.header(
        "Access-Control-Allow-Headers",
        "Origin, X-Requested-With, ContentType, Content-Type, Accept, Authorization"
    );
    res.header(
        "Strict-Transport-Security",
        "max-age=15552000; includeSubDomains; preload;"
    );
    res.header("X-Frame-Options", "DENY");
    res.header("X-XSS-Protection", "1; mode=block");
    res.header("X-Content-Type-Options", "nosniff");
    res.header(
        "Content-Security-Policy",
        "default-src 'none'; script-src 'self'; style-src 'self'; object-src 'none'"
    );
    res.header("Referrer-Policy", "same-origin");
    next();
});

//Add Api Log
app.use(addApiLog);

//Routes
router(app);

//Error Handling
app.use(apiErrorHandler);

//Not Found
app.use(notFoundHandler);

module.exports = app;
