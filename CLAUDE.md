# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

This is a **company standards repository** that serves as the content source for the LitD MCP Server. It contains engineering standards, guidelines, templates, and skill definitions organized by team and category.

## Architecture

### Content Structure

```
teams/          # Team-specific standards and code templates
  ├── frontend/ # React/TypeScript standards
  ├── backend/  # Go standards
  ├── mobile/   # iOS/Android standards
  ├── devops/   # Infrastructure standards
  └── data/     # Data engineering standards

skills/         # Learning resources with level-based progression
  ├── .skills-config.toml  # Active config used by MCP server
  ├── core/     # Problem-solving, debugging, testing
  ├── collaboration/  # Code review, documentation
  ├── architecture/    # Design patterns, SOLID
  ├── security/  # Secure coding, data protection
  ├── performance/  # Optimization, monitoring
  ├── tools/    # Git workflow, CLI
  └── ai-assisted-development/  # Prompt engineering

shared/         # Cross-team standards (API versioning, error handling, logging)
agents/         # AI agent instructions for code review, testing, documentation
```

### MCP Integration

The `.skills-config.toml` file is **actively parsed by the LitD MCP Server** to provide:
- Skill metadata (display names, tags, required flags)
- Category groupings
- Level-based skill requirements (junior → principal)

When modifying skills, always update the corresponding TOML configuration to ensure the MCP server returns accurate metadata.

### Skill Document Format

All skill files (`skills/**/*.md`) must follow the standard format:

```markdown
# Skill Name

## Description
[Brief description]

## Level: Junior
### Learning Objectives
### Required Skills
### Resources

## Level: Mid-Level
### Learning Objectives
### Required Skills
### Optional Skills

## Level: Senior
### Learning Objectives
### Required Skills

## Tags
`tag1`, `tag2`, `tag3`
```

## Commit Conventions

This repository uses **Conventional Commits** format:
```
<type>[optional scope]: <description>

<body>

<footer>
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`

Example: `refactor(tools): standardize git-commit skill format`

## Contributing

When adding or modifying standards:
1. Update the relevant category file
2. If adding a skill, register it in `.skills-config.toml`
3. Follow the skill document format exactly
4. Use conventional commits for all changes
