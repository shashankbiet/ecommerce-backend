const app = require("./app");
require("dotenv").config();

//Setting up server
const port = process.env.PORT || 3002;
app.listen(port, () => console.log(`User Service running on port ${port}`));
