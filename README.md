# HC Server

HC Server is a small, production-oriented HTTP server baseline written in Go. It
uses Go's highly concurrent `net/http` stack and is designed to be extended with
application routes without changing the operational foundations.

## Features

- Concurrent HTTP request handling through `net/http`
- `GET /healthz` liveness endpoint
- `GET /readyz` readiness endpoint
- Request IDs, access logs, panic recovery, and JSON error responses
- Configurable bind address, graceful-shutdown timeout, and log level
- Graceful shutdown on `SIGINT` and `SIGTERM`
- No external runtime dependencies

## Requirements

- Go 1.21 or newer

## Run

```bash
go run ./cmd/server
```

The server listens on `:8080` by default.

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
```

## Configuration

| Variable | Default | Description |
| --- | --- | --- |
| `HC_SERVER_ADDR` | `:8080` | TCP address to bind |
| `HC_SERVER_SHUTDOWN_TIMEOUT` | `10s` | Maximum graceful-shutdown duration |
| `HC_SERVER_LOG_LEVEL` | `INFO` | `DEBUG`, `INFO`, `WARN`, or `ERROR` |

Example:

```bash
HC_SERVER_ADDR=127.0.0.1:9000 HC_SERVER_LOG_LEVEL=DEBUG go run ./cmd/server
```

## Test and build

```bash
go test ./...
go build ./cmd/server
```

## Project layout

```text
cmd/server/          Application entry point
internal/server/     Configuration, routes, middleware, and server lifecycle
```

## Adding application routes

Application routes should be registered in `internal/server.NewHandler`. Keep
health and readiness endpoints available so load balancers and orchestrators can
distinguish process liveness from service readiness.

## License

This project is provided as-is for development and extension.
