# Incident Response Workflow

## Description
Standard workflow for responding to and resolving production incidents.

## Severity Levels

| Severity | Description | Response Time |
|----------|-------------|---------------|
| SEV-1 | Critical system down | < 15 minutes |
| SEV-2 | Major functionality impacted | < 1 hour |
| SEV-3 | Minor functionality impacted | < 4 hours |
| SEV-4 | Cosmetic issues | Next business day |

## Steps

### 1. Detection & Triage
- **Owner**: On-call engineer
- **Actions**:
  - Acknowledge alert
  - Determine severity level
  - Create incident channel
  - Notify appropriate team members
- **Output**: Incident created, team assembled

### 2. Investigation
- **Owner**: Incident commander
- **Actions**:
  - Gather metrics and logs
  - Identify affected systems
  - Determine root cause hypothesis
  - Communicate status updates
- **Output**: Understanding of issue scope

### 3. Mitigation
- **Owner**: Incident commander
- **Actions**:
  - Implement temporary fix
  - Verify system recovery
  - Monitor for stability
  - Communicate resolution
- **Output**: System operational

### 4. Resolution
- **Owner**: On-call engineer
- **Actions**:
  - Implement permanent fix
  - Deploy to production
  - Verify complete resolution
  - Close incident
- **Output**: Issue fully resolved

### 5. Post-Incident
- **Owner**: Engineering manager
- **Actions**:
  - Schedule post-mortem
  - Document root cause
  - Create action items
  - Update runbooks
- **Output**: Post-mortem document, improvements

## Communication Templates

### Initial Announcement
```
🚨 Incident declared: [Brief description]
Severity: SEV-X
Impact: [Affected systems/users]
Channel: #incident-XXX
Next update: [Time]
```

### Status Updates
```
⏳ Update on incident #XXX
Status: [Investigating|Mitigating|Resolved]
Summary: [What's happening]
ETA: [If available]
```

### Resolution
```
✅ Incident #XXX resolved
Duration: [Time to resolve]
Root cause: [Summary]
Follow-up: [Post-mortem planned]
```

## Roles

- **Incident Commander**: Leads response, makes decisions
- **Communication Lead**: Manages stakeholder updates
- **Scribe**: Documents actions and decisions
- **Subject Matter Expert**: Provides technical expertise

## Outputs

| Output | Description | Format |
|--------|-------------|--------|
| Incident log | Timeline of actions | Document |
| Post-mortem | Analysis and improvements | Document |
| Runbook updates | Updated procedures | Documentation |

## Tags
`shared`, `workflow`, `incident-response`, `sre`, `operations`
