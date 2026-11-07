# Ecommerce Backend

A comprehensive, microservices-based ecommerce backend platform built with modern tools and best practices. This project demonstrates a scalable architecture with independent services for API Gateway, User Management, Inventory Management, and Product Search.

## Project Overview

**Ecommerce Backend** is a monorepo containing four tightly integrated microservices that work together to provide a complete ecommerce platform. Each service is independently deployable, scalable, and ownable by different teams.


## Services Overview

### 1. **API Gateway** (Node.js/Express)
**Port:** 3001

Single entry point for all client requests. Handles routing, authentication verification, request logging, and response aggregation.

- HTTP routing and request proxying
- Centralized health checks
- Request/response logging (Winston)
- Prometheus metrics collection
- Auth token validation middleware

**Quick Start:** (load `.env` first; see [step 2](#2-configure-local-secrets-and-service-env))
```bash
cd api-gateway
npm install
set -a; . ./.env; set +a
npm start
```

**Documentation:** [API Gateway README](api-gateway/README.md)

---

### 2. **User Service** (Node.js/Express)
**Port:** 3002 | **Database:** MongoDB 

Comprehensive user management including authentication, profiles, and address management with JWT-based authorization.

- User registration and authentication
- JWT token generation and validation
- User profile management
- Address management (multiple addresses per user)
- Password hashing with bcrypt
- Prometheus metrics and Winston logging

**Quick Start:** (load `.env` first; see [step 2](#2-configure-local-secrets-and-service-env))
```bash
cd user-service
npm install
set -a; . ./.env; set +a
npm start
```

**Documentation:** [User Service README](user-service/README.md)

---

### 3. **Inventory Service** (Go 1.26+)
**Port:** 3003 | **Database:** PostgreSQL 

Product and inventory management with real-time stock tracking and category organization.

- Product CRUD operations
- Category and subcategory management
- Inventory tracking and updates
- Kafka producer for product updates
- Transactional database operations
- REST HTTP API endpoints

**Quick Start:** (load `.env` first; see [step 2](#2-configure-local-secrets-and-service-env))
```bash
cd inventory-service
go mod download
set -a; . ./.env; set +a
make run
```

**Documentation:** [Inventory Service README](inventory-service/README.md)

---

### 4. **Search Service** (Go 1.26+)
**Port:** 9004 | **Search Engine:** Elasticsearch

Product search and retrieval via gRPC. Currently implements product lookup by ID; full-text search coming soon.

- Get product by ID (gRPC)
- Kafka consumer for product updates

**Quick Start:** (load `.env` first; see [step 2](#2-configure-local-secrets-and-service-env))
```bash
cd search-service
go mod download
set -a; . ./.env; set +a
make run
```

**Documentation:** [Search Service README](search-service/README.md)

---

## Tech Stack Summary

### **Runtime & Languages**
- **Node.js 20+** - API Gateway, User Service
- **Go 1.26+** - Inventory Service, Search Service

### **API Gateway & Routing**
- **Express.js** - HTTP server and routing framework
- **http-proxy-middleware** - Request proxying and load balancing
- **Winston** - Structured logging

### **Application Frameworks**
- **Express.js** - User Service (REST HTTP API)
- **http/net** - Inventory Service (REST HTTP API)
- **gRPC** - Search Service (RPC framework)
- **Protocol Buffers** - Service contract definition

### **Databases**
- **MongoDB** - User profiles, addresses, authentication data
- **PostgreSQL 16** - Products, categories, inventory, stock tracking
- **Elasticsearch** - Full-text search index (Coming Soon)

### **Message Queue & Events**
- **Apache Kafka** - Event streaming platform
- **Zookeeper** - Kafka coordination and broker management
- **Topics:**
  - `product_update` - Product creation/update events
  - `inventory_update` - Stock level changes

### **Logging & Monitoring**
- **Winston** - Structured logging (Node.js services)
- **Zap** - High-performance logging (Go services)
- **Loki** - Log aggregation and querying
- **Prometheus** - Metrics collection and time-series database
- **Grafana** - Metrics dashboards and alerting

### **Security & Authentication**
- **bcrypt** - Password hashing and verification
- **jwt-simple** - JWT token generation and validation
- **Custom Middleware** - Authentication and request validation

### **Containerization & Orchestration**
- **Podman** - Container images with multi-stage builds
- **Podman Compose** - Local environment orchestration
- **Alpine/Slim Images** - Lightweight container images
- **.dockerignore** - Build optimization

### **Package Management**
- **npm** - Node.js dependencies
- **go mod** - Go module management

### **Build Tools**
- **Make** - Go service automation
- **npm scripts** - Node.js build and run commands

## Quick Start Guide

### Prerequisites
- Podman & Podman Compose
- Node.js 20+ (for local development)
- Go 1.26+ (for local development)
- golangci-lint v2 (for local Go linting; install using the [official instructions](https://golangci-lint.run/welcome/install/))
- OpenSSL (used to generate local secret files)
- Git

### Run Linters Locally
From the repository root, run the Go linter for both Go services:
```bash
make lint-go
```

To run all Go and Node.js linters, first install the Node.js dependencies:
```bash
npm --prefix api-gateway ci
npm --prefix user-service ci
make lint
```

### 1. Clone Repository
```bash
git clone git@github.com:shashankbiet/ecommerce-backend.git
cd ecommerce-backend
```

### 2. Configure Local Secrets and Service Env
Each service reads its own `.env` file (Git-ignored). Copy the templates and
create the secret files; secret values are stored one per file under
`secrets/` (also Git-ignored) and mounted into containers by Compose.
```bash
mkdir -p secrets
for name in api_gateway_access_token_secret user_service_admin_verification_token \
  mongodb_root_password sql_config_password elastic_search_password grafana_admin_password; do
  openssl rand -hex 32 > "secrets/$name"
done
for svc in api-gateway user-service inventory-service search-service; do
  cp -n "$svc/.env.example" "$svc/.env"
done
```
Optionally, run the drift check to confirm the `.env` files, `secrets/`, and
`docker-compose.yml` agree:
```bash
python3 scripts/check_env_drift.py
```
The generated values are for local testing only. The MongoDB root username
(`mongo-ecommerce`) and PostgreSQL username (`postgres_user`) and database
(`inventorydb`) are set in `docker-compose.yml`; keep them equal to the values in
the service `.env` files. Image initialization values apply only when the
database is first created. If you already have data in `bind-mount/db/mongodb`
or `bind-mount/db/postgres`, back it up and make sure its existing credentials
match the secret files; the initialization variables do not rotate credentials
in an already-initialized database. MongoDB data created before authentication
was enabled needs a deliberate authentication migration.

### 3. Start All Services with Podman Compose
```bash
podman-compose up --build
```

Compose builds application images from the repository root and uses the root
`.dockerignore`; each Dockerfile copies only its service's files.

For a memory-constrained machine, stop the running stack and
serialize image builds to avoid compiler processes competing for memory:
```bash
podman-compose stop
podman-compose --parallel 1 up --build
```

To run a service directly on the host instead of in a container, load its
`.env` into the shell from that service's directory (its `*_FILE` paths are
relative to the service directory):
```bash
cd inventory-service
set -a; . ./.env; set +a
make run
```

This starts the application services and their dependencies:
- **API Gateway**: host loopback port 3001; intended client entry point.
- **User Service**: internal to the Compose network at `user-service:3002`.
- **Inventory Service**: bound to the host loopback address on port 3003.
- **Search Service**: bound to host loopback port 9004.
- **MongoDB**: authenticated and bound to host loopback port 27017.
- **PostgreSQL**: bound to host loopback port 5433.
- **Elasticsearch**: host loopback port 9200 is published.
- **Kafka**: host loopback ports 9092 and 9999 are published; port 29092
  remains available through the host-gateway listener for Docker clients.
- **Zookeeper**: available on the Compose network; no host port is published.
- **Loki**: bound to the host loopback address on port 3100.
- **Prometheus and Grafana**: bound to host loopback ports 9090 and 3000 when
  the monitoring profile is enabled.

### Seed Local Test Data

After starting the stack (inventory-service must have run once to create its
tables), load the seed files from `scripts/seed/`:
```bash
podman exec -i mongodb mongo --username mongo-ecommerce \
  --password "$(cat secrets/mongodb_root_password)" --authenticationDatabase admin \
  userService < scripts/seed/mongo.js
podman exec -e PGPASSWORD="$(cat secrets/sql_config_password)" -i postgres \
  psql -U postgres_user -d inventorydb < scripts/seed/postgres.sql
curl -H 'Content-Type: application/x-ndjson' -XPOST \
  'http://localhost:9200/_bulk?refresh=true' \
  --data-binary @scripts/seed/elasticsearch.ndjson
```
This creates 25 demo users, 15 categories, 60 subcategories, and 105 products
with inventory records, and indexes the products into Elasticsearch. The
MongoDB and PostgreSQL scripts are safe to rerun; the Elasticsearch bulk load
overwrites the same document ids. Demo accounts use the password
`DemoPass123!`; these credentials and records are for local testing only.

To enable Prometheus, Grafana, and Podman container log collection, write a
local Grafana admin password to `secrets/grafana_admin_password` (created in
step 2) and start the monitoring profile:
```bash
podman-compose --profile monitoring up -d
```
Add `--build` only after changing service code or dependencies.

Grafana and Prometheus are available on `http://localhost:3000` and
`http://localhost:9090`. Grafana is bound to localhost and anonymous admin
access is disabled; its password comes from `secrets/grafana_admin_password`. Loki remains enabled
because the Node.js services send logs to it directly. Promtail reads container
logs from the systemd journal at `/srv/journal`. Create that path once before
starting the profile:

- **Linux:** `sudo ln -s /var/log/journal /srv/journal`
- **macOS (Podman machine):** `podman-compose` resolves bind-mount sources on the
  Mac, where `/var/log/journal` does not exist, so expose the VM journal under
  `/var/srv/journal` with a persistent mount unit:
  ```bash
  podman machine ssh podman-machine-default 'sudo tee /etc/systemd/system/var-srv-journal.mount >/dev/null <<EOF
  [Unit]
  Description=Expose the systemd journal for Promtail

  [Mount]
  What=/var/log/journal
  Where=/var/srv/journal
  Options=bind

  [Install]
  WantedBy=local-fs.target
  EOF
  sudo systemctl daemon-reload && sudo systemctl enable --now var-srv-journal.mount'
  ```

Database files are stored under `bind-mount/db/`: MongoDB and PostgreSQL are
already persisted there, and Elasticsearch uses `bind-mount/db/elastic-search`.
When monitoring is enabled, Prometheus data, Grafana state, Loki logs, and
Promtail journal positions are stored under `bind-mount/monitoring/`. These are
local development data directories and are excluded from Git. Back them up
separately; deleting their contents deletes the associated data. Prometheus
retains 15 days of metrics. The `:U` bind option lets Podman set directory
ownership for the container processes; this stack configuration uses Podman
Compose for local data mounts.

### Monitoring and Logs

Metrics and logs are collected by different paths:

- **Metrics (pull):** Prometheus scrapes every 15 seconds, as configured in
  [prometheus.yml](./prometheus.yml). Each service exposes its own endpoint:

  | Job | Target | Path | Main metrics |
  | --- | --- | --- | --- |
  | `api-gateway` | `api-gateway:3001` | `/metrics` | `gateway_http_request_duration_sec` |
  | `user-service` | `user-service:3002` | `/api/user/metrics` | `rest_api_response_time_sec`, `db_response_time_sec`, `nodejs_*` |
  | `inventory-service` | `inventory-service:3003` | `/metrics` | `inventory_service_http_requests_total`, `inventory_service_http_request_duration_seconds`, `go_*`, `process_*` |
  | `search-service` | `search-service:9104` | `/metrics` | `grpc_server_*`, `ecommerce_backend_search_service_*` |

- **Logs (push and collect):**
  - Node.js services (`api-gateway`, `user-service`) send Winston logs directly
    to Loki with the label `app`.
  - All containers write to stdout. Podman sends this to the systemd journal,
    and Promtail reads it from `/srv/journal` and labels each line with
    `container`.

- **Grafana** reads both sources. Its Prometheus and Loki data sources are
  provisioned from [datasources.yml](./datasources.yml).

Quick checks after starting the profile:
```bash
# Gateway metrics (exposed on the host loopback port)
curl http://localhost:3001/metrics | grep gateway_http_request_duration_sec

# Scrape status for every job
curl -s http://localhost:9090/api/v1/targets | grep -o '"job":"[^"]*"' | sort -u

# Logs from one container, via Loki
curl -G -s http://localhost:3100/loki/api/v1/query_range \
  --data-urlencode 'query={container="api-gateway"}' --data-urlencode 'limit=5'
```
Prometheus target status is also visible at `http://localhost:9090/targets`.

Known limitations:
- Node.js log lines can appear twice in Loki, once from the direct Winston push
  (`app` label) and once from the journal (`container` label).
- The Winston log stream and the journal stream do not share labels, so a single
  label selector does not match both.
- Prometheus scrapes the four application services plus itself, Loki, and
  Promtail. MongoDB, PostgreSQL, Elasticsearch, and Kafka are not scraped.

### Observability Check and Load Testing

Scripts live in [scripts/load-test/](./scripts/load-test/). Install their
dependencies once, then run them against the running stack:
```bash
cd scripts/load-test
npm install
```

**Observability check** verifies that Prometheus, Loki, and Grafana are healthy
and receiving data from the services. Start the monitoring profile first:
```bash
USER_PASSWORD='DemoPass123!' npm run observability-check
```
It checks Prometheus readiness, scrape target status, series for each service
job, and the HTTP request rate; Loki readiness and log labels; and Grafana
health. It exits non-zero if any check fails. The URLs default to
`PROMETHEUS_URL=http://localhost:9090`, `LOKI_URL=http://localhost:3100`, and
`GRAFANA_URL=http://localhost:3000`.

**Load test** sends traffic to 12 endpoints (gateway, inventory, and search
gRPC) and prints requests, RPS, error rate, p50/p95/p99/max latency, and status
counts per scenario:
```bash
USER_PASSWORD='DemoPass123!' npm run load-test
```
`USER_PASSWORD` is required. The script logs in through the gateway to get a
token for the authenticated routes. Each scenario runs concurrently with
`WORKERS_PER_SCENARIO` workers (default `3`) for `DURATION_SECONDS` (default
`300`). Other optional variables:

| Variable | Default |
| --- | --- |
| `GATEWAY_URL` | `http://localhost:3001` |
| `INVENTORY_URL` | `http://localhost:3003` |
| `SEARCH_ADDRESS` | `localhost:9004` |
| `USER_EMAIL` | `demo.customer01@example.com` |

Use a short run to confirm the setup before a full run:
```bash
USER_PASSWORD='DemoPass123!' DURATION_SECONDS=10 npm run load-test
```
Run load tests only against local or test environments. Watch the results live in
the provisioned **Ecommerce Backend Overview** dashboard in Grafana (folder
`Ecommerce`), or query Prometheus directly at `http://localhost:9090`.

### 4. Verify Services
```bash
# Check API Gateway health
curl http://localhost:3001/health

# Check User Service health (internal-only service; run from container namespace)
podman exec user-service wget -qO- http://localhost:3002/health

# Check Inventory Service health
curl http://localhost:3003/health
```

Useful container commands:
```bash
# Show service status
podman-compose ps

# Follow logs for one service
podman-compose logs -f inventory-service

# Stop containers and keep data
podman-compose stop

# Remove containers and networks, keep bind-mount data
podman-compose --profile monitoring down
```
Data in `bind-mount/` is not removed by `down`. Delete those directories
manually only when you want to reset the local databases and monitoring data.

### 5. Stop and Clean Up
Use `podman-compose stop` to pause the stack, or the `down` command above to
remove its containers and networks. Re-run `podman-compose up` to start again.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
