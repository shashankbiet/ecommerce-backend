const userModel = require("../shared/models/user-model");
const hashing = require("../../util/hashing");
const authService = {};

authService.authenticate = async (email, phoneNumber, plainPass) => {
    const orConditions = [];
    if (email) orConditions.push({ email });
    if (phoneNumber) orConditions.push({ phoneNumber });
    if (orConditions.length === 0) {
        return null;
    }

    const user = await userModel.findOne({ isActive: true, $or: orConditions });
    if (user && (await hashing.compare(plainPass, user.password))) {
        return { userId: user.userId, role: user.role };
    }
    return null;
};

module.exports = authService;
