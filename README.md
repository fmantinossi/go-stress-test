# Go Load Testing Tool

A simple CLI tool for load testing web services written in Go.

## Features

- Concurrent request execution
- Configurable number of total requests
- Detailed test report including:
  - Total execution time
  - Number of successful requests
  - Status code distribution

## Usage

### Building and Running Locally

```bash
go build -o load-test
./load-test --url=http://example.com --requests=1000 --concurrency=10
```

### Using Docker

```bash
# Build the Docker image
docker build -t go-load-test .

# Run the container
docker run go-load-test --url=http://example.com --requests=1000 --concurrency=10
```

## Parameters

- `--url`: URL of the service to test (required)
- `--requests`: Total number of requests to make (required)
- `--concurrency`: Number of concurrent requests (default: 1)

## Example Output

```
=== Load Test Report ===
Total Time: 5.234s
Total Requests: 1000
Successful Requests (200): 980

Status Code Distribution:
Status 200: 980 requests
Status 404: 15 requests
Status 500: 5 requests
```