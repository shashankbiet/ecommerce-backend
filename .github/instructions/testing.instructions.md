# Testing instructions

## Goals

- Preserve behavior while the project evolves toward AI-first development.
- Keep tests fast, readable, and useful as learning artifacts.

## Service-level commands

- Node services:
  - `cd api-gateway && npm test`
  - `cd user-service && npm test`
- Go services:
  - `cd inventory-service && go test ./...`
  - `cd search-service && go test ./...`

## Authoring guidance

- Add tests only for changed behavior and realistic edge cases.
- Prefer existing test folders and patterns.
- Keep tests deterministic: avoid clock/network randomness where possible.
- For bug fixes, add a test that fails before the fix and passes after.

## AI-first guidance

- Write test names as intent statements so AI tools can infer behavior quickly.
- When introducing AI-assisted features, add explicit validation cases for:
  - malformed input handling
  - contract compatibility
  - fallback/error paths