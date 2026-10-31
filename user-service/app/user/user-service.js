const userModel = require("../shared/models/user-model");
const mongoError = require("../shared/constants/mongo-error");
const sequencing = require("../../util/sequencing");
const hashing = require("../../util/hashing");
const { adminVerificationToken } = require("../../config");
const { dbResponseTimeHistogram } = require("../../util/metrics");
const userService = {};

userService.isAdmin = (token) => {
    return Boolean(adminVerificationToken) && token === adminVerificationToken;
};

userService.register = async (body) => {
    const metricsLabels = {
        operation: "CreateUser",
    };
    let timer = dbResponseTimeHistogram.startTimer();
    try {
        let { phoneNumber, email, password, role, firstName, lastName } = body;
        let dbUser = await userModel.findOne({
            $or: [{ email: email }, { phoneNumber: phoneNumber }],
        });
        if (dbUser) {
            timer({ ...metricsLabels, success: false });
            return null;
        }
        password = await hashing.encrypt(password);
        let userId = await sequencing.getNewUserId();
        let user = new userModel({
            userId,
            phoneNumber,
            email,
            password,
            role,
            firstName,
            lastName,
        });

        await user.save();
        timer({ ...metricsLabels, success: true });
        return userId;
    } catch (err) {
        timer({ ...metricsLabels, success: false });
        if (err.code && err.code == mongoError.DUPLICATE_KEY.code) {
            return null;
        }
        throw err;
    }
};

userService.get = async (userId) => {
    const metricsLabels = {
        operation: "GetUser",
    };
    let timer = dbResponseTimeHistogram.startTimer();
    try {
        let user = await userModel.findOne(
            { userId: userId, isActive: true },
            { password: 0, _id: 0, address: 0 }
        );
        timer({ ...metricsLabels, success: true });
        if (user) {
            return user.toObject();
        }
        return null;
    } catch (err) {
        timer({ ...metricsLabels, success: false });
        throw err;
    }
};

userService.update = async (userId, body) => {
    const { firstName, lastName } = body;
    const result = await userModel.findOneAndUpdate(
        { userId, isActive: true },
        { $set: { firstName, lastName, lastModifiedDate: new Date() } }
    );
    return Boolean(result);
};

userService.delete = async (userId) => {
    const result = await userModel.findOneAndUpdate(
        { userId, isActive: true },
        { $set: { isActive: false, lastModifiedDate: new Date() } }
    );
    return Boolean(result);
};

userService.updatePassword = async (userId, currentPassword, newPassword) => {
    const user = await userModel.findOne(
        { userId, isActive: true },
        { password: 1 }
    );
    if (!user) {
        return null;
    }
    const auth = await hashing.compare(currentPassword, user.password);
    if (!auth) {
        return false;
    }
    await userModel.findOneAndUpdate(
        { userId, isActive: true },
        {
            $set: {
                password: await hashing.encrypt(newPassword),
                lastModifiedDate: new Date(),
            },
        }
    );
    return true;
};

module.exports = userService;
