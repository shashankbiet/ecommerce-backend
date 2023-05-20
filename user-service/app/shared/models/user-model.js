const mongoose = require("mongoose");
const Schema = mongoose.Schema;
const collectionName = require("../constants/collection-names");
const userRoles = require("../constants/user-roles");
const addressTypes = require("../constants/address-types");
const ObjectId = Schema.ObjectId;
const userSchema = new Schema(
    {
        userId: {
            type: Number,
            unique: true,
        },
        phoneNumber: {
            type: String,
            required: true,
            unique: true,
        },
        email: {
            type: String,
            required: true,
            unique: true,
        },
        password: {
            type: String,
            required: true,
        },
        role: {
            type: String,
            required: true,
            enum: Object.values(userRoles),
        },
        firstName: {
            type: String,
            required: true,
        },
        lastName: {
            type: String,
            required: false,
        },
        createdDate: {
            type: Date,
            required: false,
            default: Date.now,
        },
        lastModifiedDate: {
            type: Date,
            required: false,
        },
        isActive: {
            type: Boolean,
            required: true,
            default: true,
        },
        address: [
            {
                addressId: {
                    type: ObjectId,
                    required: true,
                },
                type: {
                    type: String,
                    required: true,
                    enum: Object.values(addressTypes),
                },
                line1: {
                    type: String,
                    required: true,
                },
                line2: {
                    type: String,
                    required: false,
                },
                district: {
                    type: String,
                    required: true,
                },
                landmark: {
                    type: String,
                    required: false,
                },
                state: {
                    type: String,
                    required: true,
                },
                pincode: {
                    type: Number,
                    required: true,
                },
                _id: false,
            },
        ],
    },
    { versionKey: false }
);

module.exports = mongoose.model(collectionName.USER, userSchema);
