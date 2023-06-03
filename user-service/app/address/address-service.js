const addressService = {};
const userModel = require("../shared/models/user-model");
const mongoose = require("mongoose");

addressService.add = async (userId, body) => {
    try {
        let { type, line1, line2, district, landmark, state, pincode } = body;
        let address = {
            addressId: new mongoose.Types.ObjectId(),
            type,
            line1,
            line2,
            district,
            landmark,
            state,
            pincode,
        };
        let result = await userModel.findOneAndUpdate(
            { userId: userId, isActive: true },
            { $push: { address: address } }
        );
        if (result) {
            return true;
        }
        return false;
    } catch (err) {
        throw new Error(err);
    }
};

addressService.get = async (userId, addressId) => {
    try {
        if (addressId) {
            let user = await userModel.findOne(
                {
                    userId: userId,
                    isActive: true,
                    "address.addressId": addressId,
                },
                { "address.$": 1 }
            );
            if (user) {
                return user.toObject().address.length == 1
                    ? user.toObject().address[0]
                    : null;
            }
            return null;
        } else {
            let user = await userModel.findOne(
                { userId: userId, isActive: true },
                { address: 1 }
            );
            if (user) {
                return user.toObject().address;
            }
            return null;
        }
    } catch (err) {
        throw new Error(err);
    }
};

addressService.update = async (userId, addressId, body) => {
    try {
        let { type, line1, line2, district, landmark, state, pincode } = body;
        let result = await userModel.findOneAndUpdate(
            { userId: userId, isActive: true, "address.addressId": addressId },
            {
                $set: {
                    "address.$.type": type,
                    "address.$.line1": line1,
                    "address.$.line2": line2,
                    "address.$.district": district,
                    "address.$.landmark": landmark,
                    "address.$.state": state,
                    "address.$.pincode": pincode,
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

addressService.delete = async (userId, addressId) => {
    try {
        let result = await userModel.findOneAndUpdate(
            { userId: userId, isActive: true, "address.addressId": addressId },
            { $pull: { address: { addressId: addressId } } }
        );
        if (result) {
            return true;
        }
        return false;
    } catch (err) {
        throw new Error(err);
    }
};

module.exports = addressService;
