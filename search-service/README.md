# Search Service

Go gRPC service for product search. It consumes product updates from Kafka and indexes them in Elasticsearch.

- **Ports:** 9004 (gRPC), 9104 (metrics)
- **Stack:** Go, gRPC, Elasticsearch, Kafka

## Run

```bash
cd search-service
go mod download
make run       # builds and runs ./.bin/app
make build
```

With Compose, from the repository root:

```bash
podman-compose up --build search-service
```

## Configuration

Values are in [.env.example](.env.example).

| Variable | Description |
| --- | --- |
| `SEARCH_SERVICE_GRPC_PORT` | gRPC port |
| `ELASTIC_SEARCH_ADDRESS` | Elasticsearch URL |
| `ELASTIC_SEARCH_USERNAME` | Elasticsearch username |
| `ELASTIC_SEARCH_PASSWORD_FILE` | Secret file for the Elasticsearch password |
| `ELASTIC_SEARCH_INDEX_NAME` | Product index name |
| `KAFKA_CONFIG_SERVER` | Kafka broker |
| `KAFKA_CONFIG_SEARCH_SERVICE_GROUP_ID` | Kafka consumer group ID |
| `PRODUCT_UPDATE_TOPIC_NAME` | Topic consumed for product changes |

## API

gRPC service `searchpb.SearchService`, defined in [search.proto](proto/search/search.proto).

| RPC | Description |
| --- | --- |
| `GetProduct` | Search products by `category`, `sub_category`, and `keywords` |

Example:

```bash
grpcurl -plaintext -emit-defaults \
  -proto search-service/proto/search/search.proto \
  -import-path . \
  -d '{"category":"Electronics","keywords":"keyboard","sub_category":"Accessories"}' \
  localhost:9004 searchpb.SearchService.GetProduct
```

## Kafka Topics

Subscribed (consumer group `search_service`):
- `product_update` - product change events, indexed into Elasticsearch
- `inventory_update` - inventory change events

## Health and Metrics

- Health: `search-service healthcheck` checks `http://127.0.0.1:9104/metrics`
- Metrics: `GET :9104/metrics` (Prometheus; keep it off public networks)

## Tests

```bash
go test ./...
```

## License

MIT
