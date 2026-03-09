---
description: Commit all changes with conventional commit messages
---

# git-commit

You are an expert at creating conventional commits following the specification at https://www.conventionalcommits.org/en/v1.0.0-beta.4/#specification

## Your Task

Commit all staged and unstaged changes in the repository following the Conventional Commits specification.

## Conventional Commits Format

A conventional commit structure:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Commit Types

- **feat**: ✨ A new feature
- **fix**: 🐛 A bug fix
- **docs**: 📝 Documentation only changes
- **style**: 💄 Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **refactor**: ♻️ A code change that neither fixes a bug nor adds a feature
- **perf**: ⚡ A code change that improves performance
- **test**: ✅ Adding missing tests or correcting existing tests
- **build**: 📦 Changes that affect the build system or external dependencies
- **ci**: 👷 Changes to CI configuration files and scripts
- **chore**: 🔧 Other changes that don't modify src or test files
- **revert**: ↩️ Reverts a previous commit

### Scope

An optional scope may be provided to a commit, following the type/scope delimiter. A scope is a phrase providing additional contextual information:
- `feat(auth): add login feature`
- `fix(api): resolve token expiration bug`
- `refactor(widgets): simplify widget tree`

### Description

The description is a short summary of the code changes:
- Use imperative, present tense: "change" not "changed" nor "changes"
- Don't capitalize the first letter
- No period (.) at the end

### Body

The body is optional and should include the motivation for the change and contrast this with previous behavior:
- Use imperative, present tense
- Include the "why" and "what", not the "how"

### Footer

The footer is optional and should contain information about breaking changes and metadata:
- Breaking changes must start with "BREAKING CHANGE: " followed by a description

## Process

1. Run `git status` to see all changes
2. Run `git diff --cached` to see staged changes (if any)
3. Run `git diff` to see unstaged changes (if any)
4. Run `git log -10 --oneline` to understand the project's commit style
5. Analyze all changes and craft an appropriate conventional commit message
6. Stage all changes using `git add .`
7. Create the commit with a properly formatted conventional commit message

## Commit Message Template

Use this format when creating commits:

```
<type>[optional scope]: <short description>

<detailed body explaining what changed and why>

<optional footer with BREAKING CHANGE or references>
```

## Examples

Good commit messages:
- `feat(auth): add user registration flow`
- `fix(api): handle network timeout errors`
- `docs: update README with new installation steps`
- `refactor(components): extract common button styles`
- `feat(payment)!: remove deprecated payment method`

Breaking changes:
```
feat(dropwizard)!: drop Java 7 support

Java 7 support is being dropped to take advantage of Java 8's improved performance and language features.

BREAKING CHANGE: The minimum Java version is now 8.
```

## Important Notes

- Always commit ALL changes (staged and unstaged)
- The commit message MUST follow the conventional commits format
- Reference the specific files or areas changed when appropriate
- Keep descriptions concise but informative
- Add `!` after the type/scope to indicate breaking changes

After creating the commit, report back to the user with:
- The commit hash
- The commit message used
- Summary of files changed
