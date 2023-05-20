const userModel = require("../shared/models/user-model");
const hashing = require("../../util/hashing");
const authService = {};

authService.authenticate = async (email, phoneNumber, plainPass) => {
    try {
        let user = await userModel.findOne({
            $or: [{ email: email }, { phoneNumber: phoneNumber }],
            isActive: true,
        });
        if (user) {
            let userId = user["userId"],
                hashPass = user["password"];
            let auth = await hashing.compare(plainPass, hashPass);
            if (auth) {
                return { userId, role: user["role"] };
            }
        }
        return null;
    } catch (err) {
        throw new Error(err);
    }
};

module.exports = authService;
