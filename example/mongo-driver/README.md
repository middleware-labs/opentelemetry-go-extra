# mongo-driver example for OpenTelemetry

[![PkgGoDev](https://pkg.go.dev/badge/go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo)](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo)

To run this example you need a MongoDB server. You can start one with Docker:

```shell
docker-compose up -d
```

You can run this example with different exporters by providing environment variables.

**Stdout** exporter (default):

```shell
go run .
```

**Jaeger** exporter:

```shell
OTEL_EXPORTER_JAEGER_ENDPOINT=http://localhost:14268/api/traces go run .
```

[Uptrace](https://github.com/middleware-labs/uptrace/) exporter:

```shell
UPTRACE_DSN="https://<token>@uptrace.dev/<project_id>" go run .
```

## Links

- [OpenTelemetry Tracing API](https://uptrace.dev/opentelemetry/go-tracing.html)
- [DataDog Competitors](https://uptrace.dev/blog/datadog-competitors.html)
