# Inventory Service

Go microservice providing product inventory management, category management, and stock tracking. Designed for the ecommerce backend with PostgreSQL persistence and Kafka integration.

HTTP-based inventory service for product management, category organization, inventory tracking, and real-time updates via Kafka with comprehensive logging and Prometheus metrics.

## Key Features
- Product management (CRUD operations)
- Category and subcategory management
- Inventory tracking and stock management
- Real-time inventory updates via Kafka producer
- HTTP REST API endpoints
- Request/response logging with Zap
- Prometheus metrics collection
- PostgreSQL persistence
- Database health checks
- Transactional operations


## Tech Stack
- Runtime: Go 1.26+
- Framework: HTTP REST
- Database: PostgreSQL 16
- Message Queue: Kafka
- Logging: Zap
- Metrics: Prometheus client
- Containerization: Docker


## Installation
```bash
cd inventory-service
go mod download
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
CREATE TABLE Products (
  Id BIGSERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Description TEXT,
  Brand VARCHAR(255) NOT NULL,
  Category VARCHAR(255),
  SubCategory VARCHAR(255),
  ImageId VARCHAR(255),
  Weight FLOAT,
  CreatedAt TIMESTAMP,
  UpdatedAt TIMESTAMP,
  UNIQUE (Name, Category, SubCategory),
  FOREIGN KEY (Category) REFERENCES Categories(Name),
  FOREIGN KEY (SubCategory, Category) REFERENCES SubCategories(Name, Category)
);
```

### Categories Table
```sql
CREATE TABLE Categories (
  Id SMALLSERIAL PRIMARY KEY,
  Name VARCHAR(255) UNIQUE NOT NULL,
  CreatedAt TIMESTAMP
);
```

### Subcategories Table
```sql
CREATE TABLE SubCategories (
  Id SMALLSERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Category VARCHAR(255),
  CreatedAt TIMESTAMP,
  UNIQUE (Name, Category),
  FOREIGN KEY (Category) REFERENCES Categories(Name)
);
```

### Inventory Table
```sql
CREATE TABLE Inventory (
  ProductId BIGINT NOT NULL PRIMARY KEY,
  SKU INT,
  PurchasePrice INT,
  SalePrice INT,
  CreatedAt TIMESTAMP,
  UpdatedAt TIMESTAMP,
  FOREIGN KEY (ProductId) REFERENCES Products(Id)
);
```


## Integration with Other Services

### Search Service
- Publishes product updates to `product_update` Kafka topic
- Search service consumes these events to maintain product index


## Metrics

Prometheus metrics are exposed at `/metrics`. The endpoint includes Go runtime
and process metrics. In the local Compose setup, Prometheus scrapes this
endpoint when the `monitoring` profile is enabled. Keep it on a trusted network
and do not expose it publicly without access control.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
