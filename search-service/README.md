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

### Search Products
```bash
grpcurl \
	-plaintext \
	-emit-defaults \
	-proto 'search-service/proto/search/search.proto' \
	-import-path '.' \
	-d '{"category":"Electronics","keywords":"keyboard","sub_category":"Accessories"}' \
	'localhost:9004' \
	searchpb.SearchService.GetProduct

# {
#     "products": [
#         {
#             "id": 4,
#             "name": "Mechanical Keyboard",
#             "description": "Mechanical Keyboard, a demo accessories item from Electronics.",
#             "brand": "Northstar",
#             "category": "Electronics",
#             "subCategory": "Accessories"
#         }
#     ],
#     "keywords": "keyboard",
#     "category": "Electronics",
#     "sub_category": "Accessories",
#     "total_results": 1
# }
```


## Kafka Topics

### Subscribed Topics
- `product_update` - Product change event
- `inventory_update` - Inventory change events

## Metrics

Prometheus metrics are exposed over HTTP at port `9104` on the service network.
This includes gRPC server, Go runtime, and process metrics. In the local
Compose setup, Prometheus scrapes this endpoint when the `monitoring` profile
is enabled. Keep the endpoint on a trusted network and do not expose it
publicly without access control.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
