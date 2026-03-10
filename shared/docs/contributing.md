# Contributing Guide

## Overview

This repository contains engineering standards, guidelines, and shared resources for all teams. Contributions are welcome from all team members.

## Getting Started

1. Clone the repository
2. Create a branch for your changes
3. Make your changes following the content type specifications
4. Submit a pull request for review

## Content Type Specifications

All content must follow the standard formats defined in the architecture documentation.

### Skills Format
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

### Agents Format
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

## Commit Conventions

Use Conventional Commits format:
```
<type>[optional scope]: <description>
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`

## Review Process

1. All changes require review
2. Team-specific content: reviewed by Team Lead
3. Shared content: reviewed by Standards Team
4. Get approval before merging

## File Naming

- Use kebab-case for all filenames
- Skills: `kebab-case.md`
- Agents: `kebab-case-agent.md`
- Commands: `kebab-case-cmd.md`
- Workflows: `kebab-case-workflow.md`
- Standards: `kebab-case-std.md`
- Docs: `kebab-case.md`
- Configs: `*.toml`
