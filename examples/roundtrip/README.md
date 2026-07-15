# roundtrip

A full OddDotNet push→query round-trip: export one span over OTLP to a running
OddDotNet sink, then query it back over the span query API.

It demonstrates the canonical "wait for a specific signal" pattern —
`TakeFirst` + a name filter + `Wait` — so the query returns the instant the span
is ingested rather than blocking for the whole wait window (as `TakeAll` would).

This is a separate module so the OpenTelemetry SDK dependencies it needs for the
push side stay out of the OddDotGo library module.

## Run

Point it at a running sink (defaults to `127.0.0.1:4317`):

```bash
go run . -addr 127.0.0.1:4317
```

On success:

```
round-trip ok: span "roundtrip-<n>" ingested and queried back
```
