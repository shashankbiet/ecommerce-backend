const addressService = {};
const { query } = require("../../handler/postgres");

addressService.add = async (userId, body) => {
    const { type, line1, line2, district, landmark, state, pincode } = body;
    const userResult = await query(
        `SELECT user_id
         FROM users
         WHERE user_id = $1 AND is_active = true`,
        [userId]
    );
    if (userResult.rowCount === 0) {
        return false;
    }
    const insertResult = await query(
        `INSERT INTO user_addresses
            (user_id, type, line1, line2, district, landmark, state, pincode)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
        [userId, type, line1, line2, district, landmark, state, pincode]
    );
    return insertResult.rowCount > 0;
};

addressService.get = async (userId, addressId) => {
    if (addressId) {
        const result = await query(
            `SELECT a.address_id, a.type, a.line1, a.line2, a.district, a.landmark, a.state, a.pincode
             FROM user_addresses a
             JOIN users u ON u.user_id = a.user_id
             WHERE a.user_id = $1 AND a.address_id = $2 AND u.is_active = true`,
            [userId, addressId]
        );
        if (result.rowCount === 0) {
            return null;
        }
        const row = result.rows[0];
        return {
            address_id: row.address_id,
            type: row.type,
            line1: row.line1,
            line2: row.line2,
            district: row.district,
            landmark: row.landmark,
            state: row.state,
            pincode: row.pincode,
        };
    }
    const userResult = await query(
        `SELECT user_id
         FROM users
         WHERE user_id = $1 AND is_active = true`,
        [userId]
    );
    if (userResult.rowCount === 0) {
        return null;
    }
    const result = await query(
        `SELECT address_id, type, line1, line2, district, landmark, state, pincode
         FROM user_addresses
         WHERE user_id = $1
         ORDER BY created_at ASC`,
        [userId]
    );
    return result.rows.map((row) => ({
        address_id: row.address_id,
        type: row.type,
        line1: row.line1,
        line2: row.line2,
        district: row.district,
        landmark: row.landmark,
        state: row.state,
        pincode: row.pincode,
    }));
};

addressService.update = async (userId, addressId, body) => {
    const { type, line1, line2, district, landmark, state, pincode } = body;
    const result = await query(
        `UPDATE user_addresses a
         SET type = $1,
             line1 = $2,
             line2 = $3,
             district = $4,
             landmark = $5,
             state = $6,
             pincode = $7,
             updated_at = NOW()
         FROM users u
         WHERE a.user_id = $8
           AND a.address_id = $9
           AND u.user_id = a.user_id
           AND u.is_active = true`,
        [
            type,
            line1,
            line2,
            district,
            landmark,
            state,
            pincode,
            userId,
            addressId,
        ]
    );
    return result.rowCount > 0;
};

addressService.delete = async (userId, addressId) => {
    const result = await query(
        `DELETE FROM user_addresses a
         USING users u
         WHERE a.user_id = $1
           AND a.address_id = $2
           AND u.user_id = a.user_id
           AND u.is_active = true`,
        [userId, addressId]
    );
    return result.rowCount > 0;
};

module.exports = addressService;
