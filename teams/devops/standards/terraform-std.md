# Terraform Standards

## Scope
This standard applies to all Terraform configurations used for infrastructure provisioning.

## Rules

### 1. Module Structure

**Rule:** Organize infrastructure into reusable modules.

**Rationale:** Modules promote reusability and consistency.

**Example:**
```
modules/
├── vpc/
│   ├── main.tf
│   ├── variables.tf
│   ├── outputs.tf
│   └── README.md
├── ec2/
└── rds/

environments/
├── dev/
│   ├── main.tf
│   └── terraform.tfvars
├── staging/
└── production/
```

### 2. Naming Conventions

**Rule:** Use consistent naming conventions across resources.

**Rationale:** Predictable names improve readability and automation.

**Example:**
```hcl
# Good
resource "aws_instance" "web_server" {
  name = "web-server-${var.environment}"
}

# Bad
resource "aws_instance" "web1" {
  name = "server1"
}
```

### 3. Variable Organization

**Rule:** Define variables in variables.tf with descriptions and types.

**Rationale:** Self-documenting and type-safe configurations.

**Example:**
```hcl
variable "instance_count" {
  description = "Number of instances to provision"
  type        = number
  default     = 3

  validation {
    condition     = var.instance_count > 0 && var.instance_count <= 10
    error_message = "Instance count must be between 1 and 10."
  }
}
```

### 4. State Management

**Rule:** Use remote backend with state locking.

**Rationale:** Prevents state corruption and enables collaboration.

**Example:**
```hcl
terraform {
  backend "s3" {
    bucket         = "company-terraform-state"
    key            = "apps/web/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-state-lock"
  }
}
```

### 5. Secrets Management

**Rule:** Never store secrets in Terraform files. Use vault or environment variables.

**Rationale:** Prevents accidental exposure of sensitive data.

**Example:**
```hcl
# Good
resource "aws_db_instance" "default" {
  password = var.db_password  # From secure source
}

# Bad
resource "aws_db_instance" "default" {
  password = "SuperSecret123!"  # Never do this
}
```

### 6. Resource Tagging

**Rule:** Apply consistent tags to all resources.

**Rationale:** Enables cost allocation and resource management.

**Example:**
```hcl
resource "aws_instance" "web" {
  tags = {
    Name        = "web-${var.environment}"
    Environment = var.environment
    ManagedBy   = "Terraform"
    Owner       = "team-platform"
    CostCenter  = "engineering"
  }
}
```

### 7. Provider Versioning

**Rule:** Pin provider versions.

**Rationale:** Ensures consistent deployments and prevents breaking changes.

**Example:**
```hcl
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  required_version = ">= 1.0"
}
```

### 8. Output Organization

**Rule:** Define useful outputs in outputs.tf with descriptions.

**Rationale:** Makes important values easily accessible.

**Example:**
```hcl
output "web_instance_public_ip" {
  description = "Public IP of the web instance"
  value       = aws_instance.web.public_ip
}

output "database_endpoint" {
  description = "RDS instance endpoint"
  value       = aws_db_instance.main.endpoint
  sensitive   = true
}
```

### 9. Documentation

**Rule:** Every module must have a README.md with usage examples.

**Rationale:** Enables module discovery and proper usage.

**Example:**
```markdown
# VPC Module

Creates a VPC with public and private subnets.

## Usage

\`\`\`hcl
module "vpc" {
  source = "../modules/vpc"

  cidr_block           = "10.0.0.0/16"
  availability_zones   = ["us-east-1a", "us-east-1b"]
  enable_nat_gateway   = true
  single_nat_gateway   = true
}
\`\`\`

## Requirements

No requirements.

## Inputs

| Name | Description | Type | Default |
|------|-------------|------|---------|
```

### 10. Lifecycle Configuration

**Rule:** Configure lifecycle to prevent accidental destruction.

**Rationale:** Prevents data loss and accidental resource deletion.

**Example:**
```hcl
resource "aws_db_instance" "production" {
  # ... configuration ...

  lifecycle {
    prevent_destroy = true
    ignore_changes = [
      # Allow manual changes to security groups
      security_groups,
    ]
  }
}
```

## Enforcement

### Linting
- `terraform fmt` for formatting
- `tflint` for linting
- `tfsec` for security scanning

### Review
All pull requests must:
- Pass `terraform fmt -check`
- Pass `tflint`
- Pass `tfsec` with no critical issues
- Be reviewed by DevOps team

### Pre-commit Hooks
- `terraform fmt -recursive`
- `terraform init`
- `terraform validate`

## Exceptions

Exceptions require:
- DevOps lead approval
- Security review for critical resources
- Documentation of exception rationale

## Related Standards
- AWS Standards
- Kubernetes Standards
- Security Standards

## Tags
`devops`, `terraform`, `infrastructure`, `standards`, `iac`
