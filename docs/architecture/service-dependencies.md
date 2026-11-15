# Service dependencies

## Runtime dependency graph

- `api-gateway` → `user-service`
- `inventory-service` → PostgreSQL (`inventory_db`)
- `user-service` → PostgreSQL (`user_db`)
- `inventory-service` → Kafka (`product_update` producer)
- `search-service` → Kafka (`product_update` consumer)
- `search-service` → Elasticsearch

## Data flow summary

1. Client traffic enters through `api-gateway`.
2. Gateway authenticates and forwards trusted identity context.
3. Inventory updates publish product events to Kafka.
4. Search consumes events and updates Elasticsearch index.
5. Search API serves gRPC queries from the indexed read model.

## Coupling guidance

- Keep synchronous dependencies minimal.
- Prefer event-driven updates for cross-domain propagation.
- Preserve correlation and identity headers through trusted boundaries.