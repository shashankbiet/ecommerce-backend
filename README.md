# Ecommerce Backend

Microservices ecommerce platform in a monorepo. Each service is independently deployable.

| Service | Stack | Port | Docs |
| --- | --- | --- | --- |
| API Gateway | Node.js/Express | 3001 | [README](api-gateway/README.md) |
| User Service | Node.js/Express, PostgreSQL | 3002 | [README](user-service/README.md) |
| Inventory Service | Go, PostgreSQL, Kafka | 3003 | [README](inventory-service/README.md) |
| Search Service | Go, gRPC, Elasticsearch | 9004 | [README](search-service/README.md) |

## Data Stores and Topics

- **PostgreSQL 16** (`user_db`): users, addresses, authentication data
- **PostgreSQL 16** (`inventory_db`): categories, subcategories, products, inventory
- **Elasticsearch**: product search index
- **Kafka topics:** `product_update` (inventory → search), `inventory_update`

## Prerequisites
- Podman and Podman Compose
- Node.js 20+ and Go 1.26+ (local development)
- golangci-lint v2 (Go linting)
- OpenSSL

## Setup

```bash
git clone git@github.com:shashankbiet/ecommerce-backend.git
cd ecommerce-backend
```

### 1. Create secrets and env files
Secrets and `.env` files are Git-ignored and for local testing only.
```bash
umask 077
mkdir -p secrets
for name in api_gateway_access_token_secret user_service_admin_verification_token \
  sql_config_password elastic_search_password grafana_admin_password \
  redis_username redis_password clickhouse_password minio_root_password; do
  if [ ! -s "secrets/$name" ]; then openssl rand -hex 32 > "secrets/$name"; fi
done
[ -s secrets/clickhouse_user ] || printf 'clickhouse\n' > secrets/clickhouse_user
[ -s secrets/minio_root_user ] || printf 'minio\n' > secrets/minio_root_user
chmod 644 secrets/*
for svc in api-gateway user-service inventory-service search-service; do
  cp -n "$svc/.env.example" "$svc/.env"
done
if [ ! -f .env ]; then
  {
    printf 'REDIS_AUTH=%s\n' "$(cat secrets/redis_password)"
    printf 'LANGFUSE_POSTGRES_USER=postgres_user\n'
    printf 'LANGFUSE_POSTGRES_PASSWORD=%s\n' "$(cat secrets/sql_config_password)"
    printf 'LANGFUSE_POSTGRES_DB=langfuse\n'
    printf 'SALT=%s\n' "$(openssl rand -hex 16)"
    printf 'ENCRYPTION_KEY=%s\n' "$(openssl rand -hex 32)"
    printf 'CLICKHOUSE_USER=%s\n' "$(cat secrets/clickhouse_user)"
    printf 'CLICKHOUSE_PASSWORD=%s\n' "$(cat secrets/clickhouse_password)"
    printf 'NEXTAUTH_SECRET=%s\n' "$(openssl rand -hex 32)"
    printf 'ELASTIC_PASSWORD=%s\n' "$(cat secrets/elastic_search_password)"
    printf 'MINIO_ROOT_USER=%s\n' "$(cat secrets/minio_root_user)"
    printf 'MINIO_ROOT_PASSWORD=%s\n' "$(cat secrets/minio_root_password)"
  } > .env
  chmod 600 .env
fi
```
Elasticsearch's `elastic` password (`ELASTIC_PASSWORD` in the root `.env`) and the
MinIO `MINIO_ROOT_*` values must match the files in `secrets/`.

### 2. Start the stack
```bash
podman-compose up --build
```
Give the Podman machine at least **8 GiB** of memory; with less, containers such as
Elasticsearch are killed by the OOM killer.

Published endpoints: API Gateway `localhost:3001`, Inventory `localhost:3003`,
Search (gRPC) `localhost:9004`.

To run one service on the host instead of in a container:
```bash
cd inventory-service
set -a; . ./.env; set +a
make run
```

### Seed demo data (optional)
Run after the stack is up (user-service and inventory-service must have started once):
```bash
podman exec -e PGPASSWORD="$(cat secrets/sql_config_password)" -i postgres \
  psql -U postgres_user -d postgres < scripts/seed/postgres.sql
curl -H 'Content-Type: application/x-ndjson' -XPOST \
  'http://localhost:9200/_bulk?refresh=true' \
  --data-binary @scripts/seed/elasticsearch.ndjson
```
Creates 25 demo users (password `DemoPass123!`), one address for each of the 20 demo customers, 105 products, and their inventory. Local testing only.

### Monitoring (optional)
Requires `secrets/grafana_admin_password` from step 1.
```bash
podman-compose --profile monitoring up -d
```
Grafana: http://localhost:3000 · Prometheus: http://localhost:9090 · Loki: http://localhost:3100

Promtail reads container logs from `/srv/journal`. Create it once:
- **Linux:** `sudo ln -s /var/log/journal /srv/journal`
- **macOS (Podman machine):** expose the VM journal at `/var/srv/journal` with a
  bind-mount unit:
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

Data is stored in `bind-mount/` (Git-ignored); `down` does not remove it.
To reset all local state and reseed from scratch:
```bash
rm -rf bind-mount
podman-compose up --build
```

## Metrics and logs

- **Metrics:** Prometheus scrapes `api-gateway`, `user-service`, `inventory-service`,
  and `search-service` (see [prometheus.yml](./prometheus.yml)).
- **Logs:** Node.js services push Winston logs straight to Loki. Container stdout
  is read by Promtail from the journal.
- **Known limitation:** Node.js log lines can appear twice in Loki, once per source
  (`app` and `container` labels).

## Load and observability checks

```bash
cd scripts/load-test
npm install
USER_PASSWORD='DemoPass123!' npm run observability-check
USER_PASSWORD='DemoPass123!' DURATION_SECONDS=10 npm run load-test
```
Run load tests only against local or test environments. `DURATION_SECONDS`
defaults to `300`.

## Verify, lint, and stop

```bash
curl http://localhost:3001/health      # API Gateway
curl http://localhost:3003/health      # Inventory Service
podman exec user-service wget -qO- http://localhost:3002/api/user/health   # User Service (internal)
```

```bash
make lint-go                           # Go linters (run from repo root)
make lint                              # Go + Node.js linters (needs `npm --prefix <svc> ci`)
```

```bash
podman-compose stop                    # pause containers
podman-compose --profile monitoring down --remove-orphans   # remove containers and networks, including leftovers from removed services
podman ps -a --pod                     # confirm nothing from this project is still running
```

If `down` fails because containers are still running or a pod was left behind from an
older compose file, remove the pod directly (this removes its containers):

```bash
podman pod ls                          # find the pod name or ID
podman pod rm -f <pod-id>              # force-remove the pod and its containers
podman network rm ecommerce-backend_default   # only if the network is still listed
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
