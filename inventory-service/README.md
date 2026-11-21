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
| GET | `/api/brand` | List brands alphabetically |
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
| GET | `/api/inventory/{product_id}` | Get inventory for a product |
| POST | `/api/keyword` | Create keyword |
| GET | `/api/keyword` | List keywords |
| GET | `/api/keyword/{id}` | Get keyword |
| POST | `/api/product-keyword` | Associate a keyword with a product |
| GET | `/api/product-keyword/product/{product_id}` | List keyword associations for a product |
| GET | `/api/product-keyword/keyword/{keyword_id}` | List product associations for a keyword |
| DELETE | `/api/product-keyword/{product_id}/{keyword_id}` | Delete a product-keyword association |

Brand listing returns a JSON array containing `id`, `name`, and `created_at`.
The list is ordered by brand name to provide stable frontend rendering.

Category and subcategory names are normalized by the service where names are user-entered.

## Kafka Topics

Published on writes:
- `product_update` - product create/update (consumed by search service)
- `inventory_update` - inventory create/update

Product create/update events use the persisted product identifier and timestamps so search consumers can upsert the complete current model.

## Database Schema

PostgreSQL database `inventory_db`.

```sql
CREATE TABLE categories (
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE brands (
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE sub_categories (
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  category_id SMALLINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (name, category_id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE products (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  slug VARCHAR(255),
  brand_id SMALLINT NOT NULL,
  category_id SMALLINT NOT NULL,
  sub_category_id SMALLINT NOT NULL,
  image_url VARCHAR(255),
  status BIGINT NOT NULL DEFAULT 0 CHECK (status IN (0, 1, 2)),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ,
  UNIQUE (name, category_id, sub_category_id),
  FOREIGN KEY (brand_id) REFERENCES brands(id),
  FOREIGN KEY (category_id) REFERENCES categories(id),
  FOREIGN KEY (sub_category_id) REFERENCES sub_categories(id)
);

CREATE TABLE keywords (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE product_keywords (
  product_id BIGINT NOT NULL,
  keyword_id BIGINT NOT NULL,
  PRIMARY KEY (product_id, keyword_id),
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (keyword_id) REFERENCES keywords(id)
);

CREATE TABLE inventory (
  product_id BIGINT NOT NULL PRIMARY KEY,
  sku INT,
  mrp INT,
  purchase_price INT,
  sale_price INT,
  quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
```

All request, response, and Kafka JSON fields use `snake_case`. Product payloads use `slug`, `image_url`, `status`, and `deleted_at`, with `brand`, `category`, and `sub_category` represented as numeric IDs. Inventory payloads include `mrp` in create, update, and read responses. Legacy camelCase request fields remain accepted during migration.

Product status is an enum represented as an integer: `0` = `DRAFT`, `1` = `ACTIVE`, and `2` = `ARCHIVED`. Create and update requests reject every other value. The default value is `DRAFT` when status is omitted.

## Health and Metrics

- `GET /health`
- `GET /metrics` (Prometheus; keep it off public networks)

## Generate easyjson

Structs in `app/models` marked with `//easyjson:json` get generated `*_easyjson.go` marshalers. Do not edit those files; change the struct and regenerate.

```bash
cd inventory-service
go install github.com/mailru/easyjson/easyjson@v0.9.1   # one-time; matches go.mod
~/go/bin/easyjson app/models/product_keyword.go          # one file
~/go/bin/easyjson -all app/models/keyword.go             # every struct in the file
```

`~/go/bin` must be on your `PATH` to run `easyjson` without the full path. The generator compiles the package first, so the package must build before generating.

## Tests

```bash
go test ./...
```

## License

MIT
