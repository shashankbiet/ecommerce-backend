# Database Design

## Primary Data Stores

### PostgreSQL

- **user_db / user-service**
  - `users`
  - `user_addresses`
  - authentication-related data
- **inventory_db / inventory-service**
  - `categories`
  - `brands`
  - `sub_categories`
  - `products`
  - `keywords`
  - `product_keywords`
  - `inventory`

Current inventory-service schema expectations:
- `products` includes `slug`, `image_url`, `status`, and `deleted_at`.
- `products.status` is constrained to the API enum values `0` (`DRAFT`), `1`
  (`ACTIVE`), and `2` (`ARCHIVED`). New and updated rows cannot use other
  values.
- `products.brand_id`, `products.category_id`, and
  `products.sub_category_id` are ID-based foreign-key references to `brands`,
  `categories`, and `sub_categories`, respectively.
- `inventory` includes `mrp`, `sku`, `purchase_price`, `sale_price`, and `quantity`.
- Demo seed products in `scripts/seed` use the same ID-based product shape in
  PostgreSQL and Elasticsearch, including `slug`, `image_url`, `status`, and
  timestamps. Seed integrity is checked by `python3 scripts/seed/validate.py`.
- PostgreSQL identifiers use lowercase snake_case. Foreign-key columns use
  `_id`, and timestamps use `TIMESTAMPTZ` with `created_at`, `updated_at`, and
  `deleted_at` naming. API and event payload fields also use snake_case.

The status check is added as `NOT VALID` for existing databases so deployment
does not fail if legacy invalid rows already exist; PostgreSQL still enforces
the constraint for all new inserts and updates. Existing invalid rows should be
reviewed and corrected separately. Rollback removes the `products_status_check`
constraint with `ALTER TABLE products DROP CONSTRAINT products_status_check`.

The brand foreign key is created with the `products_brand_id_fkey` constraint.
For an existing database, it is added as `NOT VALID` so startup does not fail on
legacy product rows whose brand IDs predate the `brands` table. New product
inserts and updates must reference an existing brand. Existing invalid rows
should be reconciled with matching brand records before validating the
constraint. Rollback removes the foreign key with
`ALTER TABLE products DROP CONSTRAINT products_brand_id_fkey`.

### Elasticsearch

- Used by `search-service` as the product discovery read model.
- Updated through Kafka-driven indexing flows.

## Messaging Backbone

- `product_update` is the canonical topic for product indexing changes consumed by `search-service`.
- `inventory_update` carries inventory record updates from `inventory-service`.

## Design Notes

- User and inventory domains remain isolated at service and database boundaries.
- Search is intentionally eventually consistent with inventory writes.

## AI-First Data Practices

- Document schema/index/contract changes with rationale and impact.
- Prefer backward-compatible, reversible migrations.
- Validate event payload compatibility before rollout.