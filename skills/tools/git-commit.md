# Git Commit

## Description

Writing effective commit messages using the Conventional Commits specification for better version control and automated changelog generation.

## Level: Junior

### Learning Objectives

- Understand the Conventional Commits format
- Write basic commit messages
- Learn common commit types

### Required Skills

1. **Commit Message Structure**

   - Format: `type[(scope)]: description`
   - Use imperative mood ("add" not "added")
   - Keep descriptions under 72 characters
   - No period at the end

2. **Common Commit Types**

   - `feat`: New feature
   - `fix`: Bug fix
   - `docs`: Documentation changes
   - `test`: Adding or updating tests
   - `chore`: Maintenance tasks

3. **Basic Examples**

   - `feat: add user login`
   - `fix: resolve navigation bug`
   - `docs: update README`

### Resources

- [Conventional Commits](https://www.conventionalcommits.org/)
- Team commit guidelines

## Level: Mid-Level

### Learning Objectives

- Use scopes effectively
- Write detailed commit bodies
- Understand breaking changes

### Required Skills

1. **Scopes**

   - Provide contextual information
   - Common scopes: auth, api, ui, db, config
   - Format: `type(scope): description`

2. **Commit Body**

   - Explain "what" and "why"
   - Use imperative mood
   - Wrap at 72 characters

3. **Breaking Changes**

   - Add `!` after type/scope
   - Include `BREAKING CHANGE:` footer
   - Document migration steps

### Optional Skills

- Commit message linting
- Husky commit hooks
- Commitizen tools

## Level: Senior

### Learning Objectives

- Establish commit conventions
- Implement automated validation
- Lead team adoption

### Required Skills

1. **Standards**

   - Team-specific commit types
   - Scope definitions
   - Style guides

2. **Automation**

   - Commitlint configuration
   - Pre-commit hooks
   - CI/CD integration

3. **Team Practices**

   - Training materials
   - Code review standards
   - Changelog automation

## Tags

`git`, `conventional-commits`, `workflow`
