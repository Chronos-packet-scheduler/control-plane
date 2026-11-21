# control-plane

Chronos control plane: source of truth for rules, telemetry aggregation.

```sh
go run ./cmd/control-plane          # serves /healthz on :8080
docker build -t control-plane .
```

Images are published to `ghcr.io/chronos-packet-scheduler/control-plane` on every push to `main` and on `v*` tags.