const router = require("express").Router();
const status = require("../app/shared/constants/api-status");
const moment = require("moment");

function healthCheck(req, res) {
    return res.status(status.OK.status).send({
        uptime: moment()
            .subtract(process.uptime(), "seconds")
            .format("MM-DD-YYYY HH:mm:ss"),
        message: "OK",
        timestamp: moment().format("MM-DD-YYYY HH:mm:ss"),
    });
}

router.get("/", healthCheck);

module.exports = router;
