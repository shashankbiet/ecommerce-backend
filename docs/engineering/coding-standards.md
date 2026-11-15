# Coding standards

## General

- Keep changes small, focused, and easy to review.
- Follow existing naming and folder structure.
- Prefer explicit errors over silent fallbacks.

## Node.js services (`api-gateway`, `user-service`)

- Use CommonJS style already present in the codebase.
- Keep router/controller/service responsibilities separate.
- Reuse shared middleware/utilities for auth, correlation IDs, and status handling.

## Go services (`inventory-service`, `search-service`)

- Keep handler → service → DAO layering.
- Use constructor-based dependency injection.
- Keep datastore specifics inside DAO packages.

## Documentation and AI-first practices

- Update architecture/engineering docs whenever behavior or contracts change.
- Add concise comments only for non-obvious logic.
- Write code and tests so AI tools can infer intent from structure and names.