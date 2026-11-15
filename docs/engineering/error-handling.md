# Error handling

## Principles

- Fail loudly and predictably.
- Return actionable error responses.
- Preserve root-cause context in logs.

## Node.js services

- Use centralized error flow via `next(...)`.
- Prefer existing error/status utilities where available.
- Keep transport-level formatting in controllers/middleware, not in business logic.

## Go services

- Return explicit errors from DAO/service layers.
- Map domain/data errors to appropriate HTTP/gRPC status codes in handlers.
- Avoid swallowing errors in goroutines or background consumers.

## Observability

- Include correlation IDs in request paths that cross services.
- Log enough context to trace user-visible failures across gateway, services, and async consumers.