const { Pool } = require("pg");
const { postgresConfig } = require("../config");
const logger = require("./logger");

const pool = new Pool({
    host: postgresConfig.host,
    port: postgresConfig.port,
    database: postgresConfig.database,
    user: postgresConfig.user,
    password: postgresConfig.password,
    max: postgresConfig.maxClients,
});

const initializeSchema = async () => {
    await pool.query("CREATE EXTENSION IF NOT EXISTS pgcrypto");
    await pool.query(`
        CREATE SEQUENCE IF NOT EXISTS user_id_seq
            START WITH 1001
            MINVALUE 1001
    `);
    await pool.query(`
        CREATE TABLE IF NOT EXISTS users (
            user_id BIGINT PRIMARY KEY DEFAULT nextval('user_id_seq'),
            phone_number VARCHAR(10) NOT NULL UNIQUE,
            email VARCHAR(255) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL,
            role VARCHAR(32) NOT NULL,
            first_name VARCHAR(128) NOT NULL,
            last_name VARCHAR(128),
            created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
            last_modified_date TIMESTAMPTZ,
            is_active BOOLEAN NOT NULL DEFAULT TRUE
        )
    `);
    await pool.query(`
        ALTER SEQUENCE user_id_seq
            OWNED BY users.user_id
    `);
    await pool.query(`
        CREATE TABLE IF NOT EXISTS user_addresses (
            address_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
            type VARCHAR(32) NOT NULL,
            line1 VARCHAR(255) NOT NULL,
            line2 VARCHAR(255),
            district VARCHAR(128) NOT NULL,
            landmark VARCHAR(255) NOT NULL,
            state VARCHAR(128) NOT NULL,
            pincode INTEGER NOT NULL,
            created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
            last_modified_date TIMESTAMPTZ
        )
    `);
    await pool.query(`
        SELECT setval(
            'user_id_seq',
            GREATEST((SELECT COALESCE(MAX(user_id), 1001) FROM users), 1001),
            EXISTS (SELECT 1 FROM users)
        )
    `);
};

const connect = async () => {
    try {
        await pool.query("SELECT 1");
        await initializeSchema();
        logger.info("POSTGRES CONNECTED");
    } catch (err) {
        logger.error(err);
        throw err;
    }
};

const query = (text, params) => pool.query(text, params);

module.exports = {
    connect,
    query,
    pool,
};
