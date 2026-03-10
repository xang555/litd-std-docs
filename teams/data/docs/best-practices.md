# Data Team Best Practices

## Overview
Best practices for data engineering and analytics at our company.

## Data Pipeline Design

### Principles

1. **Idempotency**: Rerunning produces same results
2. **Atomicity**: Operations complete fully or not at all
3. **Isolation**: Separate dev, staging, and production
4. **Observability**: Monitor and log everything

## Spark Best Practices

### Performance Optimization

```python
# Cache frequently used DataFrames
df_cached = df.cache()

# Use broadcast joins for small tables
from pyspark.sql.functions import broadcast
result = large_df.join(broadcast(small_df), "id")

# Coalesce reduces partitions before writing
df.coalesce(1).write.parquet(path)

# Use appropriate partition sizes
df.repartition(200, "date_column")
```

### Resource Management

```python
# Configure executor resources
spark.conf.set("spark.executor.memory", "8g")
spark.conf.set("spark.executor.cores", "4")
spark.conf.set("spark.dynamicAllocation.enabled", "true")

# Enable shuffle service
spark.conf.set("spark.shuffle.service.enabled", "true")
```

## Airflow Best Practices

### DAG Design

```python
# Use meaningful task IDs
extract_task = PythonOperator(
    task_id='extract_users_from_api',
    python_callable=extract_users
)

# Add documentation
dag.doc_md = """
This DAG processes user data daily.
"""
```

### XCom Usage

```python
# Pass small data between tasks
def extract(**context):
    data = fetch_data()
    context['ti'].xcom_push(key='data', value=data)

def transform(**context):
    data = context['ti'].xcom_pull(key='data', task_ids='extract')
```

## dbt Best Practices

### Model Organization

```
models/
├── staging/        # Raw data cleaned
├── intermediate/   # Business logic
└── marts/          # Final aggregates
```

### Testing

```sql
-- models/staging/users.sql
{{ config(
    materialized='table',
    tags=['daily', 'pii']
) }}

SELECT
    user_id,
    email,
    created_at
FROM {{ source('raw', 'users') }}

-- tests/users_tests.sql
SELECT * FROM {{ ref('users') }}
WHERE user_id IS NULL
```

## Data Quality

### Validation Framework

```python
def validate_data(df, validations):
    """
    Run validations on DataFrame

    Args:
        df: Spark DataFrame
        validations: Dict of validation rules

    Returns:
        Dict with validation results
    """
    results = {}
    for name, rule in validations.items():
        result = rule(df)
        results[name] = result
        if not result.passed:
            logger.error(f"Validation failed: {name}")
    return results
```

### Data Quality Checks

```python
# Essential checks
validations = {
    'no_null_ids': lambda df: df.filter(col('id').isNull()).count() == 0,
    'positive_amounts': lambda df: df.filter(col('amount') < 0).count() == 0,
    'unique_ids': lambda df: df.count() == df.select('id').distinct().count(),
}
```

## Data Modeling

### Dimensional Modeling

```sql
-- Star schema example
-- Fact table
CREATE TABLE fact_orders (
    order_key BIGINT,
    customer_key INT,
    product_key INT,
    order_date DATE,
    amount DECIMAL(10,2)
);

-- Dimension tables
CREATE TABLE dim_customers (
    customer_key INT,
    customer_id INT,
    name VARCHAR,
    segment VARCHAR
);
```

### Slowly Changing Dimensions

```sql
-- Type 2 SCD example
CREATE TABLE dim_customers_scd2 (
    customer_key INT,
    customer_id INT,
    name VARCHAR,
    valid_from DATE,
    valid_to DATE,
    is_current BOOLEAN
);
```

## Security

### PII Handling

```python
# Mask PII in logs
import re

email = "user@example.com"
masked = re.sub(r'(?<=.{2}).(?=.*@)', '*', email)  # us**@example.com
logger.info(f"Processing user: {masked}")
```

### Access Control

```python
# Use role-based access
def get_s3_path(environment):
    paths = {
        'dev': 's3://dev-bucket/',
        'prod': 's3://prod-bucket/'
    }
    role = get_user_role()
    if role not in ['data-engineer', 'data-scientist']:
        raise PermissionError("Insufficient privileges")
    return paths[environment]
```

## Performance

### Query Optimization

```sql
-- Use partition pruning
WHERE date >= '2024-01-01' AND date < '2024-02-01'

-- Use appropriate join types
SELECT * FROM orders
INNER JOIN customers ON orders.customer_id = customers.id

-- Use CTEs for readability
WITH customer_orders AS (
    SELECT customer_id, COUNT(*) as order_count
    FROM orders
    GROUP BY customer_id
)
SELECT * FROM customer_orders WHERE order_count > 10
```

## Monitoring

### Key Metrics

1. **Pipeline Metrics**
   - Run duration
   - Data volume processed
   - Records per second
   - Error rate

2. **Data Metrics**
   - Row counts
   - Null percentages
   - Duplicate counts
   - Schema drift

### Alerting

```python
# Alert on data quality issues
if row_count < expected_count * 0.9:
    send_alert(
        severity='WARNING',
        message=f"Low row count: {row_count} vs expected {expected_count}"
    )
```

## Documentation

### Required Documentation

- Pipeline purpose and logic
- Data source descriptions
- Data quality rules
- SLA requirements
- Runbook for common issues

### Data Dictionary

```markdown
# users_daily_metrics

| Column | Type | Description | Source |
|--------|------|-------------|--------|
| user_id | string | Unique user identifier | users.id |
| event_date | date | Date of activity | events.event_date |
| page_views | int | Number of pages viewed | COUNT(events) |
```

## Resources

- [Spark Programming Guide](https://spark.apache.org/docs/latest/)
- [Airflow Documentation](https://airflow.apache.org/docs/)
- [dbt Documentation](https://docs.getdbt.com/)
- [Data Engineering Book](https://www.dataengineeringbook.com/)
