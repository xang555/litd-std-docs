# Data Team Standards

## Overview

This document outlines the coding standards, best practices, and conventions for the Data Engineering team.

## Technology Stack

- **ETL**: Python + Apache Airflow
- **Storage**: Snowflake / BigQuery
- **Processing**: Spark
- **Orchestration**: Dagster
- **Testing**: pytest

## Code Standards

### File Naming

- Pipelines: `kebab-case.py` (e.g., `user-pipeline.py`)
- Tests: `test_*.py`
- Notebooks: `*.ipynb` with descriptive names

### Data Quality

- Schema validation
- Data profiling
- Anomaly detection
- Lineage tracking
- Documentation

## Best Practices

- Idempotent pipelines
- Incremental processing when possible
- Proper error handling and retries
- Monitoring and alerting
- Data versioning (DVC)

## Related Documents

- [ETL Patterns](./etl-patterns.md)
