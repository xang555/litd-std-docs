# DevOps Troubleshooting Guide

## Common Issues and Solutions

## Terraform Issues

### "State lock" Error

**Symptoms:**
```
Error: Error acquiring the state lock
```

**Solutions:**
1. Check if someone else is applying
2. Force unlock if safe: `terraform force-unlock <LOCK_ID>`
3. Verify no concurrent apply operations

### "Resource already managed" Error

**Symptoms:**
Terraform thinks it doesn't own a resource it created

**Solutions:**
1. Import the resource: `terraform import <address> <id>`
2. Or remove from state and recreate: `terraform state rm <address>`
3. Verify state file is correct

### "Backend configuration changed" Error

**Symptoms:**
Can't apply after changing backend configuration

**Solutions:**
1. Copy state to new backend: `terraform state push`
2. Or run: `terraform init -migrate-state`
3. Verify migration completed successfully

## Kubernetes Issues

### "CrashLoopBackOff" Status

**Symptoms:**
Pod keeps restarting

**Diagnosis:**
```bash
kubectl describe pod <pod-name>
kubectl logs <pod-name> --previous
```

**Solutions:**
1. Check application logs for errors
2. Verify environment variables and ConfigMaps
3. Check resource limits (OOMKilled)
4. Verify liveness/readiness probe configuration
5. Check image pull errors

### "ImagePullBackOff" Error

**Symptoms:**
Pod can't pull container image

**Solutions:**
1. Verify image name and tag
2. Check image pull secrets
3. Verify registry credentials
4. Check network connectivity to registry

### Pod Stuck in "Pending" State

**Symptoms:**
Pod never starts

**Diagnosis:**
```bash
kubectl describe pod <pod-name>
kubectl get events --sort-by=.metadata.creationTimestamp
```

**Solutions:**
1. Insufficient resources: Scale nodes or reduce requests
2. Node selector: Verify labels match
3. Taints and tolerations: Check node restrictions
4. PVC pending: Verify storage class availability

## Deployment Issues

### Rollout Stuck

**Symptoms:**
Deployment not progressing

**Solutions:**
1. Check deployment status: `kubectl rollout status deployment/<name>`
2. View deployment events: `kubectl describe deployment <name>`
3. Check replica set status
4. Rollback if needed: `kubectl rollout undo deployment/<name>`

### High Error Rate After Deployment

**Symptoms:**
5xx errors increased after deployment

**Solutions:**
1. Immediate rollback: `kubectl rollout undo deployment/<name>`
2. Investigate logs: `kubectl logs -f deployment/<name>`
3. Check metrics in dashboard
4. Review recent changes

## CI/CD Issues

### Pipeline Failing Randomly

**Symptoms:**
Flaky tests or intermittent failures

**Solutions:**
1. Add retries for external dependencies
2. Use fixed versions instead of `latest`
3. Increase timeout values
4. Check for resource exhaustion
5. Add proper cleanup between steps

### Docker Build Fails

**Symptoms:**
Build fails with no clear error

**Solutions:**
1. Check Dockerfile syntax
2. Verify base image exists
3. Check for cached layer issues: `--no-cache`
4. Verify build context
5. Check disk space

## Performance Issues

### High CPU Usage

**Diagnosis:**
```bash
kubectl top nodes
kubectl top pods
```

**Solutions:**
1. Identify consuming pods
2. Check for inefficient queries
3. Profile application code
4. Scale horizontally if needed

### High Memory Usage

**Diagnosis:**
```bash
kubectl top pods --sort-by=memory
```

**Solutions:**
1. Check for memory leaks
2. Review application logs
3. Adjust memory limits
4. Profile memory usage

### Slow API Response Times

**Diagnosis:**
1. Check application metrics
2. Review database query performance
3. Check network latency
4. Profile the application

**Solutions:**
1. Add caching
2. Optimize database queries
3. Add indexes
4. Scale horizontally

## Security Issues

### Unauthorized Access Attempt

**Symptoms:**
Failed authentication in logs

**Solutions:**
1. Review audit logs
2. Check compromised credentials
3. Rotate affected secrets
4. Review IAM policies

### Security Scan Failures

**Symptoms:**
Vulnerabilities detected in container images

**Solutions:**
1. Update base images
2. Update dependencies
3. Apply security patches
4. Review and accept false positives

## Network Issues

### Pod Can't Reach Service

**Symptoms:**
Connection refused or timeout

**Diagnosis:**
```bash
kubectl exec -it <pod> -- nslookup <service>
kubectl exec -it <pod> -- curl http://<service>:<port>
```

**Solutions:**
1. Verify service exists
2. Check service selector
3. Verify port configuration
4. Check network policies
5. Check DNS configuration

### External Access Not Working

**Symptoms:**
Can't access application from outside

**Solutions:**
1. Verify Ingress configuration
2. Check LoadBalancer service
3. Verify security groups/firewall rules
4. Check DNS records

## Monitoring Issues

### Alerts Not Firing

**Symptoms:**
Problems occur but no alerts

**Solutions:**
1. Verify alert rules exist
2. Check alert evaluation interval
3. Verify notification channels
4. Test alert rules

### Metrics Missing

**Symptoms:**
Expected metrics not appearing

**Solutions:**
1. Check Prometheus targets
2. Verify ServiceMonitor configuration
3. Check application metrics endpoint
4. Verify scraping configuration

## Debugging Tools

```bash
# Kubernetes debugging
kubectl debug -it <pod> --image=nicolaka/netshoot
kubectl logs -f <pod> --all-containers=true
kubectl get events --watch

# Terraform debugging
TF_LOG=DEBUG terraform apply
terraform plan -out=tfplan
terraform show tfplan

# General debugging
kubectl exec -it <pod> -- /bin/sh
docker logs -f <container>
```

## Getting Help

1. **Check logs**: Application, system, and security logs
2. **Metrics**: Check monitoring dashboards
3. **Documentation**: Runbooks and knowledge base
4. **Team**: Alert appropriate team members
5. **Escalate**: Follow escalation procedures if needed

## Incident Response

1. **Detect**: Identify the issue
2. **Respond**: Initial mitigation
3. **Resolve**: Fix the root cause
4. **Learn**: Post-mortem and improvements
