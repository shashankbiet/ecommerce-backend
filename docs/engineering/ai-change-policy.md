# AI change policy

## Purpose

Define safe, repeatable rules for AI-assisted engineering in this repository.

## Allowed AI assistance

- Drafting implementation plans and technical options
- Generating/refactoring code within approved scope
- Producing and improving tests
- Improving developer documentation and runbooks

## Required human controls

- Human review required for all production-impacting changes
- Do not merge AI-generated changes without validation evidence
- Security-sensitive changes require explicit reviewer attention
- Contract/data changes must include migration and rollback notes

## Prohibited patterns

- Copying generated code without understanding behavior
- Merging with missing tests for changed behavior
- Silent fallback behavior that hides failures
- Storing secrets/credentials in code, prompts, or docs

## Minimum PR requirements for AI-assisted changes

1. AI usage disclosed in PR template
2. Scope and impact documented
3. Tests/lint/build evidence attached
4. Risk and rollback documented
5. Related docs updated

