# API design

## External API boundary

`api-gateway` is the external HTTP boundary.

- Validates bearer token for protected routes.
- Preserves/creates `correlation-id`.
- Forwards trusted identity headers (`user-id`, `user-role`) downstream.
- Removes raw `authorization` header before proxying to user-service endpoints.

## Field naming

All HTTP request and response JSON fields, gRPC JSON fields, and Kafka payload fields use `snake_case`. Examples include `user_id`, `phone_number`, `sub_category`, `image_url`, `created_at`, and `updated_at`.

The services accept legacy camelCase request and event fields at their boundaries and normalize them internally. Responses, newly issued JWT claims, and newly published Kafka events always use snake_case. New integrations must use only snake_case.

## Internal service interfaces

- **user-service**: REST-style endpoints for auth, profile, and addresses.
- **inventory-service**: HTTP handlers over service + DAO abstractions.
- **search-service**: gRPC interface defined in `search-service/proto/search/search.proto`.

`GET /api/brand` is the inventory-service read endpoint for frontend brand
selectors. It returns a JSON array sorted alphabetically by brand name:

```json
[
  {
    "id": 1,
    "name": "Demo Electronics",
    "created_at": "2026-11-22T00:00:00Z"
  }
]
```

## Contract governance

- Any API contract change must update:
  1. service implementation
  2. consumers/callers
  3. architecture docs
  4. tests

## Product update event

`inventory-service` publishes the complete persisted product model to `product_update` after successful create or update. `search-service` validates `id` and uses the event for an Elasticsearch `doc_as_upsert` operation. The public gRPC product response preserves legacy string category fields and adds the ID and product fields; multi-word protobuf fields have explicit snake_case JSON names.

Product `status` is the numeric enum `0` (`DRAFT`), `1` (`ACTIVE`), or `2`
(`ARCHIVED`). Producers, consumers, and HTTP create/update validation reject any
other value. Omitted status values default to `DRAFT`.

## AI-first API principles

- Keep request/response schemas explicit and stable.
- Prefer additive changes over breaking changes.
- Include example payloads and failure modes in follow-up docs/PR descriptions.