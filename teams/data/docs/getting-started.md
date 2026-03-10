# Data Team - Getting Started

## Welcome to the Data Team

This guide helps you get started with data engineering at our company.

## Tech Stack

- **Language**: Python
- **Processing**: Apache Spark
- **Orchestration**: Airflow
- **Storage**: S3, Redshift
- **Streaming**: Kafka

## Setup

1. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```

2. Configure Airflow:
   ```bash
   airflow initdb
   ```

3. Start local Spark:
   ```bash
   spark-shell
   ```

## Project Structure

```
dags/
├── examples/       # Example DAGs
├── ingestion/      # Data ingestion
├── processing/     # Data processing
└── analytics/      # Analytics jobs
jobs/
├── batch/          # Batch jobs
└── streaming/      # Streaming jobs
libs/
├── extract/        # Data extraction
├── transform/      # Data transformation
└── load/           # Data loading
```

## Code Style

Follow the coding standards defined in `teams/data/standards/`.

## Resources

- Team-specific agents: `teams/data/agents/`
- Team workflows: `teams/data/workflows/`
- Shared resources: `shared/`
