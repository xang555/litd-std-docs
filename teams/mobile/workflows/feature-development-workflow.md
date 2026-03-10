# Mobile Feature Development Workflow

## Description
A structured workflow for developing mobile features from design to app store submission.

## Prerequisites
- Feature requirements documented
- UI/UX designs approved
- Development environment set up
- Backend APIs available (if needed)

## Steps

### 1. Requirements Review
- **Command**: None
- **Agent**: None
- **Validation**: Requirements clear for both platforms
- **Output**: Approved requirements document

### 2. UI Design Review
- **Command**: None
- **Agent**: `mobile-review-agent`
- **Validation**: Design follows platform conventions
- **Output**: Design feedback document

### 3. Scaffold Screens
- **Command**: `screen-create`
- **Agent**: None
- **Validation**: Screens compile and run
- **Output**: Screen files for both platforms

### 4. Implement ViewModels
- **Command**: None
- **Agent**: `mobile-review-agent` (iterative review)
- **Validation**: ViewModels handle business logic
- **Output**: ViewModel implementations

### 5. Implement UI
- **Command**: None
- **Agent**: `mobile-review-agent` (iterative review)
- **Validation**: UI matches design, responsive
- **Output**: Screen implementations

### 6. Implement Data Layer
- **Command**: None
- **Agent**: None
- **Validation**: Repositories handle data correctly
- **Output**: Repository implementations

### 7. Add Navigation
- **Command**: None
- **Agent**: None
- **Validation**: Navigation flows work correctly
- **Output**: Navigation configuration

### 8. Write Tests
- **Command**: `test-generate`
- **Agent**: None
- **Validation**: Tests pass with adequate coverage
- **Output**: Test files

### 9. Code Review
- **Command**: None
- **Agent**: `mobile-review-agent`
- **Validation**: Review feedback addressed
- **Output**: Approved pull requests

### 10. Device Testing
- **Command**: None
- **Agent**: None
- **Validation**: Works on physical devices
- **Output**: Test report

### 11. Beta Release
- **Command**: `deploy-beta`
- **Agent**: None
- **Validation**: Beta builds distributed
- **Output**: Beta release notes

## Outputs

| Output | Description | Format |
|--------|-------------|--------|
| Requirements | Feature requirements | Markdown |
| Design assets | UI designs and assets | Figma/Sketch |
| Screen files | Implemented screens | Swift/Kotlin |
| Test results | Test execution results | JSON/Console |
| Beta builds | Test builds | IPA/APK |

## Estimated Time

- Simple screen: 2-3 days
- Medium feature: 1 week
- Complex feature: 2-3 weeks

## Tags
`mobile`, `workflow`, `feature-development`, `ios`, `android`
