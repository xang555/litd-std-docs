# Architecture Overview

## Repository Architecture

This repository serves as the content source for the LitD MCP Server, providing a centralized, team-centric structure for AI-assisted development standards.

## Directory Structure

```
litd-std-docs/
├── teams/              # Team-specific content
│   ├── {team}/
│   │   ├── skills/     # Team-specific learning resources
│   │   ├── agents/     # Team-specific AI agents
│   │   ├── commands/   # Team-specific command templates
│   │   ├── workflows/  # Team-specific multi-step workflows
│   │   ├── standards/  # Team coding standards
│   │   ├── templates/  # Reusable code templates
│   │   ├── docs/       # Team documentation
│   │   └── team-config.toml
│
├── shared/             # Cross-team shared content
│   ├── skills/         # Universal skills
│   ├── agents/         # Universal agents
│   ├── commands/       # Shared commands
│   ├── workflows/      # Cross-team workflows
│   ├── standards/      # Cross-cutting standards
│   ├── templates/      # Universal templates
│   ├── docs/           # Shared documentation
│   └── shared-config.toml
```

## Content Types

### Skills
Define learning objectives and requirements for specific competencies.

### Agents
Define AI agent behavior, instructions, and capabilities.

### Commands
Reusable command templates for common operations.

### Workflows
Multi-step procedures combining multiple commands and agents.

### Standards
Coding standards and conventions.

### Templates
Reusable code and file templates.

## Configuration

Content is configured through TOML files:
- `shared/shared-config.toml` - Shared configuration
- `teams/{team}/team-config.toml` - Team-specific configuration

## MCP Integration

The MCP server consumes content from this repository, providing:
- Team context awareness
- Content indexing and search
- Configuration merging (team overrides shared)
