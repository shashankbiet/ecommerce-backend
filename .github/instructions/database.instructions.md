# Database and data-layer instructions

## Current data stores

- PostgreSQL (`user-service`, `inventory-service`)
- Elasticsearch (`search-service` read model)
- Kafka (`product_update` event stream between inventory and search)

## Change rules

- Treat schema and data contract changes as backward-compatible migrations unless explicitly coordinated.
- Keep SQL/data-access logic inside existing DAO/repository layers.
- Avoid embedding ad hoc SQL in controllers/handlers.
- Keep category/subcategory normalization consistent by reusing existing service helpers in inventory.

## AI-first guidance

- Document every schema/index/contract change in `docs/architecture/database-design.md`.
- Include rollback notes for risky migrations.
- Prefer small, reviewable schema updates with matching tests.