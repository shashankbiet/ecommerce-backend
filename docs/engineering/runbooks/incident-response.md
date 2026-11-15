# Incident response runbook

## 1. Triage

- Identify impacted service(s)
- Confirm severity and user impact
- Assign incident owner

## 2. Stabilize

- Stop harmful writes/traffic if needed
- Use rollback path for recent risky change
- Capture evidence before restarting components

## 3. Diagnose

- Check service health endpoints
- Inspect recent deploys/PRs
- Correlate logs with `correlation-id` where applicable

## 4. Recover

- Apply fix or rollback
- Validate with targeted smoke checks and tests
- Monitor error rates and latency

## 5. Post-incident

- Record timeline and root cause
- Add/adjust tests to prevent recurrence
- Update docs/guardrails if process gaps were found

