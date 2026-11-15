# Backend implementation instructions

## Purpose
This repository is a learning-first microservices project and is being upgraded to an AI-first engineering workflow.

## How to implement backend changes

- Keep changes service-local unless a cross-service contract requires updates.
- For Node.js services (`api-gateway`, `user-service`):
  - Keep router/controller/service separation.
  - Reuse existing middleware and shared error/status patterns.
  - Use PostgreSQL access through the existing `handler/postgres` abstraction in `user-service`.
- For Go services (`inventory-service`, `search-service`):
  - Keep handler → service → DAO boundaries.
  - Keep DB/search logic in DAO implementations.
  - Preserve constructor-based dependency injection.
- When changing event or RPC contracts:
  - Update producer + consumer together.
  - Regenerate protobuf stubs when `search-service/proto/search/search.proto` changes.
  - Document contract updates under `docs/architecture/`.

## AI-first expectations

- Include clear docs updates for behavior changes.
- Prefer deterministic and testable logic that AI tools can reason about.
- Add focused tests for modified behavior (unit first, integration when needed).