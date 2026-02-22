# Inventory Service

Go microservice providing product inventory management, category management, and stock tracking. Designed for the ecommerce backend with MySQL persistence and Kafka integration.

HTTP-based inventory service for product management, category organization, inventory tracking, and real-time updates via Kafka with comprehensive logging and Prometheus metrics.

## Key Features
- Product management (CRUD operations)
- Category and subcategory management
- Inventory tracking and stock management
- Real-time inventory updates via Kafka producer
- HTTP REST API endpoints
- Request/response logging with Zap
- Prometheus metrics collection
- MySQL persistence
- Database health checks
- Transactional operations


## Tech Stack
- Runtime: Go 1.25+
- Framework: HTTP REST
- Database: MySQL 8.0
- Message Queue: Kafka
- Logging: Zap
- Metrics: Prometheus client
- Containerization: Docker


## Installation
```bash
cd inventory-service
go mod download
```

## Environment Variables

Example `conf/dev.yaml`:
```yaml
ENVIRONMENT: "dev"
APP_NAME: "inventory-service"
SQL_CONFIG:
    USERNAME: "testuser"
    PASSWORD: "testpassword"
    HOST: "mysql"
    PORT: "3306"
    DATABASE: "testdb"
    MAX_LIFE_TIME: 3
    MAX_OPEN_CONNS: 10
    MAX_IDLE_CONNS: 10
HTTP_SERVER:
    PORT: 3003
KAFKA_CONFIG:
    BOOTSTRAP_SERVERS: "kafka:19092"
PRODUCT_UPDATE_TOPIC:
    NAME: "product_update"
INVENTORY_UPDATE_TOPIC:
    NAME: "inventory_update"
```

## Running the Application

### Locally
1. Start directly
    ```bash
    make run
    ```

2. Start with live reload (if configured):
    ```bash
    make dev
    ```

3. Build the binary
    ```bash
    make build
    ```

### Using Docker Compose (Recommended)
```bash
docker-compose up --build inventory-service
```

### Run tests
```bash
make test
```


## Usage Examples

### Health Check
```bash
curl http://localhost:3003/health

# 200 {"status":"ok"}
```

### Add Category
```bash
curl --location 'http://localhost:3003/api/category' \
--header 'Content-Type: application/json' \
--data '{
    "name": "mobile"
}'
```

### Get All Category
```bash
curl --location 'http://localhost:3003/api/category'

# 200{
#     "GROCERY": {
#         "id": 1,
#         "name": "GROCERY",
#         "createdAt": "2026-02-22T10:25:30Z"
#     },
#     "MOBILE": {
#         "id": 2,
#         "name": "MOBILE",
#         "createdAt": "2026-02-22T10:27:06Z"
#     }
#     }
# }
```

### Get Category by Id
```bash
curl --location 'http://localhost:3003/api/category/2'
```


### Add Sub-Category
```bash
curl --location 'http://localhost:3003/api/subcategory' \
--header 'Content-Type: application/json' \
--data '{
    "name": "smartphone",
    "category":"mobile"
}'
```

### Get All Sub-Category
```bash
curl --location 'http://localhost:3003/api/subcategory'

# 200 {
#   "PULSE_AND_CEREALS": {
#     "id": 1,
#     "name": "PULSE_AND_CEREALS",
#     "category": "GROCERY",
#     "createdAt": "2026-02-22T10:25:55Z"
#   },
#   "SMARTPHONE": {
#     "id": 2,
#     "name": "SMARTPHONE",
#     "category": "MOBILE",
#     "createdAt": "2026-02-22T10:27:29Z"
#   }
# }
```

### Get Sub-Category by Id
```bash
curl --location 'http://localhost:3003/api/subcategory/2'
```

### Add Product
```bash
curl --location 'http://localhost:3003/api/product' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Nord 42820",
    "description":"4G mobile phone",
    "brand": "Oneplus",
    "category": "mobile",
    "subCategory":"smartphone",
    "imageId":"oneplus_4_phone_image.png",
    "weight":145.9
}'

# 201 {}
```

### List All Products
```bash
curl --location 'http://localhost:3003/api/product'

# 200 [
# {
#     "id":4,
#     "name":"Nord 42820",
#     "description":"4G mobile phone",
#     "brand":"Oneplus",
#     "category":"MOBILE",
#     "subCategory":"SMARTPHONE",
#     "imageId":"oneplus_4_phone_image.png",
#     "weight":145.9,
#     "createdAt":"2026-02-22T10:27:32Z",
#     "updatedAt":"2026-02-22T10:27:32Z"
# }
# ]
```

### Create Product
```bash
curl --location 'http://localhost:3003/api/product/4'

# 201 {
#   "id": 4,
#   "name": "Nord 42820",
#   "description": "4G mobile phone",
#   "brand": "Oneplus",
#   "category": "MOBILE",
#   "subCategory": "SMARTPHONE",
#   "imageId": "oneplus_4_phone_image.png",
#   "weight": 145.9,
#   "createdAt": "2026-02-22T10:27:32Z",
#   "updatedAt": "2026-02-22T10:27:32Z"
# }
```

### Update Product
```bash
curl --location --request PUT 'http://localhost:3003/api/product' \
--header 'Content-Type: application/json' \
--data '{
    "id":11,
    "name": "Nord 4",
    "description":"4G mobile phone",
    "brand": "Oneplus",
    "category": "mobile",
    "subCategory":"smartphone",
    "imageId":"oneplus_4_image.png",
    "weight":100
}'

# 200 {}
```

### Add Inventory
```bash
curl --location 'http://localhost:3003/api/inventory' \
--header 'Content-Type: application/json' \
--data '{
    "productId": 4,
    "sku":6,
    "purchasePrice": 100,
    "salePrice": 125
}'
```

### Get All Inventory
```bash
curl --location 'http://localhost:3003/api/inventory'
```

### Get Inventory by Id
```bash
curl --location 'http://localhost:3003/api/inventory/10'
```

### Update Inventory
```bash
curl --location --request PUT 'http://localhost:3003/api/inventory' \
--header 'Content-Type: application/json' \
--data '{
    "productId": 10,
    "sku":11,
    "purchasePrice": 120,
    "salePrice": 145
}'
```



## API Endpoints

### Categories
- `GET /category` - List all categories
- `GET /category/:id` - Get category by id
- `POST /category` - Create new category

### Subcategories
- `GET /subcategory` - List subcategories for category
- `GET /subcategory/:id` - Get subcategory by ID
- `POST /subcategory` - Create new subcategory

### Products
- `GET /product` - List all products
- `GET /product/:id` - Get product by ID
- `POST /product` - Create new product
- `PUT /product` - Update product details

### Inventory
- `GET /inventory` - List all inventory
- `GET /inventory/:id` - Get product by ID
- `POST /inventory` - Create new product
- `PUT /inventory` - Update product details


### Health & Monitoring
- `GET /health` - Health check endpoint
- `GET /metrics` - Prometheus metrics


## Kafka Topics

### Published Topics
- `product_update` - Product change event
- `inventory_update` - Inventory change events


## Database Schema

### Products Table
```sql
CREATE TABLE `Products` (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `Description` text,
  `Brand` varchar(255) NOT NULL,
  `Category` varchar(255) DEFAULT NULL,
  `SubCategory` varchar(255) DEFAULT NULL,
  `ImageId` varchar(255) DEFAULT NULL,
  `Weight` float DEFAULT NULL,
  `CreatedAt` datetime DEFAULT NULL,
  `UpdatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UniqueId` (`Name`,`Category`,`SubCategory`),
  KEY `Category` (`Category`),
  KEY `SubCategory` (`SubCategory`),
  CONSTRAINT `products_ibfk_1` FOREIGN KEY (`Category`) REFERENCES `Categories` (`Name`),
  CONSTRAINT `products_ibfk_2` FOREIGN KEY (`SubCategory`) REFERENCES `SubCategories` (`Name`)
);
```

### Categories Table
```sql
CREATE TABLE `Categories` (
  `Id` smallint NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `CreatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Name` (`Name`)
);
```

### Subcategories Table
```sql
CREATE TABLE `SubCategories` (
  `Id` smallint NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `Category` varchar(255) DEFAULT NULL,
  `CreatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UniqueId` (`Name`,`Category`),
  KEY `Category` (`Category`),
  CONSTRAINT `subcategories_ibfk_1` FOREIGN KEY (`Category`) REFERENCES `Categories` (`Name`)
);
```

### Inventory Table
```sql
CREATE TABLE `Inventory` (
  `ProductId` bigint NOT NULL,
  `SKU` int DEFAULT NULL,
  `PurchasePrice` int DEFAULT NULL,
  `SalePrice` int DEFAULT NULL,
  `CreatedAt` datetime DEFAULT NULL,
  `UpdatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`ProductId`),
  CONSTRAINT `inventory_ibfk_1` FOREIGN KEY (`ProductId`) REFERENCES `Products` (`Id`)
)
```


## Integration with Other Services

### Search Service
- Publishes product updates to `product_update` Kafka topic
- Search service consumes these events to maintain product index


## Metrics

Metrics are exposed at `/metrics` endpoint. Ensure metrics endpoint is not publicly exposed in production without access control.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
