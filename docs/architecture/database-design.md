# Database design

## Primary stores

### PostgreSQL

- **user_db / user-service**
  - users
  - addresses
  - auth-related data
- **inventory_db / inventory-service**
  - categories
  - subcategories
  - products
  - inventory

### Elasticsearch

- Search read model for product discovery in `search-service`.
- Fed by Kafka `product_update` events.

## Messaging backbone

- Kafka topic `product_update` is the canonical stream for product indexing changes from inventory to search.

## Design notes

- User and inventory domains are separated at service and database boundaries.
- Search remains eventually consistent with inventory by design.

## AI-first data practices

- Document schema and index changes with rationale.
- Keep migration steps reversible where practical.
- Validate event payload compatibility before deployment.