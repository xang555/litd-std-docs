# API Development Workflow

## Description
A structured workflow for developing new REST APIs from design to deployment.

## Prerequisites
- API requirements documented
- Data models designed
- Development environment set up
- Database schema approved

## Steps

### 1. API Design
- **Command**: None (manual design)
- **Agent**: `api-review-agent` (for reviewing design)
- **Validation**: API design approved
- **Output**: OpenAPI specification document

### 2. Database Schema
- **Command**: None (manual design)
- **Agent**: None
- **Validation**: Schema normalized and indexed
- **Output**: Migration files

### 3. Scaffold Service
- **Command**: `service-create`
- **Agent**: None
- **Validation**: Service compiles and runs
- **Output**: Service structure

### 4. Implement Models
- **Command**: None (manual implementation)
- **Agent**: None
- **Validation**: Models match schema
- **Output**: Model files

### 5. Implement Repository
- **Command**: None (manual implementation)
- **Agent**: `api-review-agent` (for reviewing queries)
- **Validation**: Queries are optimized
- **Output**: Repository layer

### 6. Implement Service Layer
- **Command**: None (manual implementation)
- **Agent**: None
- **Validation**: Business logic is correct
- **Output**: Service files

### 7. Implement Handlers
- **Command**: None (manual implementation)
- **Agent**: `api-review-agent` (for review)
- **Validation**: Handlers follow REST conventions
- **Output**: Handler files

### 8. Write Tests
- **Command**: `test-generate`
- **Agent**: None
- **Validation**: All tests pass with adequate coverage
- **Output**: Test files

### 9. Update Documentation
- **Command**: None (manual update)
- **Agent**: None
- **Validation**: OpenAPI spec is complete
- **Output**: Updated API documentation

### 10. Code Review
- **Command**: None (PR creation)
- **Agent**: `api-review-agent` (automated review)
- **Validation**: Review feedback addressed
- **Output**: Approved pull request

### 11. Integration Testing
- **Command**: None (manual testing)
- **Agent**: None
- **Validation**: API works in integrated environment
- **Output**: Test report

### 12. Deployment
- **Command**: `deploy-service`
- **Agent**: None
- **Validation**: Service deployed successfully
- **Output**: Deployment confirmation

## Outputs

| Output | Description | Format |
|--------|-------------|--------|
| OpenAPI spec | API specification | YAML |
| Migration files | Database migrations | SQL/Go |
| Service code | Implemented service | Go |
| Test results | Test execution results | JSON |
| API docs | API documentation | Markdown/OpenAPI |

## Estimated Time

- Simple CRUD API: 1-2 days
- Medium complexity: 3-5 days
- Complex API: 1-2 weeks

## Tags
`backend`, `workflow`, `api-development`, `rest`, `agile`
