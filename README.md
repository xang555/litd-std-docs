# Company Development Standards

This repository contains the official development standards, guidelines, and shared resources for all engineering teams. It serves as the content source for the LitD MCP Server.

## Table of Contents

- [Overview](#overview)
- [Structure](#structure)
- [Content Types](#content-types)
- [Teams](#teams)
- [Quick Start](#quick-start)
- [Contributing](#contributing)
- [Documentation](#documentation)
- [License](#license)

## Overview

This is a **company standards repository** that provides:

- **Engineering Standards** - Coding conventions and best practices
- **Learning Resources** - Skill definitions for career development
- **AI Agents** - Specialized agent behaviors for different tasks
- **Reusable Templates** - Code templates with embedded examples
- **Workflows** - Multi-step procedures and guides
- **Documentation** - Team-specific and shared documentation

Content is organized by **team** (frontend, backend, mobile, devops, data) and **shared** (universal resources).

## Structure

```
teams/                          # Team-specific content
├── {team}/
│   ├── skills/                 # Team-specific learning resources
│   ├── agents/                 # Team-specific AI agents
│   ├── commands/               # Team-specific command templates
│   ├── workflows/              # Team-specific multi-step workflows
│   ├── standards/              # Team coding standards (individual files)
│   ├── templates/              # Reusable code templates (.md with code)
│   ├── docs/                   # Team documentation
│   │   ├── getting-started.md
│   │   ├── best-practices.md
│   │   └── troubleshooting.md
│   └── VERSION.md              # Directory version tracking

shared/                         # Cross-team shared content
├── skills/                     # Universal skills (all .md files)
│   ├── problem-solving.md
│   ├── debugging.md
│   ├── testing-fundamentals.md
│   ├── code-review.md
│   ├── documentation.md
│   ├── communication.md
│   ├── design-patterns.md
│   ├── solid-principles.md
│   ├── clean-code.md
│   ├── secure-coding.md
│   ├── data-protection.md
│   ├── authentication.md
│   ├── optimization.md
│   ├── scalability.md
│   ├── monitoring.md
│   ├── git-workflow.md
│   ├── ide-setup.md
│   ├── cli-commands.md
│   ├── git-commit.md
│   ├── prompt-engineering.md
│   ├── code-generation.md
│   ├── review-automation.md
│   └── technical-writing.md
├── agents/                     # Universal agents
├── commands/                   # Shared commands
├── workflows/                  # Cross-team workflows
├── standards/                  # Cross-cutting standards
├── templates/                  # Universal templates (.md with code)
│   ├── license/                # License templates
│   ├── dockerfile.md           # Docker build templates
│   ├── gitignore.md            # .gitignore templates
│   ├── api-response-format.md  # API response patterns
│   ├── logging-pattern.md      # Structured logging
│   └── ...
├── docs/                       # Shared documentation
│   ├── architecture.md         # Repository architecture
│   ├── contributing.md         # Contribution guidelines
│   └── migration-guide.md      # Migration instructions
├── VERSION.md                  # Directory version tracking
└── ...
```

## Content Types

| Type | Description | Format | Location |
|------|-------------|--------|----------|
| **Skills** | Learning objectives by experience level | `.md` with Junior/Mid-Level/Senior sections | `teams/{team}/skills/` or `shared/skills/` |
| **Agents** | AI agent system prompts and capabilities | `.md` with config and instructions | `teams/{team}/agents/` or `shared/agents/` |
| **Commands** | Reusable CLI command templates | `.md` with usage examples | `teams/{team}/commands/` or `shared/commands/` |
| **Workflows** | Multi-step procedures and guides | `.md` with prerequisites and steps | `teams/{team}/workflows/` or `shared/workflows/` |
| **Standards** | Coding standards and conventions | `.md` with rules and examples | `teams/{team}/standards/` or `shared/standards/` |
| **Templates** | Reusable code and file templates | `.md` with embedded code blocks | `teams/{team}/templates/` or `shared/templates/` |
| **Docs** | Documentation and guides | `.md` | `teams/{team}/docs/` or `shared/docs/` |
| **VERSION** | Directory version tracking | `.md` with directory table | `teams/{team}/VERSION.md` or `shared/VERSION.md` |

## Teams

| Team | Focus | Tech Stack |
|------|-------|------------|
| **Frontend** | Web applications | React, TypeScript, Next.js, Tailwind |
| **Backend** | API and services | Go, PostgreSQL, Redis, gRPC |
| **Mobile** | iOS & Android | Swift, Kotlin, Flutter |
| **DevOps** | Infrastructure | Kubernetes, Terraform, CI/CD, Docker |
| **Data** | Data pipelines | Python, Spark, Airflow, dbt |

## Quick Start

### Clone the Repository

```bash
git clone <repository-url>
cd litd-std-docs
```

### Explore Team-Specific Content

```bash
# Frontend resources
ls teams/frontend/skills/
ls teams/frontend/templates/

# Backend resources
ls teams/backend/standards/
ls teams/backend/agents/

# DevOps resources
ls teams/devops/workflows/
```

### Explore Shared Content

```bash
# Universal skills
ls shared/skills/core/
ls shared/skills/collaboration/

# Shared templates
ls shared/templates/

# Documentation
cat shared/docs/architecture.md
```

### Using with MCP Server

This repository serves as content for the LitD MCP Server. Configure your MCP client to point to this repository to access:

- Skill definitions for developer learning
- Agent prompts for AI assistance
- Code templates for rapid development
- Standards reference for code review

## Contributing

### Process

1. **Determine content scope** - Is this team-specific or shared?
2. **Follow content specifications** - Use the correct format for the content type
3. **Create or update files** - Place in appropriate directory
4. **Update VERSION.md** - Track changes in the relevant VERSION.md file
5. **Submit pull request** - Get approval from relevant team lead or Standards Team

### Content Type Specifications

All content types have specific formats defined in `CLAUDE.md`. Key requirements:

- **Templates**: Must be `.md` files with embedded code blocks (not raw source files)
- **Standards**: Individual files in `standards/` subdirectory (no `standards.md` at root)
- **Skills**: Include Junior, Mid-Level, and Senior sections
- **Agents**: Include Configuration, Instructions, Capabilities, and Tools Required
- **All files**: Use kebab-case naming

### Commit Convention

Use **Conventional Commits** format:

```
<type>[optional scope]: <description>

[optional body]

[optional footer]
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`

Examples:
- `feat(frontend): add React hooks skill definition`
- `fix(templates): correct Dockerfile example`
- `docs(readme): update quick start guide`

## Documentation

### Repository Documentation

- **Architecture** - `shared/docs/architecture.md` - Repository structure and design
- **Contributing** - `shared/docs/contributing.md` - Contribution guidelines
- **Migration Guide** - `shared/docs/migration-guide.md` - Updating between versions

### Team Documentation

- **Getting Started** - `teams/{team}/docs/getting-started.md`
- **Best Practices** - `teams/{team}/docs/best-practices.md`
- **Troubleshooting** - `teams/{team}/docs/troubleshooting.md`

### Configuration Files

- **CLAUDE.md** - Guidance for Claude Code working in this repository
- **VERSION.md** - Version tracking for each team and shared directory

## Accessing Standards

Content is available through:

- **LitD MCP Server** - Primary access method for AI-assisted development
- **Git Repository** - Direct access for reference and contribution
- **Documentation Site** - Published documentation (if configured)

## Version Control

- **Main branch** (`main`) reflects current production standards
- **Semantic versioning** for major updates
- **Breaking changes** documented in migration guides
- **Conventional commits** for change tracking

## License

See [LICENSE](LICENSE) for details.

## Contact

For questions or suggestions:
- **Engineering Standards Team** - standards@company.com
- **Issue Tracker** - Report issues or requests
- **Discussions** - Community discussions and Q&A
