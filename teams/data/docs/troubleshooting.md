# Data Team Troubleshooting Guide

## Common Issues and Solutions

## Spark Issues

### "OutOfMemoryError" in Spark

**Symptoms:**
Job fails with OOM error

**Solutions:**
1. Increase executor memory: `--executor-memory`
2. Increase shuffle partitions: `spark.sql.shuffle.partitions`
3. Use `repartition()` or `coalesce()` appropriately
4. Cache selectively
5. Use broadcast joins for small tables

```python
spark.conf.set("spark.sql.shuffle.partitions", "400")
df.cache()
```

### "Task Failed" Due to Data Skew

**Symptoms:**
Some tasks take much longer than others

**Solutions:**
1. Add salting for skewed keys
2. Increase partitions on skewed column
3. Use broadcast join if appropriate
4. Repartition on different column

```python
# Add salt for skewed keys
from pyspark.sql.functions import rand, floor, concat, col

df_skewed = df.withColumn("salted_key",
    concat(col("skewed_key"), lit("_"), floor(rand() * 10).cast("string"))
)
```

### Stage Taking Too Long

**Symptoms:**
Single stage stalls

**Diagnosis:**
```python
spark.sparkContext.setLogLevel("INFO")
# Look for skewed stages in UI
```

**Solutions:**
1. Check shuffle spill metrics
2. Increase shuffle memory fraction
3. Optimize join order
4. Use DataFrame API instead of RDD

## Airflow Issues

### DAG Not Showing Up

**Symptoms:**
DAG not visible in UI

**Solutions:**
1. Check file is in DAG folder
2. Verify no syntax errors: `python -m py_compile dag.py`
3. Check DAG owner has permissions
4. Restart scheduler: `airflow scheduler`

### Tasks Stuck in "Queued" State

**Symptoms:**
Tasks not executing

**Solutions:**
1. Check worker status: `airflow workers`
2. Verify queue is not empty
3. Check for resource limits
4. Review parallelism settings

### DAG Runs Too Slow

**Symptoms:**
Tasks take longer than expected

**Solutions:**
1. Increase worker parallelism
2. Use appropriate pool sizes
3. Optimize task dependencies
4. Check for blocking operations

## Data Quality Issues

### Row Count Mismatch

**Symptoms:**
Output has fewer/more rows than expected

**Diagnosis:**
```python
print(f"Input: {input_count}, Output: {output_count}")
assert input_count == output_count, "Row count mismatch!"
```

**Solutions:**
1. Check for duplicate keys
2. Verify join conditions
3. Check filter logic
4. Review data source changes

### Null Values Unexpectedly

**Symptoms:**
Unexpected null values in output

**Solutions:**
1. Check for schema changes
2. Verify join types (inner vs left)
3. Review null handling logic
4. Check source data quality

### Data Drift

**Symptoms:**
Schema or data format changed

**Solutions:**
1. Implement schema validation
2. Monitor source metadata
3. Add data type checks
4. Version schemas

## Performance Issues

### Slow Pipeline Execution

**Diagnosis:**
```bash
# Check Spark UI for bottlenecks
# Look for:
# - Skewed stages
# - Spill to disk
# - GC time
```

**Solutions:**
1. Profile with Spark UI
2. Optimize shuffle operations
3. Increase resources if needed
4. Use appropriate file formats (Parquet)

### High Memory Usage

**Diagnosis:**
```python
# Check executor memory in Spark UI
# Look for:
# - High storage memory
# - Cached data
```

**Solutions:**
1. Release cached data: `df.unpersist()`
2. Reduce cache usage
3. Use smaller executor memory
4. Process data in batches

## File System Issues

### S3 Connection Timeout

**Symptoms:**
Timeout when reading/writing to S3

**Solutions:**
1. Check network connectivity
2. Verify credentials
3. Increase timeout: `fs.s3.connection.timeout`
4. Use S3 endpoint acceleration

### File Not Found

**Symptoms:**
```
java.io.FileNotFoundException: No such file or directory
```

**Solutions:**
1. Verify path is correct
2. Check for trailing slashes
3. Verify file exists: `aws s3 ls <path>`
4. Check permissions

## dbt Issues

### Model Not Found

**Symptoms:**
dbt can't find model

**Solutions:**
1. Check model path
2. Verify `dbt_project.yml` configuration
3. Run `dbt clean`
4. Check for naming conflicts

### Compilation Error

**Symptoms:**
SQL syntax error in model

**Solutions:**
1. Check SQL syntax
2. Verify macro usage
3. Check Jinja template syntax
4. Review error message for line number

### Test Failure

**Symptoms:**
Data quality tests fail

**Solutions:**
1. Review test logic
2. Check source data
3. Update test thresholds
4. Document expected failures

## Debugging Tools

```python
# Spark debugging
df.explain(True)  # Show physical plan
df.printSchema()  # Show schema
df.show(10, False)  # Show data without truncation
df.describe().show()  # Show statistics

# Airflow debugging
airflow dags list
airflow tasks test <dag_id> <task_id> <execution_date>
airflow dags show <dag_id>

# General debugging
import logging
logging.basicConfig(level=logging.DEBUG)
```

## Getting Help

1. **Check logs**: Spark UI, Airflow logs, dbt logs
2. **Verify assumptions**: Data counts, schemas, permissions
3. **Isolate the issue**: Run minimal reproducible example
4. **Search**: Stack Overflow, GitHub issues
5. **Ask team**: Post in data team channel
6. **Document**: Create issue for persistent problems

## Incident Response

1. **Assess impact**: Which data/products affected?
2. **Communicate**: Notify stakeholders
3. **Contain**: Stop failing pipelines
4. **Investigate**: Root cause analysis
5. **Fix**: Implement solution
6. **Verify**: Test fix thoroughly
7. **Document**: Post-mortem and improvements
