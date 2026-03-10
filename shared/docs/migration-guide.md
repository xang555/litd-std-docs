# Migration Guide

## Overview

This guide helps with migrating content and configurations to the new repository structure.

## Structure Changes

### Before
```
litd-std-docs/
├── skills/             # Global skills
├── agents/             # Global agents
├── shared/             # Cross-team standards
└── teams/              # Team standards only
```

### After
```
litd-std-docs/
├── shared/             # All shared content
│   ├── skills/
│   ├── agents/
│   ├── commands/
│   ├── workflows/
│   ├── standards/
│   ├── templates/
│   └── docs/
└── teams/              # Team-specific content
    └── {team}/
        ├── skills/
        ├── agents/
        ├── commands/
        ├── workflows/
        ├── standards/
        ├── templates/
        ├── docs/
        └── team-config.toml
```

## Migration Steps

### 1. Categorize Existing Content

Determine if content should be:
- **Shared**: Universal across all teams
- **Team-specific**: Specific to a team's tech stack

### 2. Move Content

#### Skills
- Universal skills → `shared/skills/`
- Team-specific skills → `teams/{team}/skills/`

#### Agents
- Universal agents → `shared/agents/`
- Team-specific agents → `teams/{team}/agents/`

#### Standards
- Cross-cutting standards → `shared/standards/`
- Team standards → `teams/{team}/standards/`

### 3. Update Configurations

Create or update:
- `shared/shared-config.toml` for shared content
- `teams/{team}/team-config.toml` for team content

### 4. Validate

Run validation to ensure:
- All files follow content type specifications
- All configurations are valid TOML
- All references exist

## Rollback

If migration fails, restore from git:
```bash
git reset --hard HEAD
```

## Support

For issues or questions, contact the Standards Team.
