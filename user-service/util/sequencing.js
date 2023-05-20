const mongoose = require("mongoose");
const collectionNames = require("../app/shared/constants/collection-names");
const USER_COUNTER_ID = "userid",
    USER_COUNTER_START = 1001;
const counterSchema = mongoose.Schema({
    _id: {
        type: String,
        required: true,
    },
    seq: {
        type: Number,
        required: true,
    },
});

const Counter = mongoose.model(collectionNames.COUNTER, counterSchema);

const getNewUserId = () => {
    return new Promise((resolve, reject) => {
        Counter.findOneAndUpdate(
            { _id: USER_COUNTER_ID },
            { $inc: { seq: 1 } },
            { new: true },
            async (err, result) => {
                if (result == null) {
                    const newValue = new Counter({
                        _id: USER_COUNTER_ID,
                        seq: USER_COUNTER_START,
                    });
                    await newValue.save();
                    return resolve(USER_COUNTER_START);
                } else {
                    return resolve(result.seq);
                }
            }
        );
    });
};

module.exports = { getNewUserId };
