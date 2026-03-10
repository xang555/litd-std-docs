# GitHub Actions CI/CD Pipeline

## Description

Production-ready GitHub Actions workflow for building, testing, and deploying applications with security scanning and multi-environment support.

## Complete Pipeline Template

```yaml
name: Deploy Application

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

env:
  AWS_REGION: us-east-1
  ECR_REGISTRY: ${{ secrets.AWS_ECR_REGISTRY }}
  ECR_REPOSITORY: myapp
  KUBERNETES_CLUSTER: production

jobs:
  build:
    name: Build and Push Docker Image
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v2
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ env.AWS_REGION }}

      - name: Login to Amazon ECR
        id: login-ecr
        uses: aws-actions/amazon-ecr-login@v1

      - name: Build and tag image
        run: |
          docker build -t $ECR_REGISTRY/$ECR_REPOSITORY:$GITHUB_SHA .
          docker tag $ECR_REGISTRY/$ECR_REPOSITORY:$GITHUB_SHA $ECR_REGISTRY/$ECR_REPOSITORY:latest

      - name: Push image to Amazon ECR
        run: |
          docker push $ECR_REGISTRY/$ECR_REPOSITORY:$GITHUB_SHA
          docker push $ECR_REGISTRY/$ECR_REPOSITORY:latest

      - name: Run security scan
        uses: aquasecurity/trivy-action@master
        with:
          image-ref: ${{ env.ECR_REGISTRY }}/${{ env.ECR_REPOSITORY }}:${{ github.sha }}
          format: 'sarif'
          output: 'trivy-results.sarif'

      - name: Upload Trivy results to GitHub Security tab
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: 'trivy-results.sarif'

  deploy-staging:
    name: Deploy to Staging
    needs: build
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    environment:
      name: staging
      url: https://staging.example.com
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Configure kubectl
        run: |
          aws eks update-kubeconfig --name ${{ env.KUBERNETES_CLUSTER }}-staging --region ${{ env.AWS_REGION }}

      - name: Deploy to Kubernetes
        run: |
          kubectl set image deployment/myapp \
            myapp=$ECR_REGISTRY/$ECR_REPOSITORY:$GITHUB_SHA \
            -n staging

      - name: Verify deployment
        run: |
          kubectl rollout status deployment/myapp -n staging --timeout=5m

      - name: Run smoke tests
        run: |
          ./scripts/smoke-tests.sh https://staging.example.com

  deploy-production:
    name: Deploy to Production
    needs: [build, deploy-staging]
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    environment:
      name: production
      url: https://example.com
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Configure kubectl
        run: |
          aws eks update-kubeconfig --name ${{ env.KUBERNETES_CLUSTER }}-production --region ${{ env.AWS_REGION }}

      - name: Deploy to Kubernetes
        run: |
          kubectl set image deployment/myapp \
            myapp=$ECR_REGISTRY/$ECR_REPOSITORY:$GITHUB_SHA \
            -n production

      - name: Verify deployment
        run: |
          kubectl rollout status deployment/myapp -n production --timeout=10m

      - name: Run smoke tests
        run: |
          ./scripts/smoke-tests.sh https://example.com

      - name: Notify deployment
        if: success()
        run: |
          echo "Deployment successful! Notify team..."
          # Add Slack/Teams notification here
```

## Test Job Template

```yaml
test:
  name: Run Tests
  runs-on: ubuntu-latest
  steps:
    - name: Checkout code
      uses: actions/checkout@v3

    - name: Set up Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '20'
        cache: 'npm'

    - name: Install dependencies
      run: npm ci

    - name: Run linter
      run: npm run lint

    - name: Run type check
      run: npm run type-check

    - name: Run unit tests
      run: npm run test:unit

    - name: Run integration tests
      run: npm run test:integration

    - name: Generate coverage report
      run: npm run test:coverage

    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        files: ./coverage/lcov.info
        flags: unittests
        name: codecov-umbrella
```

## Docker Multi-stage Build Template

```dockerfile
# Build stage
FROM node:20-alpine AS builder
WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm ci

# Copy source and build
COPY . .
RUN npm run build

# Production stage
FROM node:20-alpine AS production
WORKDIR /app

# Copy built files
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY package*.json ./

# Create non-root user
RUN addgroup -g 1001 -S nodejs && \
    adduser -S nodejs -u 1001
USER nodejs

EXPOSE 3000

CMD ["node", "dist/main.js"]
```

## Kubernetes Deployment Template

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
  labels:
    app: myapp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: ${ECR_REGISTRY}/${ECR_REPOSITORY}:${IMAGE_TAG}
        ports:
        - containerPort: 3000
        env:
        - name: NODE_ENV
          value: "production"
        - name: PORT
          value: "3000"
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 3000
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 3000
          initialDelaySeconds: 5
          periodSeconds: 5

---
apiVersion: v1
kind: Service
metadata:
  name: myapp-service
spec:
  selector:
    app: myapp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 3000
  type: LoadBalancer
```

## Terraform Infrastructure Template

```hcl
# ECR Repository
resource "aws_ecr_repository" "app" {
  name                 = var.app_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = var.common_tags
}

# EKS Cluster
resource "aws_eks_cluster" "main" {
  name     = var.cluster_name
  role_arn = aws_iam_role.cluster.arn
  version  = var.kubernetes_version

  vpc_config {
    subnet_ids = var.subnet_ids
  }

  tags = var.common_tags
}

# Load Balancer
resource "aws_lb" "main" {
  name               = "${var.app_name}-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb.id]
  subnets            = var.public_subnet_ids

  tags = var.common_tags
}
```

## Tags

`devops`, `cicd`, `github-actions`, `docker`, `kubernetes`, `deployment`, `pipeline`
