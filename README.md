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

**Quick Start:**
```bash
cd api-gateway
npm install
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

**Quick Start:**
```bash
cd user-service
npm install
npm start
```

**Documentation:** [User Service README](user-service/README.md)

---

### 3. **Inventory Service** (Go 1.25+)
**Port:** 3003 | **Database:** PostgreSQL 

Product and inventory management with real-time stock tracking and category organization.

- Product CRUD operations
- Category and subcategory management
- Inventory tracking and updates
- Kafka producer for product updates
- Transactional database operations
- REST HTTP API endpoints

**Quick Start:**
```bash
cd inventory-service
go mod download
make run
```

**Documentation:** [Inventory Service README](inventory-service/README.md)

---

### 4. **Search Service** (Go 1.25+)
**Port:** 9004 | **Search Engine:** Elasticsearch

Product search and retrieval via gRPC. Currently implements product lookup by ID; full-text search coming soon.

- Get product by ID (gRPC)
- Kafka consumer for product updates

**Quick Start:**
```bash
cd search-service
go mod download
make run
```

**Documentation:** [Search Service README](search-service/README.md)

---

## Tech Stack Summary

### **Runtime & Languages**
- **Node.js 20+** - API Gateway, User Service
- **Go 1.25+** - Inventory Service, Search Service

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
- Go 1.25+ (for local development)
- Git

### 1. Clone Repository
```bash
git clone git@github.com:shashankbiet/ecommerce-backend.git
cd ecommerce-backend
```

### 2. Start All Services with Podman Compose
```bash
podman-compose up --build
```

For a memory-constrained machine, stop the running stack and
serialize image builds to avoid compiler processes competing for memory:
```bash
podman-compose stop
podman-compose --parallel 1 up --build
```

This starts the application services and their dependencies:
- **API Gateway**: host port 3001; intended client entry point.
- **User Service**: internal to the Compose network at `user-service:3002`.
- **Inventory Service**: bound to the host loopback address on port 3003.
- **Search Service**: host port 9004 is published.
- **MongoDB**: host port 27017 is published.
- **PostgreSQL**: host port 5433 is published.
- **Elasticsearch**: host ports 9200 and 9300 are published.
- **Kafka**: host ports 9092, 29092, and 9999 are published.
- **Zookeeper**: available on the Compose network; no host port is published.
- **Loki**: bound to the host loopback address on port 3100.
- **Prometheus and Grafana**: bound to host loopback ports 9090 and 3000 when
  the monitoring profile is enabled.

### Seed Local Test Data

After starting the stack (inventory-service must have run once to create its
tables), load the seed files from `scripts/seed/`:
```bash
podman exec -i user-mongodb mongo userService < scripts/seed/mongo.js
podman exec -i inventory-pg psql -U inventoryuser -d inventorydb < scripts/seed/postgres.sql
curl -H 'Content-Type: application/x-ndjson' -XPOST \
  'http://localhost:9200/_bulk?refresh=true' \
  --data-binary @scripts/seed/elasticsearch.ndjson
```
This creates 25 demo users, 15 categories, 60 subcategories, and 105 products
with inventory records, and indexes the products into Elasticsearch. The
MongoDB and PostgreSQL scripts are safe to rerun; the Elasticsearch bulk load
overwrites the same document ids. Demo accounts use the password
`DemoPass123!`; these credentials and records are for local testing only.

To enable Prometheus, Grafana, and Podman container log collection, set a local
Grafana admin password and start the monitoring profile:
```bash
GRAFANA_ADMIN_PASSWORD='choose-a-local-password' podman-compose --parallel 1 \
  --profile monitoring up --build
```
Grafana and Prometheus are available on `http://localhost:3000` and
`http://localhost:9090`. Grafana is bound to localhost and anonymous admin
access is disabled (the local default login is `admin`/`admin`; set the
environment variable above to change it). Loki remains enabled because the
Node.js services send logs to it directly. Promtail reads Podman machine
journald records for container logs when the monitoring profile is enabled.

Database files are stored under `bind-mount/db/`: MongoDB and PostgreSQL are
already persisted there, and Elasticsearch uses `bind-mount/db/elasticsearch`.
When monitoring is enabled, Prometheus data, Grafana state, Loki logs, and
Promtail journal positions are stored under `bind-mount/monitoring/`. These are
local development data directories and are excluded from Git. Back them up
separately; deleting their contents deletes the associated data. Prometheus
retains 15 days of metrics. The `:U` bind option lets Podman set directory
ownership for the container processes; this stack configuration uses Podman
Compose for local data mounts.

### 3. Verify Services
```bash
# Check API Gateway health
curl http://localhost:3001/health

# Check User Service health
curl http://localhost:3002/health

# Check Inventory Service health
curl http://localhost:3003/health
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
