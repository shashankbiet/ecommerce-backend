const userModel = require("../shared/models/user-model");
const mongoError = require("../shared/constants/mongo-error");
const sequencing = require("../../util/sequencing");
const hashing = require("../../util/hashing");
require("dotenv").config();
const userService = {};

userService.isAdmin = (token) => {
    return token == process.env.ADMIN_VERIFICATION_TOKEN_SECRET;
};

userService.register = async (body) => {
    try {
        let { phoneNumber, email, password, role, firstName, lastName } = body;
        password = await hashing.encrypt(password);
        let dbUser = await userModel.findOne({
            $or: [{ email: email }, { phoneNumber: phoneNumber }],
        });
        if (dbUser) {
            return null;
        } else {
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
            return userId;
        }
    } catch (err) {
        if (err.code && err.code == mongoError.DUPLICATE_KEY.code) {
            return null;
        }
        throw new Error(err);
    }
};

userService.get = async (userId) => {
    try {
        let user = await userModel.findOne(
            { userId: userId, isActive: true },
            { password: 0, _id: 0, address: 0 }
        );
        if (user) {
            return user.toObject();
        }
        return null;
    } catch (err) {
        throw new Error(err);
    }
};

userService.update = async (userId, body) => {
    try {
        let { firstName, lastName } = body;
        let result = await userModel.findOneAndUpdate(
            { userId: userId, isActive: true },
            {
                $set: {
                    firstName: firstName,
                    lastName: lastName,
                    modifiedDate: new Date(),
                },
            }
        );
        if (result) {
            return true;
        }
        return false;
    } catch (err) {
        throw new Error(err);
    }
};

userService.delete = async (userId) => {
    try {
        let result = await userModel.findOneAndUpdate(
            { userId: userId, isActive: true },
            {
                $set: {
                    isActive: false,
                    modifiedDate: new Date(),
                },
            }
        );
        if (result) {
            return true;
        }
        return false;
    } catch (err) {
        throw new Error(err);
    }
};

userService.updatePassword = async (userId, currentPassword, newPassword) => {
    try {
        let user = await userModel.findOne(
            { userId: userId, isActive: true },
            { password: 1 }
        );
        if (user) {
            let hashPass = user["password"];
            let auth = await hashing.compare(currentPassword, hashPass);
            if (auth) {
                newPassword = await hashing.encrypt(newPassword);
                await userModel.findOneAndUpdate(
                    { userId: userId, isActive: true },
                    {
                        $set: {
                            password: newPassword,
                            modifiedDate: new Date(),
                        },
                    }
                );
                return true;
            }
            return false;
        }
        return null;
    } catch (err) {
        throw new Error(err);
    }
};

module.exports = userService;
