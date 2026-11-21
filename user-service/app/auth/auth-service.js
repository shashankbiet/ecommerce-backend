const hashing = require("../../util/hashing");
const { query } = require("../../handler/postgres");
const authService = {};

authService.authenticate = async (email, phoneNumber, plainPass) => {
    const orConditions = [];
    if (email) orConditions.push({ email });
    if (phoneNumber) orConditions.push({ phone_number: phoneNumber });
    if (orConditions.length === 0) {
        return null;
    }

    const clauses = [];
    const values = [];
    if (email) {
        clauses.push(`email = $${values.length + 1}`);
        values.push(email);
    }
    if (phoneNumber) {
        clauses.push(`phone_number = $${values.length + 1}`);
        values.push(phoneNumber);
    }
    const result = await query(
        `SELECT user_id, role, password
         FROM users
         WHERE is_active = true AND (${clauses.join(" OR ")})
         LIMIT 1`,
        values
    );
    const user = result.rows[0];
    if (user && (await hashing.compare(plainPass, user.password))) {
        return { user_id: Number(user.user_id), role: user.role };
    }
    return null;
};

module.exports = authService;
