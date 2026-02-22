# API Gateway

Lightweight HTTP gateway that routes and aggregates requests to backend services (User Service). Designed for local development and as the single entry point for client apps.

Node.js Express-based gateway providing routing, request proxying, basic auth hooks, and response aggregation for the ecommerce backend.


## Key Features
- Route proxying to internal services (users)
- Centralized health & readiness endpoints
- Simple auth middleware hooks (extensible)
- Request/response logging
- Basic metrics endpoint compatible with Prometheus
- Config-driven per-environment settings


## Tech Stack
- Runtime: Node.js
- Framework: Express.js
- Database: MongoDB
- Containerization: Docker


## Installation
```bash
cd api-gateway
npm install
```


## Environment Variables
Example .env:
```env
PORT=3001
ACCESS_TOKEN_SECRET=9876543210
USER_SERVICE_TARGET=http://user-service:3002
```


## Running the Application

### Locally
1. Start directly
    ```bash
    npm start
    ```

2. Start with nodemon for hot reload (if configured):
    ```bash
    npm run dev
    ```

### Using Docker Compose (Recommended)
```bash
docker-compose up --build api-gateway
```

### Lint and format
```bash
npm run lint
npm run format
```

## Usage Examples

### Health
```bash
curl http://localhost:3001/health

# 200 {"status":"ok"}
```

### Create User
```bash
curl --location 'http://localhost:3001/api/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phoneNumber": "9876543210",
    "email": "test@test.com",
    "password": "test@",
    "firstName": "test",
    "lastName": "test",
    "role": "ADMIN",
    "token": "1234567890"
}'

# 200 {}
```

### Login User
```bash
curl --location 'http://localhost:3001/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "phoneNumber":9876543210,
    "password":"test"
}'

# 200 {eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEwMDYsInJvbGUiOiJBRE1JTiIsImlhdCI6MTc3MTc1ODczMywiZXhwIjoxNzcxODQ1MTMzfQ.47CAst2SnDU9QIImvl-G0zBbWYyUyzy6bQes2-s_M_0}
```

### Get User
```bash
curl --location 'http://localhost:3001/api/user' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEwMDMsInJvbGUiOiJBRE1JTiIsImlhdCI6MTc3MTc1ODU0MCwiZXhwIjoxNzcxODQ0OTQwfQ.oPg7B3HfBv8G-W7LOWL4FJX_35ppv0937CZREEPChH8'

# 200 {
#   "userId": 1006,
#   "phoneNumber": "9876543210",
#   "email": "test@test.com",
#   "role": "ADMIN",
#   "firstName": "test",
#   "lastName": "test",
#   "isActive": true,
#   "createdDate": "2026-02-22T11:11:47.481Z"
# }
```

### Additional Endpoints
Similarly, you can access other user service APIs as documented in the [user-service README](../user-service/README.md).



## Metrics

Expose Prometheus metrics at /metrics (configure METRICS_PATH). Ensure METRICS_PATH is not publicly exposed in production without access control.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.