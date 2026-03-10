# DevOps Team - Getting Started

## Welcome to the DevOps Team

This guide helps you get started with DevOps practices at our company.

## Tech Stack

- **Container Orchestration**: Kubernetes
- **Infrastructure as Code**: Terraform
- **CI/CD**: GitHub Actions
- **Monitoring**: Prometheus, Grafana
- **Logging**: ELK Stack

## Setup

1. Configure kubectl:
   ```bash
   aws eks update-kubeconfig --name production
   ```

2. Install Terraform:
   ```bash
   brew install terraform
   ```

3. Configure AWS CLI:
   ```bash
   aws configure
   ```

## Project Structure

```
infrastructure/
├── terraform/
│   ├── modules/       # Reusable modules
│   ├── environments/  # Environment configs
│   └── examples/      # Usage examples
├── kubernetes/
│   ├── base/          # Base manifests
│   └── overlays/      # Environment overlays
└── scripts/           # Utility scripts
```

## Code Style

Follow the coding standards defined in `teams/devops/standards/`.

## Resources

- Team-specific agents: `teams/devops/agents/`
- Team workflows: `teams/devops/workflows/`
- Shared resources: `shared/`
