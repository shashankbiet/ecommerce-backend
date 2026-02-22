# User Service

Node.js microservice providing user management, authentication, and account operations. Designed for the ecommerce backend with MongoDB persistence and Prometheus metrics.

Express.js-based user service for registration, authentication, user profiles, and address management with JWT token support and comprehensive API logging.


## Key Features
- User registration with email and phone authentication
- JWT-based authentication and authorization
- User profile management (CRUD operations)
- Address management for users
- Admin role verification and token validation
- Request/response logging with Winston
- Prometheus metrics collection
- MongoDB persistence
- Health check endpoints


## Tech Stack
- Runtime: Node.js
- Framework: Express.js
- Database: MongoDB
- Authentication: JWT (JSON Web Tokens)
- Password Hashing: bcrypt
- Logging: Winston
- Metrics: Prometheus client
- Containerization: Docker


## Installation
```bash
cd user-service
npm install
```

## Environment Variables
Example .env:
```env
PORT=3002
NODE_ENV=development
MONGODB_URL=mongodb://mongodb/userService
ADMIN_VERIFICATION_TOKEN=1234567890
ACCESS_TOKEN_SECRET=9876543210
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
docker-compose up --build user-service
```

### Lint and format
```bash
npm run lint
npm run format
```

## Usage Examples

### Health Check
```bash
curl http://localhost:3002/health

# 200 {"status":"ok"}
```

### Register User
```bash
curl --location 'http://localhost:3002/api/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phoneNumber": "9876543210",
    "email": "test@test.com",
    "password": "test@123",
    "firstName": "John",
    "lastName": "Doe",
    "role": "USER",
    "token": "1234567890"
}'

# 201 {
#   "userId": "507f1f77bcf86cd799439011",
#   "message": "User registered successfully"
# }
```

### Authenticate User
```bash
curl --location 'http://localhost:3002/auth/authenticate' \
--header 'Content-Type: application/json' \
--data '{
    "phoneNumber": "9876543210",
    "password": "test@123"
}'

# 200 {
#   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI1MDdmMWY3N2JjZjg2Y2Q3OTk0MzkwMTEiLCJyb2xlIjoiVVNFUiIsImlhdCI6MTcxNzM4MDAzMCwiZXhwIjoxNzE3NDY2NDMwfQ.kX5c4q9pK2m8nL3oQ1r6sT7uV8wX9yZ0aB1cD2eF3gH"
# }
```

### Get User Profile
```bash
curl --location 'http://localhost:3002/api/user' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI1MDdmMWY3N2JjZjg2Y2Q3OTk0MzkwMTEiLCJyb2xlIjoiVVNFUiIsImlhdCI6MTcxNzM4MDAzMCwiZXhwIjoxNzE3NDY2NDMwfQ.kX5c4q9pK2m8nL3oQ1r6sT7uV8wX9yZ0aB1cD2eF3gH'

# 200 {
#   "userId": "507f1f77bcf86cd799439011",
#   "phoneNumber": "9876543210",
#   "email": "test@test.com",
#   "firstName": "John",
#   "lastName": "Doe",
#   "role": "USER",
#   "isActive": true,
#   "createdDate": "2026-02-22T11:11:47.481Z"
# }
```

### Add User Address
```bash
curl --location 'http://localhost:3002/api/address' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "HOME",
    "line1": "123 MG Road",
    "line2": "Apartment 4B",
    "district": "Bangalore",
    "landmark": "Near Indiranagar Tech Park",
    "state": "Karnataka",
    "pincode": 560038
}'

# 201 {
#   "addressId": "507f1f77bcf86cd799439012",
#   "message": "Address added successfully"
# }
```

### Get User Addresses
```bash
curl --location 'http://localhost:3002/api/address' \
--header 'Authorization: Bearer <token>'

# 200 [
#   {
#     "addressId": "507f1f77bcf86cd799439012",
#     "type": "HOME",
#     "line1": "123 MG Road",
#     "line2": "Apartment 4B",
#     "district": "Bangalore",
#     "landmark": "Near Indiranagar Tech Park",
#     "state": "Karnataka",
#     "pincode": 560038
#   }
# ]
```

## API Endpoints
- `POST /api/user/register` - Register new user
- `POST /auth/authenticate` - User authenticate
- `GET /api/user` - Get user profile (requires auth)
- `PUT /api/user` - Update user profile (requires auth)
- `DELETE /api/user` - Delete user profile (requires auth)
- `POST /api/address` - Add new address (requires auth)
- `GET /api/address` - Get user addresses (requires auth)
- `PUT /api/address/:id` - Update address (requires auth)
- `DELETE /api/address/:id` - Delete address (requires auth)
- `GET /health` - Health check
- `GET /metrics` - Metrics

## Metrics

Expose Prometheus metrics at `/metrics`. Ensure this is not publicly exposed in production without access control.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
