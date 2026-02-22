# Ecommerce Backend

A comprehensive, microservices-based ecommerce backend platform built with modern tools and best practices. This project demonstrates a scalable architecture with independent services for API Gateway, User Management, Inventory Management, and Product Search.

## Project Overview

**Ecommerce Backend** is a monorepo containing four tightly integrated microservices that work together to provide a complete ecommerce platform. Each service is independently deployable, scalable, and ownable by different teams.


## Services Overview

### 1. **API Gateway** (Node.js/Express)
**Port:** 3001

Single entry point for all client requests. Handles routing, authentication verification, request logging, and response aggregation.

- ✅ HTTP routing and request proxying
- ✅ Centralized health checks
- ✅ Request/response logging (Winston)
- ✅ Prometheus metrics collection
- ✅ Auth token validation middleware

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

- ✅ User registration and authentication
- ✅ JWT token generation and validation
- ✅ User profile management
- ✅ Address management (multiple addresses per user)
- ✅ Password hashing with bcrypt
- ✅ Prometheus metrics and Winston logging

**Quick Start:**
```bash
cd user-service
npm install
npm start
```

**Documentation:** [User Service README](user-service/README.md)

---

### 3. **Inventory Service** (Go 1.25+)
**Port:** 3003 | **Database:** MySQL 

Product and inventory management with real-time stock tracking and category organization.

- ✅ Product CRUD operations
- ✅ Category and subcategory management
- ✅ Inventory tracking and updates
- ✅ Kafka producer for product updates
- ✅ Transactional database operations
- ✅ REST HTTP API endpoints

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

- ✅ Get product by ID (gRPC)
- ✅ Kafka consumer for product updates

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
- **MySQL 8.0** - Products, categories, inventory, stock tracking
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
- **Docker** - Container images with multi-stage builds
- **Docker Compose** - Local environment orchestration
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
- Docker & Docker Compose
- Node.js 20+ (for local development)
- Go 1.25+ (for local development)
- Git

### 1. Clone Repository
```bash
git clone git@github.com:shashankbiet/ecommerce-backend.git
cd ecommerce-backend
```

### 2. Start All Services with Docker Compose
```bash
docker-compose up --build
```

This starts all services and dependencies:
- API Gateway (port 3001)
- User Service (port 3002, MongoDB)
- Inventory Service (port 3003, MySQL)
- Search Service (port 9004, Elasticsearch)
- Supporting infrastructure (Kafka, Zookeeper, Prometheus, Grafana, Loki)

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
