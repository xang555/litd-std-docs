# Backend Team Troubleshooting Guide

## Common Issues and Solutions

## Build Issues

### "Cannot find package" Error

**Symptoms:**
```
cannot find package "github.com/pkg/errors"
```

**Solutions:**
1. Run `go mod tidy` to clean up dependencies
2. Check go.mod for correct module path
3. Clear Go build cache: `go clean -cache`
4. Verify GOPATH and GORATE are set correctly

### "Import cycle" Error

**Symptoms:**
```
import cycle not allowed
```

**Solutions:**
1. Check circular dependencies in imports
2. Create a new package to break the cycle
3. Move shared code to a common package
4. Use dependency injection to reduce coupling

## Runtime Issues

### "Context Deadline Exceeded"

**Symptoms:**
Operations timing out prematurely

**Solutions:**
1. Check context timeout values
2. Identify slow operations (database, external API)
3. Add metrics to track operation duration
4. Consider increasing timeout for long operations

```go
// Adjust timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

### "Connection Refused" (Database)

**Symptoms:**
```
dial tcp 127.0.0.1:5432: connect: connection refused
```

**Solutions:**
1. Verify database is running
2. Check connection string (host, port)
3. Ensure firewall isn't blocking
4. Test with psql or other client tool

### "Too Many Open Files"

**Symptoms:**
```
too many open files
```

**Solutions:**
1. Check for file descriptor leaks
2. Ensure files are closed (use defer)
3. Increase system limits: `ulimit -n 4096`
4. Check for connection pool leaks

## Performance Issues

### High Memory Usage

**Diagnosis:**
1. Use pprof to profile memory: `go tool pprof`
2. Check for memory leaks
3. Look for large allocations in hot paths

**Solutions:**
1. Use sync.Pool for object reuse
2. Stream large responses instead of loading in memory
3. Limit concurrent operations
4. Profile with: `go tool pprof http://localhost:6060/debug/pprof/heap`

### Slow Database Queries

**Diagnosis:**
1. Enable slow query log
2. Use EXPLAIN ANALYZE
3. Check for missing indexes

**Solutions:**
1. Add appropriate indexes
2. Rewrite queries (avoid SELECT *)
3. Use connection pooling
4. Consider caching for frequent queries

### High CPU Usage

**Diagnosis:**
1. Use CPU profiler: `go tool pprof`
2. Check for tight loops
3. Look for inefficient algorithms

**Solutions:**
1. Profile with: `go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30`
2. Optimize hot paths
3. Use more efficient data structures
4. Implement caching

## Concurrency Issues

### Race Conditions

**Symptoms:**
Inconsistent behavior, data corruption

**Diagnosis:**
```bash
go test -race ./...
```

**Solutions:**
1. Use mutexes to protect shared state
2. Use channels for communication
3. Avoid shared mutable state when possible
4. Run tests with race detector

```go
// Good: Protected access
var mu sync.Mutex
var count int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}
```

### Deadlocks

**Symptoms:**
Application hangs, goroutines stuck

**Solutions:**
1. Acquire locks in consistent order
2. Use context with timeouts
3. Avoid holding locks while calling external functions
4. Use deadlock detector in tests

```go
// Good: Consistent lock order
func transfer(from, to *Account, amount int) error {
    // Always lock smaller ID first
    if from.ID < to.ID {
        from.Lock()
        to.Lock()
    } else {
        to.Lock()
        from.Lock()
    }
    defer from.Unlock()
    defer to.Unlock()
    // ...
}
```

## API Issues

### 404 Not Found on All Routes

**Symptoms:**
All API endpoints return 404

**Solutions:**
1. Check router is properly registered
2. Verify handler functions exist
3. Check middleware isn't blocking requests
4. Ensure server is listening on correct port

### CORS Errors

**Symptoms:**
Browser blocks API requests

**Solutions:**
1. Add CORS middleware
2. Configure allowed origins properly
3. Handle preflight requests (OPTIONS)

```go
// Good: CORS middleware
func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        if r.Method == "OPTIONS" {
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

## Deployment Issues

### Container Crashes

**Symptoms:**
Pod/container restarting repeatedly

**Solutions:**
1. Check logs: `kubectl logs <pod>`
2. Verify environment variables
3. Check resource limits
4. Ensure health check endpoints work

### Health Check Failing

**Symptoms:**
Service marked unhealthy

**Solutions:**
1. Verify health check endpoint
2. Check database connectivity
3. Ensure all dependencies are available
4. Add detailed logging to health checks

## Debugging Tools

### Built-in Tools

```bash
# pprof endpoints
curl http://localhost:6060/debug/pprof/heap
curl http://localhost:6060/debug/pprof/profile

# Environment variables
go env -w GODEBUG=gctrace=1

# Race detector
go test -race
```

### External Tools

- `dlv` - Go debugger
- `pprof` - Performance profiling
- `trace` - Execution tracing

## Getting Help

1. **Check logs**: Application logs, system logs
2. **Enable debug mode**: Set log level to debug
3. **Search**: Go forums, Stack Overflow
4. **Ask team**: Post in backend team channel
5. **Create issue**: Document the problem
