# System overview

## Project intent

This repository is a test/learning project maintained by an engineering manager and originally created around 3 years ago.  
The current direction is to evolve it into an AI-first backend platform while keeping the architecture understandable for hands-on learning.

## Services

- **api-gateway** (Node.js/Express): public entrypoint, auth token validation, request proxying.
- **user-service** (Node.js/Express + PostgreSQL): authentication, user profile, and address management.
- **inventory-service** (Go + PostgreSQL): categories, subcategories, products, and inventory.
- **search-service** (Go + gRPC + Elasticsearch): search index and query API.

## Integration model

- Synchronous flow: API Gateway routes user-facing HTTP traffic to backend services.
- Asynchronous flow: Inventory publishes `product_update` to Kafka; Search consumes and indexes into Elasticsearch.

## AI-first direction

- Keep architecture docs and implementation synchronized.
- Favor clear contracts and testable boundaries.
- Add AI-ready development guardrails in `.github/` instructions, workflows, and skill/agent definitions.