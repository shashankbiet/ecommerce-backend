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

const getNewUserId = async () => {
    let counter = await Counter.findOneAndUpdate(
        { _id: USER_COUNTER_ID },
        { $inc: { seq: 1 } },
        { new: true }
    );
    if (!counter) {
        try {
            counter = await Counter.create({
                _id: USER_COUNTER_ID,
                seq: USER_COUNTER_START,
            });
            return counter.seq;
        } catch (err) {
            // Concurrent initialization race condition: fetch newly created counter
            counter = await Counter.findOneAndUpdate(
                { _id: USER_COUNTER_ID },
                { $inc: { seq: 1 } },
                { new: true }
            );
        }
    }
    return counter.seq;
};

module.exports = { getNewUserId };
