# Inventory Service

Go HTTP service for categories, subcategories, products, and inventory. It stores data in PostgreSQL and publishes changes to Kafka.

- **Port:** 3003
- **Stack:** Go, PostgreSQL 16, Kafka

## Run

```bash
cd inventory-service
go mod download
make run       # builds and runs ./.bin/app
make build
```

With Compose, from the repository root:

```bash
podman-compose up --build inventory-service
```

## Configuration

Values are in [.env.example](.env.example).

| Variable | Description |
| --- | --- |
| `INVENTORY_SERVICE_PORT` | HTTP port |
| `INVENTORY_SERVICE_READ_TIMEOUT_MS` | Read timeout in milliseconds |
| `INVENTORY_SERVICE_WRITE_TIMEOUT_MS` | Write timeout in milliseconds |
| `SQL_CONFIG_USERNAME` | PostgreSQL username |
| `SQL_CONFIG_PASSWORD_FILE` | Secret file for the PostgreSQL password |
| `SQL_CONFIG_HOST` | PostgreSQL host |
| `SQL_CONFIG_PORT` | PostgreSQL port |
| `SQL_CONFIG_DATABASE` | PostgreSQL database name |
| `SQL_CONFIG_MAX_LIFE_TIME` | Connection max lifetime |
| `SQL_CONFIG_MAX_OPEN_CONNS` | Max open connections |
| `SQL_CONFIG_MAX_IDLE_CONNS` | Max idle connections |
| `KAFKA_CONFIG_SERVER` | Kafka broker |
| `PRODUCT_UPDATE_TOPIC_NAME` | Product update topic |
| `INVENTORY_UPDATE_TOPIC_NAME` | Inventory update topic |

## API

| Method | Path | Description |
| --- | --- | --- |
| POST | `/api/category` | Create category |
| GET | `/api/category` | List categories |
| GET | `/api/category/{id}` | Get category |
| POST | `/api/subcategory` | Create subcategory |
| GET | `/api/subcategory` | List subcategories |
| GET | `/api/subcategory/{id}` | Get subcategory |
| POST | `/api/product` | Create product (publishes `product_update`) |
| PUT | `/api/product` | Update product (publishes `product_update`) |
| GET | `/api/product` | List products |
| GET | `/api/product/{id}` | Get product |
| POST | `/api/inventory` | Create inventory record (publishes `inventory_update`) |
| PUT | `/api/inventory` | Update inventory record (publishes `inventory_update`) |
| GET | `/api/inventory` | List inventory |
| GET | `/api/inventory/{productId}` | Get inventory for a product |

Category and subcategory names are normalized by the service.

## Kafka Topics

Published on writes:
- `product_update` - product create/update (consumed by search service)
- `inventory_update` - inventory create/update

## Database Schema

PostgreSQL database `inventory_db`.

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

CREATE TABLE Categories (
  Id SMALLSERIAL PRIMARY KEY,
  Name VARCHAR(255) UNIQUE NOT NULL,
  CreatedAt TIMESTAMP
);

CREATE TABLE SubCategories (
  Id SMALLSERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Category VARCHAR(255),
  CreatedAt TIMESTAMP,
  UNIQUE (Name, Category),
  FOREIGN KEY (Category) REFERENCES Categories(Name)
);

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

## Health and Metrics

- `GET /health`
- `GET /metrics` (Prometheus; keep it off public networks)

## Tests

```bash
go test ./...
```

## License

MIT
