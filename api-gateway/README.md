# API Gateway

Express gateway and single entry point for client apps. It validates bearer tokens, proxies user requests to the user service, and forwards trusted identity headers.

- **Port:** 3001
- **Stack:** Node.js, Express

## Run

```bash
cd api-gateway
npm install
npm start      # production mode
npm run dev    # nodemon, development mode
```

With Compose, from the repository root:

```bash
podman-compose up --build api-gateway
```

## Configuration

Values are in [.env.example](.env.example).

| Variable | Description |
| --- | --- |
| `API_GATEWAY_PORT` | HTTP port |
| `API_GATEWAY_ACCESS_TOKEN_SECRET_FILE` | Secret file used to verify bearer tokens |
| `USER_SERVICE_TARGET` | User service base URL |
| `USER_SERVICE_USER_PATH` | Proxied user path (`/api/user`) |
| `USER_SERVICE_AUTH_PATH` | Proxied auth path (`/api/auth`) |

## API

| Method | Path | Auth | Description |
| --- | --- | --- | --- |
| POST | `/auth/login` | No | Log in and return a token |
| POST | `/api/user/register` | No | Register a user |
| ALL | `/api/user/*` | Yes | Proxied to the user service |

The user service documents its endpoints in [user-service](../user-service/README.md).

## Health and Metrics

- `GET /health`
- `GET /metrics` (Prometheus; keep it off public networks)

## Tests and Lint

```bash
npm test
npm run lint
npm run format
```

## License

MIT
