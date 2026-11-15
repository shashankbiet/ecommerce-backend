# API design

## External API boundary

`api-gateway` is the external HTTP boundary.

- Validates bearer token for protected routes.
- Preserves/creates `correlation-id`.
- Forwards trusted identity headers (`user-id`, `user-role`) downstream.
- Removes raw `authorization` header before proxying to user-service endpoints.

## Internal service interfaces

- **user-service**: REST-style endpoints for auth, profile, and addresses.
- **inventory-service**: HTTP handlers over service + DAO abstractions.
- **search-service**: gRPC interface defined in `search-service/proto/search/search.proto`.

## Contract governance

- Any API contract change must update:
  1. service implementation
  2. consumers/callers
  3. architecture docs
  4. tests

## AI-first API principles

- Keep request/response schemas explicit and stable.
- Prefer additive changes over breaking changes.
- Include example payloads and failure modes in follow-up docs/PR descriptions.