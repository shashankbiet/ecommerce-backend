---
name: review-code
description: Perform an independent production code review covering correctness, security, performance, maintainability, and architectural compliance.
---

# Code Review Skill

## Objective

Review implementation changes for production readiness without modifying source code.

## Review areas

1. Business logic correctness
2. Architecture and design consistency
3. Security vulnerabilities
4. Authentication and authorization
5. Database transaction and data integrity concerns
6. Race conditions and concurrency issues
7. Performance and scalability risks
8. Error handling and observability
9. Backward compatibility
10. Test completeness and gaps

## Rules

- Review actual changed code and relevant surrounding context.
- Do not report speculation as confirmed defects.
- Do not approve changes only because tests pass.
- Do not modify source code while reviewing.
- Focus on meaningful defects, not cosmetic style nits.

## Severity model

- **CRITICAL**: security breach or major data corruption risk
- **HIGH**: serious functional or production risk
- **MEDIUM**: scenario-specific incorrect behavior
- **LOW**: maintainability or minor improvement

## Expected output

### 1. Overall Assessment
### 2. Critical Findings
### 3. High Priority Findings
### 4. Medium Priority Findings
### 5. Low Priority Findings
### 6. Missing Tests
### 7. Recommendations

For each finding, include impacted file path, relevant symbol/function, and why it matters.
