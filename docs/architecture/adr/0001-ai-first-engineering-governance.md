# ADR 0001: AI-first engineering governance

## Status

Accepted

## Context

This repository began as a learning project and is now evolving into an AI-first engineering project.  
We need explicit governance so AI assistance increases velocity without reducing quality or safety.

## Decision

Adopt an AI-first workflow with:

- Documented repository instructions and role-specific agent guidance
- Standard PR/issue templates capturing AI usage, risk, and validation
- Explicit testing and documentation expectations for changed behavior
- Metrics to monitor quality outcomes of AI-assisted delivery

## Consequences

### Positive

- More consistent delivery process across contributors
- Better auditability of engineering decisions
- Faster onboarding and clearer expectations

### Negative

- Slightly higher process overhead per change
- Requires discipline to keep docs and templates current

## Alternatives considered

1. **Ad hoc AI usage with no governance**
   - Rejected due to inconsistency and higher risk.
2. **No AI usage**
   - Rejected because it conflicts with project direction and learning goals.

