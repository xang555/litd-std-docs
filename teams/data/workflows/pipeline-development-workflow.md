# Pipeline Development Workflow

## Description
A structured workflow for developing data pipelines from design to production.

## Prerequisites
- Data requirements documented
- Source systems accessible
- Data model designed
- Target schema defined

## Steps

### 1. Requirements Analysis
- **Command**: None
- **Agent**: `data-pipeline-review-agent`
- **Validation**: Requirements clear and feasible
- **Output:** Requirements document

### 2. Data Model Design
- **Command**: None
- **Agent**: None
- **Validation:** Model normalized and performant
- **Output:** ERD and schema files

### 3. Scaffold Pipeline
- **Command**: `pipeline-create`
- **Agent**: None
- **Validation:** Pipeline compiles
- **Output:** Pipeline structure

### 4. Implement Transformation Logic
- **Command**: None
- **Agent**: `data-pipeline-review-agent` (iterative review)
- **Validation:** Logic correct and efficient
- **Output:** Transformation code

### 5. Implement Data Quality Tests
- **Command**: `test-generate`
- **Agent**: None
- **Validation:** Tests cover critical quality checks
- **Output:** Test files

### 6. Local Testing
- **Command**: `test-local`
- **Agent**: None
- **Validation:** Pipeline runs successfully locally
- **Output:** Test results

### 7. Staging Deployment
- **Command**: `deploy-staging`
- **Agent**: None
- **Validation:** Pipeline runs in staging
- **Output:** Staging run results

### 8. Data Validation
- **Command**: `validate-data`
- **Agent**: None
- **Validation:** Output data matches expectations
- **Output:** Validation report

### 9. Performance Testing
- **Command**: `performance-test`
- **Agent**: None
- **Validation:** Pipeline runs within SLA
- **Output:** Performance metrics

### 10. Code Review
- **Command**: None
- **Agent**: `data-pipeline-review-agent`
- **Validation:** Review feedback addressed
- **Output:** Approved pull request

### 11. Production Deployment
- **Command**: `deploy-production`
- **Agent**: None
- **Validation:** Pipeline runs in production
- **Output:** Production run confirmation

### 12. Monitoring Setup
- **Command**: `setup-monitoring`
- **Agent**: None
- **Validation:** Alerts configured
- **Output:** Monitoring dashboard

## Outputs

| Output | Description | Format |
|--------|-------------|--------|
| Requirements document | Data requirements | Markdown |
| ERD diagrams | Data model diagrams | Diagram |
| Pipeline code | Implemented pipeline | Python/SQL |
| Test results | Data quality test results | JSON/HTML |
| Validation report | Output data validation | Markdown |
| Performance metrics | Pipeline performance | Dashboard |

## Estimated Time

- Simple ETL: 1-2 days
- Complex transformation: 3-5 days
- New data mart: 1-2 weeks

## Tags
`data`, `workflow`, `pipeline-development`, `etl`, `agile`
