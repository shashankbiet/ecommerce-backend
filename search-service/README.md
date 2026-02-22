# Search Service

Go microservice providing product search, filtering, and indexing capabilities. Designed for the ecommerce backend with Elasticsearch persistence and Prometheus metrics.

gRPC-based search service for product search, advanced filtering, index management, and real-time search updates via Kafka with comprehensive logging and metrics.

## Key Features
- Full-text product search with Elasticsearch
- Real-time index updates via Kafka consumer
- gRPC API endpoints for high-performance search
- Request/response logging with Zap
- Prometheus metrics collection
- Elasticsearch persistence
- Health check endpoints


## Tech Stack
- Runtime: Go 1.25+
- Protocol: gRPC with Protocol Buffers
- Search Engine: Elasticsearch
- Message Queue: Kafka
- Logging: Zap
- Metrics: Prometheus client
- Containerization: Docker


## Installation
```bash
cd search-service
go mod download
```

## Environment Variables

Example `conf/dev.yaml`:
```yaml
ENVIRONMENT: "dev"
APP_NAME: "search-service"
GRPC_SERVER:
    PORT: 9004
ELASTIC_SEARCH:
    ADDRESS: http://elasticsearch:9200
    INDEX_NAME: searchindex
KAFKA_CONFIG:
    SERVERS: "kafka:19092"
    GROUP_ID: "inventory_service"
PRODUCT_UPDATE_TOPIC:
    NAME: "product_update"
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
docker-compose up --build search-service


```
## Usage Examples

### Get Product by ID (Currently Implemented)
```bash
grpcurl \
	-plaintext \
	-emit-defaults \
	-proto 'https://github.com/shashankbiet/ecommerce-backend/search-service/proto/search/search.proto' \
	-import-path 'https://github.com/shashankbiet/ecommerce-backend/search-service/proto/search' \
	-d '{"category":"mobile","keywords":"n","sub_category":"smartphone"}' \
	'localhost:9004' \
	searchpb.SearchService.GetProduct

# {
#     "products": [
#         {
#             "id": 0,
#             "name": "Nord 42820",
#             "description": "4G mobile phone",
#             "brand": "Oneplus",
#             "category": "MOBILE",
#             "subCategory": "SMARTPHONE"
#         }
#     ],
#     "keywords": "n",
#     "category": "mobile",
#     "sub_category": "smartphone",
#     "total_results": 1
# }
```


## Kafka Topics

### Subscribed Topics
- `product_update` - Product change event
- `inventory_update` - Inventory change events

## Metrics

Expose Prometheus metrics via gRPC endpoint or HTTP. Ensure metrics endpoint is not publicly exposed in production without access control.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
