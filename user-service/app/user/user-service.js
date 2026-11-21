const hashing = require("../../util/hashing");
const { adminVerificationToken } = require("../../config");
const { dbResponseTimeHistogram } = require("../../util/metrics");
const { query } = require("../../handler/postgres");
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
        let {
            phone_number: phoneNumber,
            email,
            password,
            role,
            first_name: firstName,
            last_name: lastName,
        } = body;
        const dbUser = await query(
            `SELECT user_id FROM users
             WHERE email = $1 OR phone_number = $2`,
            [email, phoneNumber]
        );
        if (dbUser.rowCount > 0) {
            timer({ ...metricsLabels, success: false });
            return null;
        }
        password = await hashing.encrypt(password);
        const insertUser = await query(
            `INSERT INTO users (phone_number, email, password, role, first_name, last_name)
             VALUES ($1, $2, $3, $4, $5, $6)
             RETURNING user_id`,
            [phoneNumber, email, password, role, firstName, lastName]
        );
        timer({ ...metricsLabels, success: true });
        return Number(insertUser.rows[0].user_id);
    } catch (err) {
        timer({ ...metricsLabels, success: false });
        if (err.code === "23505") {
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
        const result = await query(
            `SELECT user_id, phone_number, email, role, first_name, last_name,
                    created_at, updated_at, is_active
             FROM users
             WHERE user_id = $1 AND is_active = true`,
            [userId]
        );
        timer({ ...metricsLabels, success: true });
        if (result.rowCount > 0) {
            const user = result.rows[0];
            return {
                user_id: Number(user.user_id),
                phone_number: user.phone_number,
                email: user.email,
                role: user.role,
                first_name: user.first_name,
                last_name: user.last_name,
                created_at: user.created_at,
                updated_at: user.updated_at,
                is_active: user.is_active,
            };
        }
        return null;
    } catch (err) {
        timer({ ...metricsLabels, success: false });
        throw err;
    }
};

userService.update = async (userId, body) => {
    const { first_name: firstName, last_name: lastName } = body;
    const result = await query(
        `UPDATE users
         SET first_name = $1,
             last_name = $2,
             updated_at = NOW()
         WHERE user_id = $3 AND is_active = true`,
        [firstName, lastName, userId]
    );
    return result.rowCount > 0;
};

userService.delete = async (userId) => {
    const result = await query(
        `UPDATE users
         SET is_active = false,
             updated_at = NOW()
         WHERE user_id = $1 AND is_active = true`,
        [userId]
    );
    return result.rowCount > 0;
};

userService.updatePassword = async (userId, currentPassword, newPassword) => {
    const userResult = await query(
        `SELECT password
         FROM users
         WHERE user_id = $1 AND is_active = true`,
        [userId]
    );
    if (userResult.rowCount === 0) {
        return null;
    }
    const auth = await hashing.compare(
        currentPassword,
        userResult.rows[0].password
    );
    if (!auth) {
        return false;
    }
    await query(
        `UPDATE users
         SET password = $1,
             updated_at = NOW()
         WHERE user_id = $2 AND is_active = true`,
        [await hashing.encrypt(newPassword), userId]
    );
    return true;
};

module.exports = userService;
