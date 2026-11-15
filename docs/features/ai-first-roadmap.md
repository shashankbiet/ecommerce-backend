# AI-first roadmap

## Why this exists

This backend started as a personal learning project and is now being modernized into an AI-first engineering playground.

## Phase 1: Baseline (now)

- Document architecture and engineering standards.
- Keep CI checks green for Node + Go services.
- Make service boundaries and contracts explicit.
- Establish AI governance assets:
  - PR/Issue templates in `.github/`
  - AI change policy and metrics in `docs/engineering/`
  - ADR records in `docs/architecture/adr/`

## Phase 2: AI-assisted delivery

- Strengthen `.github` Copilot instructions for implementation, database, and testing.
- Require docs updates with behavior/contract changes.
- Expand targeted automated tests for regressions found by AI-assisted development.

## Phase 3: AI-enabled product capabilities

- Evaluate AI-driven search enrichment and ranking improvements.
- Add safe experimentation paths (feature flags/config toggles).
- Define observability and quality gates for AI-influenced features.

## Success criteria

- Faster iteration without dropping reliability.
- Better onboarding value for engineers learning microservices + AI-assisted workflows.
- Clear, auditable technical decisions across docs, code, and CI.
