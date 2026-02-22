var bcrypt = require("bcrypt");
let hashing = {};

hashing.encrypt = async (password) => {
    const saltRounds = 10;
    let hashPass = await bcrypt.hash(password, saltRounds);
    return hashPass;
};

hashing.compare = async (plainPass, hashPass) => {
    const match = await bcrypt.compare(plainPass, hashPass);
    if (match) return true;
    else return false;
};

module.exports = hashing;
