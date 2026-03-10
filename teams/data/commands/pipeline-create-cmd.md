# Pipeline Create Command

## Description
Command template for scaffolding new data pipelines with Airflow and Spark.

## Usage
```
/pipeline-create <PipelineName> [options]
```

## Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| PipelineName | string | true | - | Name of the pipeline (snake_case) |
| --type | string | false | batch | Pipeline type: batch, streaming |
| --engine | string | false | spark | Execution engine: spark, sql, dbt |
| --schedule | string | false | daily | Schedule: hourly, daily, weekly |
| --with-tests | boolean | false | true | Include data quality tests |
| --with-monitoring | boolean | false | true | Include monitoring and alerting |

## Examples

### Batch pipeline
```bash
/pipeline-create user_analytics --type=batch --schedule=daily
```
Creates daily batch processing pipeline

### Streaming pipeline
```bash
/pipeline-create event_stream --type=streaming --engine=spark
```
Creates streaming data pipeline

### dbt pipeline
```bash
/pipeline-create mart_orders --engine=dbt --schedule=hourly
```
Creates dbt transformation pipeline

## Generated Structure

```
dags/
└── pipeline_name/
    ├── pipeline_name_dag.py        # Airflow DAG
    ├── jobs/
    │   ├── spark_job.py            # Spark job
    │   └── sql_queries.sql         # SQL queries
    ├── tests/
    │   └── test_pipeline.py        # Data quality tests
    ├── monitoring/
    │   └── alerts.yaml             # Alert configurations
    ├── config/
    │   ├── dev.yaml                # Development config
    │   └── prod.yaml               # Production config
    └── README.md                   # Pipeline documentation
```

## DAG Template

```python
from airflow import DAG
from airflow.providers.apache.spark.operators.spark import SparkSubmitOperator
from datetime import datetime, timedelta
import yaml

DEFAULT_ARGS = {
    'owner': 'data-team',
    'depends_on_past': False,
    'start_date': datetime(2024, 1, 1),
    'email_on_failure': True,
    'email_on_retry': False,
    'retries': 2,
    'retry_delay': timedelta(minutes=5),
}

with DAG(
    'pipeline_name',
    default_args=DEFAULT_ARGS,
    description='Pipeline description',
    schedule_interval='@daily',
    catchup=False,
    max_active_runs=1,
) as dag:

    spark_job = SparkSubmitOperator(
        task_id='run_spark_job',
        application='jobs/spark_job.py',
        conn_id='spark_default',
        conf={
            'spark.dynamicAllocation.enabled': 'true',
            'spark.shuffle.service.enabled': 'true',
        },
        application_args=['--date', '{{ ds }}'],
    )

    data_quality_checks = SparkSubmitOperator(
        task_id='data_quality_checks',
        application='jobs/tests/test_pipeline.py',
        conn_id='spark_default',
    )

    spark_job >> data_quality_checks
```

## Spark Job Template

```python
from pyspark.sql import SparkSession
from pyspark.sql.functions import col, count, sum as _sum
import argparse

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--date', required=True)
    args = parser.parse_args()

    spark = SparkSession.builder \
        .appName("pipeline_name") \
        .getOrCreate()

    try:
        # Read input data
        df = spark.read.parquet(f"s3://bucket/input/{args.date}")

        # Transform data
        result = df.groupBy("category") \
            .agg(
                count("*").alias("record_count"),
                _sum("amount").alias("total_amount")
            )

        # Write output
        result.write \
            .mode("overwrite") \
            .parquet(f"s3://bucket/output/{args.date}")

        spark.stop()
    except Exception as e:
        spark.stop()
        raise e

if __name__ == "__main__":
    main()
```

## Tags
`data`, `pipeline`, `airflow`, `spark`, `cli`, `etl`
