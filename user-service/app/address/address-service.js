const addressService = {};
const userModel = require("../shared/models/user-model");
const mongoose = require("mongoose");

addressService.add = async (userId, body) => {
    const { type, line1, line2, district, landmark, state, pincode } = body;
    const address = {
        addressId: new mongoose.Types.ObjectId(),
        type,
        line1,
        line2,
        district,
        landmark,
        state,
        pincode,
    };
    const result = await userModel.findOneAndUpdate(
        { userId, isActive: true },
        { $push: { address } }
    );
    return Boolean(result);
};

addressService.get = async (userId, addressId) => {
    if (addressId) {
        const user = await userModel.findOne(
            { userId, isActive: true, "address.addressId": addressId },
            { "address.$": 1 }
        );
        return user && user.address && user.address.length > 0
            ? user.address[0]
            : null;
    }
    const user = await userModel.findOne(
        { userId, isActive: true },
        { address: 1 }
    );
    return user ? user.address || [] : null;
};

addressService.update = async (userId, addressId, body) => {
    const { type, line1, line2, district, landmark, state, pincode } = body;
    const result = await userModel.findOneAndUpdate(
        { userId, isActive: true, "address.addressId": addressId },
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
    return Boolean(result);
};

addressService.delete = async (userId, addressId) => {
    const result = await userModel.findOneAndUpdate(
        { userId, isActive: true, "address.addressId": addressId },
        { $pull: { address: { addressId } } }
    );
    return Boolean(result);
};

module.exports = addressService;
