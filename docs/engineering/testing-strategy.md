# Testing strategy

## Goals

- Keep the project reliable as a learning environment.
- Support AI-assisted changes with fast, deterministic feedback.

## Test layers

- **Unit tests** for service and helper logic (default).
- **Integration-style tests** for DAO and boundary behavior where risk is higher.
- **Contract checks** for gRPC/protobuf and event payload compatibility when interfaces change.

## Commands

- `cd api-gateway && npm test`
- `cd user-service && npm test`
- `cd inventory-service && go test ./...`
- `cd search-service && go test ./...`

## Scope policy

- Add tests for changed behavior and bug fixes.
- Include meaningful edge cases (invalid input, empty results, downstream failure).
- Keep fixtures minimal and easy to understand for contributors learning the stack.