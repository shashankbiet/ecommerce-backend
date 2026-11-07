.DEFAULT_GOAL := help

GO_SERVICES := inventory-service search-service
NODE_SERVICES := api-gateway user-service

.PHONY: help test test-go test-node lint lint-go lint-node format

help:
	@echo "Targets:"
	@echo "  test     Run Go and Node unit tests"
	@echo "  lint     Run Go and Node linters locally"
	@echo "  lint-go  Run golangci-lint on the Go services"
	@echo "  format   Apply gofmt and Prettier to the services"

test: test-go test-node

test-go:
	@for s in $(GO_SERVICES); do echo "==> go test $$s"; (cd $$s && go test ./...) || exit 1; done

test-node:
	@for s in $(NODE_SERVICES); do echo "==> npm test $$s"; (cd $$s && npm test) || exit 1; done

lint: lint-go lint-node

lint-go:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint is required; see https://golangci-lint.run/welcome/install/" >&2; exit 1; }
	@for s in $(GO_SERVICES); do echo "==> golangci-lint $$s"; (cd $$s && golangci-lint run ./...) || exit 1; done

lint-node:
	@for s in $(NODE_SERVICES); do echo "==> npm run lint $$s"; (cd $$s && npm run lint) || exit 1; done

format:
	@for s in $(GO_SERVICES); do (cd $$s && gofmt -w .); done
	@for s in $(NODE_SERVICES); do (cd $$s && npx prettier --write "**/*.js"); done
