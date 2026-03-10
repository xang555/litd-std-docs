# Deploy Command

## Description
Command template for deploying applications to Kubernetes with proper validation and rollback.

## Usage
```
/deploy <Environment> <Service> [options]
```

## Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| Environment | string | true | - | Target environment: dev, staging, production |
| Service | string | true | - | Service name to deploy |
| --image | string | false | auto | Docker image tag |
| --namespace | string | false | default | Kubernetes namespace |
| --timeout | int | false | 300 | Deployment timeout in seconds |
| --dry-run | boolean | false | false | Validate without deploying |
| --verify | boolean | false | true | Verify deployment health |

## Examples

### Deploy to development
```bash
/deploy dev user-service
```
Deploys user-service to development environment

### Deploy with specific image
```bash
/deploy production order-service --image=v2.3.1
```
Deploys specific version to production

### Validate deployment without deploying
```bash
/deploy staging payment-service --dry-run
```
Validates configuration without deploying

### Deploy to custom namespace
```bash
/deploy dev api-gateway --namespace=team-a
```
Deploys to specific namespace

## Deployment Process

1. **Pre-deployment Checks**
   - Validates environment exists
   - Checks service configuration
   - Verifies image exists in registry

2. **Backup Current State**
   - Saves current deployment config
   - Records running pods

3. **Apply New Configuration**
   - Updates deployment with new image
   - Applies ConfigMaps and Secrets
   - Creates or updates Services

4. **Health Verification**
   - Waits for rollout completion
   - Checks pod health
   - Runs smoke tests

5. **Rollback on Failure**
   - Automatic rollback if health checks fail
   - Reverts to previous stable version

## Output

```
✓ Pre-deployment checks passed
✓ Current deployment backed up
✓ Applying new configuration...
✓ Waiting for rollout completion...
  - Replicas: 3/3 ready
✓ Health checks passed
✓ Deployment successful

Service: user-service
Environment: development
Image: user-service:v2.3.0
Replicas: 3/3 ready
URL: https://user-service.dev.example.com
```

## Rollback

If deployment fails, automatic rollback is triggered:

```bash
✗ Health checks failed
✓ Rolling back to previous version...
✓ Rollback complete
```

## Tags
`devops`, `deploy`, `kubernetes`, `cli`, `cicd`
