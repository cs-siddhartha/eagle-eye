# Go API

This directory contains the Go control-plane and telemetry-query API. Raw OTLP
ingestion will be handled by the OpenTelemetry Collector in a later increment.

## Run

```sh
cd backend
go run ./cmd/api
```

The API listens on `:8080` by default.

```sh
curl -i http://localhost:8080/readyz
```

## Test

```sh
cd backend
go test ./...
```

## Configuration

| Variable | Default |
| --- | --- |
| `API_ADDRESS` | `:8080` |
| `API_READ_HEADER_TIMEOUT` | `5s` |
| `API_READ_TIMEOUT` | `15s` |
| `API_WRITE_TIMEOUT` | `15s` |
| `API_IDLE_TIMEOUT` | `60s` |
| `API_SHUTDOWN_TIMEOUT` | `10s` |

Durations use Go duration syntax, such as `500ms`, `5s`, or `1m`.
