---
name: generate-tests
description: Generate focused unit, integration, and regression tests for implemented changes while preserving existing behavior and architecture.
---

# Test Generation Skill

## Objective

Create meaningful tests for changed behavior using the repository's existing test frameworks and conventions.

## Instructions

1. Understand the requirement and implementation.
2. Identify behavior changes and risk areas.
3. Add tests that fail before and pass after the change.
4. Cover success, failure, and edge-case scenarios.
5. Reuse existing helpers, fixtures, and patterns.
6. Keep tests deterministic and maintainable.
7. Run the smallest relevant test scope first.

## Rules

- Do not modify production logic only to make tests pass.
- Do not add redundant tests that only restate implementation.
- Avoid brittle time/network-dependent assertions.
- Keep test data minimal and readable.

## Expected output

### 1. Test Strategy
### 2. Scenarios Covered
### 3. Tests Added or Updated
### 4. Execution Results
### 5. Remaining Risks
