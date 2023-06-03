const app = require("./app");
const config = require("./config");
//Setting up server
const port = config.port || 3002;
app.listen(port, () => console.log(`User Service running on port ${port}`));
