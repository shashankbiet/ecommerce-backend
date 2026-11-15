# Repository guide

## Project context

- This repository is a **learning-focused test project**.
- Changes should prioritize clear architecture, maintainability, and educational value over shortcuts.

## Architecture

This repository is a monorepo of independently built services; each Node.js or Go service has its own dependency manifest and runs from its own directory. `docker-compose.yml` wires the services and local infrastructure together.

- `api-gateway` is the client-facing Express entry point. It validates bearer tokens, proxies user requests to `user-service`, and forwards trusted `user-id`, `user-role`, and `correlation-id` headers.
- `user-service` owns registration, authentication, user profiles, and addresses in PostgreSQL.
- `inventory-service` owns products, categories, subcategories, and inventory in PostgreSQL. Its HTTP handlers call services, which use DAO interfaces. Product creation/update publishes `product_update` Kafka messages after persistence; the product ID must be set before publishing.
- `search-service` consumes `product_update`, upserts product documents in Elasticsearch, and serves search requests over gRPC. Its API contract is `search-service/proto/search/search.proto`; generated Go protobuf files are checked in.
- Kafka connects inventory writes to search indexing. PostgreSQL, Elasticsearch, Kafka/Zookeeper, and optional monitoring services are configured in Compose.

## Build, test, and run

Run commands from the indicated service directory; the Go services are separate modules.

| Service | Commands |
| --- | --- |
| `api-gateway`, `user-service` | `npm install`; `npm start` for production mode or `npm run dev` for nodemon development mode |
| `inventory-service`, `search-service` | `go mod download`; `make build`; `make run` (builds and runs `./.bin/app`) |

Run Go tests from either Go service directory with `go test ./...`. The current focused tests are in `inventory-service`; run one with `go test ./app/service -run '^TestNormalizeName$'` or `go test ./pkg/db -run '^TestPostgresDSN$'`.

The repository has a root `Makefile` with `test`, `lint`, and `format` targets. Node services include `start`, `dev`, `test`, and `lint` scripts in their package manifests.

For the full local stack, see the root README; it uses
`podman-compose up --build`. Run the optional data
seeder with `podman-compose --profile seed run --rm seed-data`.

## Repository conventions

- Node services use CommonJS. Keep HTTP routing, controller/request-response handling, and business logic separated in the existing router, controller, and service modules. User persistence is through the existing PostgreSQL query/data-access modules; use the established Express `next(...)` flow for errors and HTTP statuses.
- Gateway authentication is upstream of proxying. Preserve its correlation and trusted identity headers; downstream user-service requests should not receive the bearer token.
- Go services use constructor-injected service/DAO interfaces. Keep PostgreSQL or Elasticsearch operations in DAO implementations, with HTTP/gRPC handlers delegating to services.
- Go configuration is loaded from `conf/{ENVIRONMENT}.yaml` (`default.yaml` when unset); Viper enables environment lookup and maps dots in config keys to underscores. Node configuration selects `dev` by default and `prod` when `NODE_ENV=prod`.
- Normalize catalog category and subcategory names through the existing service helper rather than adding ad hoc casing at individual handlers or DAOs.
- When changing the search RPC contract, edit `search-service/proto/search/search.proto` and regenerate the checked-in stubs using the `protoc` command documented at the top of that file.
- For AI-first evolution, prefer adding documentation, prompts/instructions, tests, and automation guardrails together with code changes.
