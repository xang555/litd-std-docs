# Pipeline Development Standards

## Scope
This standard applies to all data pipelines built with Spark, Airflow, and dbt.

## Rules

### 1. Naming Conventions

**Rule:** Use snake_case for all pipeline-related naming.

**Rationale:** Consistent naming improves discoverability.

**Example:**
```python
# Good
user_analytics_dag
transform_orders
calculate_metrics

# Bad
UserAnalyticsDAG
Transform-Orders
calcMetrics
```

### 2. DAG Organization

**Rule:** Each DAG should have a single, clear purpose.

**Rationale:** Monolithic DAGs are hard to maintain and debug.

**Example:**
```python
# Good
# DAG: process_orders.py - Processes daily orders
# DAG: sync_customers.py - Syncs customer data

# Bad
# DAG: data_pipeline.py - Does everything
```

### 3. Idempotency

**Rule:** All pipeline tasks must be idempotent.

**Rationale:** Enables safe re-runs and recovery.

**Example:**
```python
# Good - Uses overwrite mode
df.write.mode("overwrite").parquet(path)

# Good - Checks for existence
if not spark.read.parquet(path).count() > 0:
    df.write.parquet(path)

# Bad - Appends without checking
df.write.mode("append").parquet(path)
```

### 4. Configuration Externalization

**Rule:** Keep configuration separate from code.

**Rationale:** Enables different configurations per environment.

**Example:**
```python
# Good
import yaml

with open('config/prod.yaml') as f:
    config = yaml.safe_load(f)

input_path = config['input_path']

# Bad
input_path = "s3://production-bucket/data/"  # Hardcoded
```

### 5. Error Handling

**Rule:** Implement proper error handling and logging.

**Rationale:** Makes debugging easier and prevents silent failures.

**Example:**
```python
# Good
import logging

logger = logging.getLogger(__name__)

try:
    df = spark.read.parquet(input_path)
except Exception as e:
    logger.error(f"Failed to read {input_path}: {e}")
    raise

# Bad
df = spark.read.parquet(input_path)  # No error handling
```

### 6. Data Documentation

**Rule:** Document all data sources, transformations, and outputs.

**Rationale:** Essential for data governance and debugging.

**Example:**
```python
"""
User Analytics Pipeline

Description: Calculates daily user engagement metrics

Input:
  - users.parquet: User profile data
  - events.parquet: User interaction events

Transformations:
  - Aggregates events by user and date
  - Joins with user profile data
  - Calculates engagement score

Output:
  - user_metrics.parquet: Daily metrics per user

Schedule: Daily at 2 AM UTC
Owner: data-team
"""
```

### 7. Testing

**Rule:** All pipelines must have data quality tests.

**Rationale:** Ensures data integrity and early issue detection.

**Example:**
```python
from airflow.sensors.sql import SqlSensor

def check_data_quality(**context):
    df = spark.read.parquet(output_path)
    assert df.count() > 0, "Output is empty"
    assert df.filter(col("user_id").isNull()).count() == 0, "Null user_ids found"
    return True

data_quality_check = PythonOperator(
    task_id='data_quality_check',
    python_callable=check_data_quality
)
```

### 8. Partitioning

**Rule:** Partition data by date for time-series data.

**Rationale:** Improves query performance and enables efficient data management.

**Example:**
```python
# Good
df.write.partitionBy("date").parquet(path)

# Bad
df.write.parquet(path)  # No partitioning
```

### 9. Schema Validation

**Rule:** Validate schemas at read and write.

**Rationale:** Catches schema drift early.

**Example:**
```python
from pyspark.sql.types import StructType, StructField

EXPECTED_SCHEMA = StructType([
    StructField("user_id", StringType(), nullable=False),
    StructField("event_date", DateType(), nullable=False),
    StructField("metric", DoubleType(), nullable=True),
])

df = spark.read.schema(EXPECTED_SCHEMA).parquet(path)
```

### 10. Monitoring

**Rule:** Include logging and metrics for all pipelines.

**Rationale:** Enables operational visibility.

**Example:**
```python
from airflow.providers.apache.spark.operators.spark import SparkSubmitOperator

job = SparkSubmitOperator(
    task_id='run_job',
    application='jobs/transform.py',
    conf={
        'spark.metrics.conf': '*.sink.prometheus.class=org.apache.spark.metrics.sink.PrometheusSink'
    }
)
```

## Enforcement

### Linting
- `pylint` for Python code
- `sqlfluff` for SQL
- `dbt lint` for dbt models

### Review
All pull requests must:
- Pass linting checks
- Have data quality tests
- Be reviewed by data engineer
- Include updated documentation

### Pre-commit Hooks
- Python linting and formatting
- SQL linting
- Schema validation

## Exceptions

Exceptions require:
- Data lead approval
- Risk assessment
- Mitigation plan
- Documentation

## Related Standards
- Data Modeling Standards
- Data Quality Standards
- Security Standards

## Tags
`data`, `pipeline`, `spark`, `airflow`, `standards`, `etl`
