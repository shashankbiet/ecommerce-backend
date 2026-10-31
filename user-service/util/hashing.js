const bcrypt = require("bcrypt");
const hashing = {};

hashing.encrypt = async (password) => {
    const saltRounds = 10;
    return bcrypt.hash(password, saltRounds);
};

hashing.compare = async (plainPass, hashPass) => {
    return bcrypt.compare(plainPass, hashPass);
};

module.exports = hashing;
