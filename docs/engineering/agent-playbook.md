# Agent playbook

## Purpose

Map common engineering tasks to the right agent/skill with expected outputs.

## Task routing

### 1) New feature design
- Use: Architect agent or `analyze-feature` skill
- Output: impacted modules, contracts/data changes, risks, open questions

### 2) Feature implementation
- Use: Developer agent or `implement-feature` skill
- Output: scoped implementation, tests, validation results

### 3) Test expansion/regression safety
- Use: Test Engineer agent or `generate-tests` skill
- Output: added scenarios, deterministic tests, execution results

### 4) Code review
- Use: Code Reviewer agent or `review-code` skill
- Output: severity-ranked findings, missing tests, recommendations

### 5) Safe cleanup/refactor
- Use: `refactor-code` skill
- Output: behavior-preserving improvements + regression check evidence

## Definition of done (AI-assisted tasks)

- Requested scope fully implemented
- Relevant tests executed and results captured
- Contracts/data changes documented
- Risk + rollback path documented

