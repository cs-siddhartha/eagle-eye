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

## Checks

```sh
cd backend
go test ./...
go vet ./...
```

There are currently no automated test files. `go test ./...` is used as a
compile check until automated coverage is added in a later reviewed increment.

## Configuration

| Variable | Default |
| --- | --- |
| `API_ADDRESS` | `:8080` |
| `API_MAX_REQUEST_BYTES` | `1048576` |
| `API_READ_HEADER_TIMEOUT` | `5s` |
| `API_READ_TIMEOUT` | `15s` |
| `API_WRITE_TIMEOUT` | `15s` |
| `API_IDLE_TIMEOUT` | `60s` |
| `API_SHUTDOWN_TIMEOUT` | `10s` |

Durations use Go duration syntax, such as `500ms`, `5s`, or `1m`.

## Middleware behavior

Every request receives:

- An `X-Request-ID` response header
- Structured completion logging with method, path, status, response size, and duration
- Panic recovery with a JSON `500` response
- A maximum request-body size
- JSON content-type enforcement for `POST`, `PUT`, and `PATCH` requests with bodies
- JSON `404` and `405` responses

## Manual verification

Start the API:

```sh
cd backend
go run ./cmd/api
```

In another terminal, verify readiness:

```sh
curl -i -H 'X-Request-ID: manual-ready' \
  http://localhost:8080/readyz
```

Verify unsupported methods:

```sh
curl -i -X POST \
  -H 'Content-Type: application/json' \
  -d '{}' \
  http://localhost:8080/readyz
```

Verify unknown resources:

```sh
curl -i http://localhost:8080/missing
```

Verify JSON enforcement:

```sh
curl -i -X POST \
  -H 'Content-Type: text/plain' \
  -d 'not-json' \
  http://localhost:8080/api/v1/traces
```

Verify the request-size limit:

```sh
API_MAX_REQUEST_BYTES=16 go run ./cmd/api

curl -i -X POST \
  -H 'Content-Type: application/json' \
  -d '{"message":"this payload is too large"}' \
  http://localhost:8080/api/v1/traces
```
