# ETL Pipeline Template

## Description

Production-ready ETL pipeline template using Python with data validation, error handling, and monitoring.

## Pipeline Structure

```
etl_pipeline/
├── config/
│   └── config.yaml         # Pipeline configuration
├── extract/
│   ├── __init__.py
│   └── extractors.py       # Data extraction modules
├── transform/
│   ├── __init__.py
│   ├── transformers.py     # Data transformation logic
│   └── validators.py       # Data validation
├── load/
│   ├── __init__.py
│   └── loaders.py          # Data loading modules
├── monitoring/
│   ├── __init__.py
│   └── metrics.py          # Monitoring and logging
└── main.py                 # Pipeline orchestration
```

## Main Pipeline Template

```python
#!/usr/bin/env python3
"""
ETL Pipeline Main Orchestration

This module orchestrates the Extract, Transform, Load process
with proper error handling and monitoring.
"""

import logging
import sys
from datetime import datetime
from typing import Any, Dict

import yaml
from pathlib import Path

from extract import DatabaseExtractor, APIExtractor
from transform import DataTransformer, DataValidator
from load import DatabaseLoader, S3Loader
from monitoring import MetricsCollector, setup_logging

logger = logging.getLogger(__name__)


class ETLPipeline:
    """Orchestrates the ETL process."""

    def __init__(self, config_path: str):
        """Initialize the ETL pipeline with configuration."""
        self.config = self._load_config(config_path)
        self.metrics = MetricsCollector()
        self.validator = DataValidator(self.config.get('validation', {}))
        self.transformer = DataTransformer(self.config.get('transform', {}))

    def _load_config(self, config_path: str) -> Dict[str, Any]:
        """Load pipeline configuration from YAML file."""
        with open(config_path, 'r') as f:
            return yaml.safe_load(f)

    def run(self) -> bool:
        """Execute the complete ETL pipeline."""
        start_time = datetime.now()
        pipeline_name = self.config.get('name', 'etl-pipeline')

        logger.info(f"Starting pipeline: {pipeline_name}")

        try:
            # Extract phase
            logger.info("Starting extraction phase")
            raw_data = self._extract()

            # Transform phase
            logger.info("Starting transformation phase")
            transformed_data = self._transform(raw_data)

            # Load phase
            logger.info("Starting load phase")
            self._load(transformed_data)

            # Record metrics
            duration = (datetime.now() - start_time).total_seconds()
            self.metrics.record_pipeline_success(duration, len(transformed_data))

            logger.info(f"Pipeline completed successfully in {duration:.2f}s")
            return True

        except Exception as e:
            duration = (datetime.now() - start_time).total_seconds()
            self.metrics.record_pipeline_failure(str(e), duration)
            logger.error(f"Pipeline failed: {e}", exc_info=True)
            return False

    def _extract(self) -> Any:
        """Extract data from source systems."""
        source_type = self.config['source']['type']

        if source_type == 'database':
            extractor = DatabaseExtractor(self.config['source'])
        elif source_type == 'api':
            extractor = APIExtractor(self.config['source'])
        else:
            raise ValueError(f"Unsupported source type: {source_type}")

        return extractor.extract()

    def _transform(self, raw_data: Any) -> Any:
        """Transform and validate the extracted data."""
        # Validate input data
        validation_errors = self.validator.validate(raw_data)
        if validation_errors:
            raise ValueError(f"Data validation failed: {validation_errors}")

        # Apply transformations
        transformed_data = self.transformer.transform(raw_data)

        # Validate output data
        output_errors = self.validator.validate_output(transformed_data)
        if output_errors:
            raise ValueError(f"Output validation failed: {output_errors}")

        return transformed_data

    def _load(self, data: Any) -> None:
        """Load data into destination systems."""
        destination_type = self.config['destination']['type']

        if destination_type == 'database':
            loader = DatabaseLoader(self.config['destination'])
        elif destination_type == 's3':
            loader = S3Loader(self.config['destination'])
        else:
            raise ValueError(f"Unsupported destination type: {destination_type}")

        loader.load(data)


def main():
    """Main entry point for the ETL pipeline."""
    # Setup logging
    setup_logging()

    # Load configuration
    config_path = sys.argv[1] if len(sys.argv) > 1 else 'config/config.yaml'

    # Run pipeline
    pipeline = ETLPipeline(config_path)
    success = pipeline.run()

    # Exit with appropriate code
    sys.exit(0 if success else 1)


if __name__ == '__main__':
    main()
```

## Database Extractor Template

```python
"""Data extraction modules."""

import logging
from typing import Any, Dict, List
import sqlalchemy
from sqlalchemy import create_engine, text

logger = logging.getLogger(__name__)


class DatabaseExtractor:
    """Extracts data from SQL databases."""

    def __init__(self, config: Dict[str, Any]):
        """Initialize database connection."""
        self.config = config
        self.engine = self._create_engine()

    def _create_engine(self) -> sqlalchemy.engine.Engine:
        """Create SQLAlchemy engine."""
        connection_string = self.config['connection_string']
        return create_engine(connection_string)

    def extract(self) -> List[Dict[str, Any]]:
        """Extract data using configured query."""
        query = self.config.get('query')
        batch_size = self.config.get('batch_size', 1000)

        logger.info(f"Starting data extraction with batch size: {batch_size}")

        try:
            with self.engine.connect() as conn:
                result = conn.execute(text(query))
                columns = result.keys()
                data = [dict(zip(columns, row)) for row in result]

            logger.info(f"Extracted {len(data)} records")
            return data

        except Exception as e:
            logger.error(f"Extraction failed: {e}")
            raise


class APIExtractor:
    """Extracts data from REST APIs."""

    def __init__(self, config: Dict[str, Any]):
        """Initialize API client."""
        self.config = config
        self.base_url = config['base_url']
        self.headers = config.get('headers', {})
        self.timeout = config.get('timeout', 30)

    def extract(self) -> List[Dict[str, Any]]:
        """Extract data from API endpoints."""
        import requests

        endpoints = self.config.get('endpoints', [])
        all_data = []

        for endpoint in endpoints:
            url = f"{self.base_url}/{endpoint}"
            logger.info(f"Fetching data from: {url}")

            try:
                response = requests.get(
                    url,
                    headers=self.headers,
                    timeout=self.timeout
                )
                response.raise_for_status()

                data = response.json()
                all_data.extend(data if isinstance(data, list) else [data])

            except Exception as e:
                logger.error(f"API request failed for {url}: {e}")
                raise

        logger.info(f"Extracted {len(all_data)} records from API")
        return all_data
```

## Data Transformer Template

```python
"""Data transformation modules."""

import logging
from typing import Any, Dict, List
import pandas as pd

logger = logging.getLogger(__name__)


class DataTransformer:
    """Transforms data according to business rules."""

    def __init__(self, config: Dict[str, Any]):
        """Initialize transformer with configuration."""
        self.config = config
        self.rules = config.get('rules', [])

    def transform(self, raw_data: List[Dict[str, Any]]) -> List[Dict[str, Any]]:
        """Apply transformation rules to data."""
        logger.info(f"Transforming {len(raw_data)} records")

        # Convert to DataFrame for easier manipulation
        df = pd.DataFrame(raw_data)

        # Apply each transformation rule
        for rule in self.rules:
            df = self._apply_rule(df, rule)

        # Convert back to list of dictionaries
        transformed_data = df.to_dict('records')

        logger.info(f"Transformation complete: {len(transformed_data)} records")
        return transformed_data

    def _apply_rule(self, df: pd.DataFrame, rule: Dict[str, Any]) -> pd.DataFrame:
        """Apply a single transformation rule."""
        rule_type = rule['type']

        if rule_type == 'rename':
            df = df.rename(columns=rule['mapping'])

        elif rule_type == 'filter':
            condition = rule['condition']
            df = df.query(condition)

        elif rule_type == 'derive':
            # Derive new columns from existing ones
            df[rule['column']] = df.eval(rule['expression'])

        elif rule_type == 'aggregate':
            # Group by and aggregate
            group_by = rule['group_by']
            aggregations = rule['aggregations']
            df = df.groupby(group_by).agg(aggregations).reset_index()

        elif rule_type == 'clean':
            # Remove nulls, duplicates, etc.
            if rule.get('remove_nulls'):
                df = df.dropna(subset=rule.get('columns'))
            if rule.get('remove_duplicates'):
                df = df.drop_duplicates()

        return df


class DataValidator:
    """Validates data against defined rules."""

    def __init__(self, config: Dict[str, Any]):
        """Initialize validator with configuration."""
        self.config = config
        self.schema = config.get('schema', {})

    def validate(self, data: List[Dict[str, Any]]) -> List[str]:
        """Validate data and return list of errors."""
        errors = []

        if not data:
            errors.append("No data to validate")
            return errors

        # Check required fields
        required_fields = self.schema.get('required_fields', [])
        for field in required_fields:
            if field not in data[0]:
                errors.append(f"Missing required field: {field}")

        # Validate data types
        field_types = self.schema.get('field_types', {})
        for record in data:
            for field, expected_type in field_types.items():
                if field in record and not isinstance(record[field], eval(expected_type)):
                    errors.append(f"Invalid type for {field}: expected {expected_type}")

        return errors

    def validate_output(self, data: List[Dict[str, Any]]) -> List[str]:
        """Validate output data before loading."""
        errors = []

        # Check for empty data
        if not data:
            errors.append("Output data is empty")

        # Check for null values in critical fields
        critical_fields = self.schema.get('critical_fields', [])
        for record in data:
            for field in critical_fields:
                if field in record and record[field] is None:
                    errors.append(f"Null value in critical field: {field}")

        return errors
```

## Database Loader Template

```python
"""Data loading modules."""

import logging
from typing import Any, Dict, List
import sqlalchemy
from sqlalchemy import create_engine, text
from sqlalchemy.types import TypeDecorator

logger = logging.getLogger(__name__)


class DatabaseLoader:
    """Loads data into SQL databases."""

    def __init__(self, config: Dict[str, Any]):
        """Initialize database connection."""
        self.config = config
        self.engine = self._create_engine()
        self.table = config['table']

    def _create_engine(self) -> sqlalchemy.engine.Engine:
        """Create SQLAlchemy engine."""
        connection_string = self.config['connection_string']
        return create_engine(connection_string)

    def load(self, data: List[Dict[str, Any]]) -> None:
        """Load data into database table."""
        import pandas as pd

        logger.info(f"Loading {len(data)} records into {self.table}")

        try:
            df = pd.DataFrame(data)

            # Configure load behavior
            if_exists = self.config.get('if_exists', 'append')
            chunk_size = self.config.get('chunk_size', 1000)

            # Load data
            df.to_sql(
                name=self.table,
                con=self.engine,
                if_exists=if_exists,
                chunksize=chunk_size,
                index=False
            )

            logger.info(f"Successfully loaded {len(data)} records")

        except Exception as e:
            logger.error(f"Load failed: {e}")
            raise

    def upsert(self, data: List[Dict[str, Any]]) -> None:
        """Upsert data into database table."""
        pass  # Implementation depends on database


class S3Loader:
    """Loads data into Amazon S3."""

    def __init__(self, config: Dict[str, Any]):
        """Initialize S3 client."""
        import boto3

        self.config = config
        self.bucket = config['bucket']
        self.prefix = config.get('prefix', '')
        self.s3_client = boto3.client('s3')

    def load(self, data: List[Dict[str, Any]) -> None:
        """Load data to S3 as JSON/Parquet."""
        import json
        import pandas as pd
        from datetime import datetime

        logger.info(f"Loading {len(data)} records to S3")

        # Generate filename with timestamp
        timestamp = datetime.now().strftime('%Y%m%d_%H%M%S')
        format_type = self.config.get('format', 'json')

        if format_type == 'json':
            filename = f"{self.prefix}/{timestamp}.json"
            content = json.dumps(data).encode('utf-8')
            content_type = 'application/json'

        elif format_type == 'parquet':
            filename = f"{self.prefix}/{timestamp}.parquet"
            df = pd.DataFrame(data)
            import io
            buffer = io.BytesIO()
            df.to_parquet(buffer, index=False)
            content = buffer.getvalue()
            content_type = 'application/octet-stream'

        # Upload to S3
        try:
            self.s3_client.put_object(
                Bucket=self.bucket,
                Key=filename,
                Body=content,
                ContentType=content_type
            )
            logger.info(f"Successfully uploaded to s3://{self.bucket}/{filename}")

        except Exception as e:
            logger.error(f"S3 upload failed: {e}")
            raise
```

## Configuration Template

```yaml
# config/config.yaml

name: "sales-data-pipeline"
description: "ETL pipeline for sales data processing"

# Source configuration
source:
  type: database
  connection_string: "postgresql://user:pass@localhost:5432/db"
  query: |
    SELECT id, customer_id, product_id, quantity, amount, created_at
    FROM sales
    WHERE created_at >= CURRENT_DATE - INTERVAL '1 day'
  batch_size: 5000

# Transformation configuration
transform:
  rules:
    - type: rename
      mapping:
        customer_id: client_id
        product_id: item_id

    - type: derive
      column: total_amount
      expression: quantity * amount

    - type: filter
      condition: "total_amount > 0"

    - type: clean
      remove_nulls: true
      columns: ["client_id", "item_id"]
      remove_duplicates: true

# Validation configuration
validation:
  schema:
    required_fields: ["id", "client_id", "item_id", "total_amount"]
    critical_fields: ["id", "client_id", "item_id"]
    field_types:
      id: "int"
      total_amount: "float"

# Destination configuration
destination:
  type: s3
  bucket: "my-data-bucket"
  prefix: "sales-data/processed"
  format: parquet
```

## Tags

`data`, `etl`, `pipeline`, `python`, `extraction`, `transformation`, `loading`
