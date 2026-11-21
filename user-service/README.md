# User Service

Node.js (Express) service for registration, authentication, user profiles, and addresses, backed by PostgreSQL.

- **Port:** 3002
- **Stack:** Node.js, Express, PostgreSQL, JWT

## Run

```bash
cd user-service
npm install
npm start      # production mode
npm run dev    # nodemon, development mode
```

With Compose, from the repository root:

```bash
podman-compose up --build user-service
```

## Configuration

Values are in [.env.example](.env.example).

| Variable | Description |
| --- | --- |
| `USER_SERVICE_PORT` | HTTP port |
| `USER_SERVICE_ADMIN_VERIFICATION_TOKEN_FILE` | Secret file for the admin verification token |
| `SQL_CONFIG_USERNAME` | PostgreSQL username |
| `SQL_CONFIG_PASSWORD_FILE` | Secret file for the PostgreSQL password |
| `SQL_CONFIG_HOST` | PostgreSQL host |
| `SQL_CONFIG_PORT` | PostgreSQL port |
| `SQL_CONFIG_DATABASE` | PostgreSQL database name |
| `SQL_CONFIG_MAX_CLIENTS` | Max PostgreSQL clients |

## Database

PostgreSQL database `user_db` (set by `SQL_CONFIG_DATABASE`). Stores user profiles, addresses, and authentication data.

## API

Paths are served by this service. The gateway forwards them and enforces the bearer token on everything except register and login.

| Method | Path | Auth | Description |
| --- | --- | --- | --- |
| POST | `/api/user/register` | No | Register a user |
| POST | `/api/auth/authenticate` | No | Log in and return a token |
| GET | `/api/user` | Yes | Get profile |
| PUT | `/api/user` | Yes | Update profile |
| DELETE | `/api/user` | Yes | Delete profile |
| PUT | `/api/user/password` | Yes | Change password |
| POST | `/api/user/address` | Yes | Add address |
| GET | `/api/user/address` | Yes | List addresses |
| GET | `/api/user/address/:address_id` | Yes | Get address |
| PUT | `/api/user/address/:address_id` | Yes | Update address |
| DELETE | `/api/user/address/:address_id` | Yes | Delete address |

All request and response JSON fields use `snake_case`, including `phone_number`, `first_name`, `last_name`, `user_id`, `address_id`, `current_password`, and `new_password`. Legacy camelCase request fields are accepted and normalized at the service boundary; responses always use snake_case.

## Health and Metrics

- `GET /api/user/health`
- `GET /api/user/metrics` (Prometheus; keep it off public networks)

## Tests and Lint

```bash
npm test
npm run lint
npm run format
```

## License

MIT
