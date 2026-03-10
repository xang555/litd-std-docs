# API Versioning

## Overview

This document outlines the standards for versioning APIs across all teams.

## Versioning Strategy

We use URL-based versioning for public APIs:

```
/api/v1/resource
/api/v2/resource
```

## Version Support

- Support current major version (N)
- Support previous major version (N-1) for 6 months
- Provide migration guide for deprecation

## Breaking Changes

A breaking change includes:
- Removing or renaming endpoints
- Changing request/response structure
- Modifying required fields
- Changing authentication method

## Non-Breaking Changes

- Adding new endpoints
- Adding optional fields
- Adding new response headers
- Bug fixes

## Deprecation Process

1. Announce deprecation with timeline
2. Add deprecation headers to responses
3. Publish migration guide
4. Monitor usage
5. Remove after sunset period

## Response Headers

```
X-API-Version: 1.0.0
X-API-Deprecated: true
X-API-Sunset: Wed, 01 Jan 2025 00:00:00 GMT
```
