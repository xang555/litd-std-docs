# Deployment Workflow

## Description
A structured workflow for deploying applications to production with proper validation and monitoring.

## Prerequisites
- Application tests passing
- Docker image built and pushed
- Infrastructure provisioned
- Monitoring configured

## Steps

### 1. Pre-deployment Checks
- **Command**: None
- **Agent**: `infrastructure-review-agent`
- **Validation**: Infrastructure code reviewed
- **Output**: Pre-deployment checklist

### 2. Security Scan
- **Command**: `security-scan`
- **Agent**: None
- **Validation**: No critical vulnerabilities
- **Output**: Security scan report

### 3. Deploy to Staging
- **Command**: `deploy staging <service>`
- **Agent**: None
- **Validation**: Deployment successful
- **Output**: Staging deployment

### 4. Integration Tests
- **Command**: `run-integration-tests`
- **Agent**: None
- **Validation**: All tests pass
- **Output**: Test results

### 5. Performance Tests
- **Command**: `run-performance-tests`
- **Agent**: None
- **Validation**: Performance within thresholds
- **Output**: Performance report

### 6. Create Deployment Plan
- **Command**: None
- **Agent**: None
- **Validation**: Plan approved
- **Output**: Deployment plan document

### 7. Notify Stakeholders
- **Command**: `notify-deployment`
- **Agent**: None
- **Validation**: Notifications sent
- **Output**: Confirmation of notifications

### 8. Deploy to Production
- **Command**: `deploy production <service>`
- **Agent**: None
- **Validation**: Deployment healthy
- **Output**: Production deployment

### 9. Smoke Tests
- **Command**: `run-smoke-tests`
- **Agent**: None
- **Validation**: Critical paths working
- **Output**: Smoke test results

### 10. Monitor Deployment
- **Command**: `monitor-deployment`
- **Agent**: None
- **Validation**: Metrics within normal range
- **Output**: Monitoring dashboard

### 11. Finalize Deployment
- **Command**: `finalize-deployment`
- **Agent**: None
- **Validation**: Deployment stable
- **Output**: Deployment complete

## Outputs

| Output | Description | Format |
|--------|-------------|--------|
| Pre-deployment checklist | Deployment readiness checklist | Markdown |
| Security scan report | Vulnerability scan results | JSON/PDF |
| Test results | Integration and performance tests | JSON/HTML |
| Deployment plan | Step-by-step deployment plan | Markdown |
| Deployment summary | Final deployment status | Email/Slack |

## Estimated Time

- Small service: 1 hour
- Medium service: 2-3 hours
- Large service: 4-6 hours

## Rollback Triggers

- Critical security issues detected
- Health checks failing
- Error rate above threshold
- Performance degradation

## Tags
`devops`, `workflow`, `deployment`, `production`, `cicd`
