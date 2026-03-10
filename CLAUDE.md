# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

This is a **company standards repository** that serves as the content source for the LitD MCP Server. It contains engineering standards, guidelines, templates, skill definitions, agents, commands, workflows, and documentation organized by team and category.

## Architecture

### Content Structure

```
teams/                          # Team-specific content
  ├── {team}/
  │   ├── skills/               # Team-specific learning resources
  │   ├── agents/               # Team-specific AI agents
  │   ├── commands/             # Team-specific command templates
  │   ├── workflows/            # Team-specific multi-step workflows
  │   ├── standards/            # Team coding standards
  │   ├── templates/            # Reusable code templates
  │   ├── docs/                 # Team documentation
  │   │   ├── getting-started.md
  │   │   ├── best-practices.md
  │   │   └── troubleshooting.md
  │   ├── VERSION.md            # Directory version tracking
  │   └── team-config.toml      # Team-specific configuration

shared/                         # Cross-team shared content
  ├── skills/                   # Universal skills (all .md files)
  │   ├── problem-solving.md    # Problem-solving techniques
  │   ├── debugging.md          # Debugging strategies
  │   ├── testing-fundamentals.md
  │   ├── code-review.md        # Code review practices
  │   ├── documentation.md      # Documentation skills
  │   ├── communication.md      # Technical communication
  │   ├── design-patterns.md    # Design patterns
  │   ├── solid-principles.md   # SOLID principles
  │   ├── clean-code.md         # Clean code practices
  │   ├── secure-coding.md      # Security best practices
  │   ├── data-protection.md    # Data protection
  │   ├── authentication.md     # Authentication methods
  │   ├── optimization.md       # Performance optimization
  │   ├── scalability.md        # Scalability patterns
  │   ├── monitoring.md         # Monitoring strategies
  │   ├── git-workflow.md       # Git workflow
  │   ├── ide-setup.md          # IDE setup
  │   ├── cli-commands.md       # CLI commands
  │   ├── git-commit.md         # Conventional commits
  │   ├── prompt-engineering.md # Prompt engineering
  │   ├── code-generation.md    # AI code generation
  │   ├── review-automation.md  # Review automation
  │   ├── technical-writing.md  # Technical writing
  │   └── .skills-config.toml   # Skills configuration
  ├── agents/                   # Universal agents
  ├── commands/                 # Shared commands
  ├── workflows/                # Cross-team workflows
  ├── standards/                # Cross-cutting standards
  ├── templates/                # Universal templates
  ├── docs/                     # Shared documentation
  │   ├── architecture.md
  │   ├── contributing.md
  │   └── migration-guide.md
  ├── VERSION.md                # Directory version tracking
  └── ...
```

### Configuration

The repository uses a hierarchical organization system where team-specific content in `teams/` can override or extend shared content in `shared/`.

## Content Type Specifications

### Skills

All skill files must follow the standard format:

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

**Location:** `teams/{team}/skills/` OR `shared/skills/`

### Agents

```markdown
# Agent Name

## Description
[Purpose and capabilities]

## Configuration
- Model: claude-opus-4-6
- Temperature: 0.7
- Max Tokens: 4000

## Instructions
[System prompt and behavior guidelines]

## Capabilities
- Capability 1

## Tools Required
- tool1

## Tags
`tag1`, `tag2`
```

**Location:** `teams/{team}/agents/` OR `shared/agents/`

### Commands

```markdown
# Command Name

## Description
[What the command does]

## Usage
\`\`\`bash
command-template [options]
\`\`\`

## Parameters
| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| param1    | string | true    | -       | Parameter description |
| param2    | number | false   | 42      | Optional parameter |

## Examples
\`\`\`bash
# Basic usage
command-name --param1 value

# With all options
command-name --param1 value --param2 100
\`\`\`

## Tags
`tag1`, `tag2`
```

**Location:** `teams/{team}/commands/` OR `shared/commands/`

### Workflows

```markdown
# Workflow Name

## Description
[What the workflow accomplishes]

## Prerequisites
- Prerequisite 1

## Steps
1. **Step 1**: [Description]
   - Command: `command-name`
   - Agent: `agent-name`

## Tags
`tag1`, `tag2`
```

**Location:** `teams/{team}/workflows/` OR `shared/workflows/`

### Standards

```markdown
# Standard Name

## Scope
[What this standard applies to]

## Rules
1. **Rule 1**: [Description]
   - Rationale: [Why this rule exists]
   - Example: [Code example showing correct usage]

## Enforcement
- Linting: [Tool configuration]
- CI: [Automated checks]

## Tags
`tag1`, `tag2`
```

**Location:** `teams/{team}/standards/` OR `shared/standards/`

**Note:** Each standard should be a separate file. Do not create `standards.md` at team root level.

### Templates

Templates are `.md` documentation files that contain embedded, reusable code examples:

```markdown
# Template Name

## Description
[Purpose of the template]

## Usage
[Brief instructions on how to use this template]

## Template
\`\`\`language
[Reusable code example]
\`\`\`

## Additional Examples
[More code variants or examples]

## Tags
`tag1`, `tag2`
```

**Location:** `teams/{team}/templates/` OR `shared/templates/`

**Important:** Templates must be `.md` files with embedded code blocks, NOT raw source code files. This ensures the MCP Server can properly index and serve the content.

### Version Tracking

Each team and shared directory contains a `VERSION.md` file for tracking updates:

```markdown
# {Team/Shared} - Version Tracking

| Directory | Notes | Last Updated |
|-----------|-------|--------------|
| skills/    | Brief description | YYYY-MM-DD |
| agents/    | Brief description | YYYY-MM-DD |
| ...
```

**Location:** `teams/{team}/VERSION.md` OR `shared/VERSION.md`

**Purpose:** Track when each subdirectory was last updated with new content or changes.

**Update:** When adding/modifying content in a directory, update the corresponding entry in the VERSION.md table.

## File Naming Conventions

- **Skills:** `kebab-case.md`
- **Agents:** `kebab-case-agent.md`
- **Commands:** `kebab-case-cmd.md`
- **Workflows:** `kebab-case-workflow.md`
- **Standards:** `kebab-case-std.md`
- **Templates:** Follow target file convention
- **Docs:** `kebab-case.md`
- **Configs:** `*.toml`

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

When adding or modifying content:
1. Determine if content is team-specific or shared
2. Update the relevant team or shared directory
3. Follow the content type specifications exactly
4. **Update the `VERSION.md` file** in the relevant team or shared directory
5. Use conventional commits for all changes
