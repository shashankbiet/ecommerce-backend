# AI quality metrics

## Purpose

Measure whether AI-assisted delivery improves speed and quality without increasing production risk.

## Core KPIs

1. **Review rework rate**
   - Percentage of AI-assisted PRs needing significant rework after review.
2. **Defect escape rate**
   - Production defects linked to AI-assisted changes per release window.
3. **Validation completeness**
   - Percentage of PRs with required test/lint/build evidence.
4. **Cycle time**
   - Time from issue start to merged PR.
5. **Rollback incidence**
   - Number of rollbacks/reverts for AI-assisted changes.

## Suggested targets (initial)

- Validation completeness: >= 95%
- Defect escape rate: non-increasing versus baseline
- Rollback incidence: <= baseline

## Review cadence

- Weekly: quick KPI trend review
- Monthly: adjust agent prompts, instructions, and guardrails

