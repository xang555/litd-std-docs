# Git Operations Command

## Description
Common Git operations for version control workflows.

## Usage
```
/git <operation> [options]
```

## Operations

### Branch Management

**Create and switch to new branch:**
```bash
/git branch create <branch-name>
```
Creates a new branch and switches to it.

**List branches:**
```bash
/git branch list
```
Shows all local and remote branches.

**Delete branch:**
```bash
/git branch delete <branch-name>
```
Deletes a local branch.

**Merge branch:**
```bash
/git branch merge <branch-name>
```
Merges specified branch into current branch.

### Commit Operations

**Stage and commit changes:**
```bash
/git commit "message" [--all]
```
Stages all changes (with --all) and commits with message.

**Amend last commit:**
```bash
/git commit amend
```
Modifies the last commit (use with caution).

**View commit history:**
```bash
/git log [oneline|graph]
```
Shows commit history in specified format.

### Sync Operations

**Pull latest changes:**
```bash
/git pull [--rebase]
```
Pulls changes from remote. Uses rebase if specified.

**Push changes:**
```bash
/git push [--force]
```
Pushes local commits to remote. Force pushes with --force.

**Fetch remote changes:**
```bash
/git fetch
```
Fetches changes without merging.

### Stash Operations

**Stash changes:**
```bash
/git stash save "message"
```
Stashes current changes with a message.

**List stashes:**
```bash
/git stash list
```
Shows all saved stashes.

**Apply stash:**
```bash
/git stash apply [stash@{n}]
```
Applies specified stash (latest if not specified).

**Drop stash:**
```bash
/git stash drop [stash@{n}]
```
Removes specified stash.

## Options

| Option | Description |
|--------|-------------|
| --verbose | Show detailed output |
| --dry-run | Show what would be done without doing it |
| --force | Force operation (use with caution) |

## Examples

### Feature development workflow
```bash
/git branch create feature/new-ui
# Make changes
/git commit "Add new UI components"
/git push
# Create PR
```

### Update from main
```bash
/git checkout main
/git pull
/git checkout feature-branch
/git branch merge main
```

### Quick stash and switch
```bash
/git stash save "work in progress"
/git checkout main
# Do something
/git checkout feature-branch
/git stash apply
```

## Tags
`git`, `version-control`, `cli`, `workflow`, `commands`
