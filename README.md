# auditFlow

## Week 1 — Foundations
[TICKET-1] Project scaffold — Initialise Go module, define Event struct (id, source, action, actor, timestamp, metadata), set up Makefile with build/test/lint targets.

[TICKET-2] HTTP ingest API — Build a POST /events endpoint using net/http (or chi) that validates and accepts an Event payload, returning 202 Accepted. No persistence yet — just log to stdout.

## Week 2 — Kafka
[TICKET-3] Kafka producer — Integrate segmentio/kafka-go (or confluentinc/confluent-kafka-go). On POST /events, publish the event to a audit-events topic. Run Kafka locally via docker-compose.

[TICKET-4] Kafka consumer — Build a consumer goroutine that reads from audit-events and logs consumed messages. Add graceful shutdown with context cancellation. Introduce consumer group concepts.

## Week 3 — Cassandra
[TICKET-5] Cassandra schema — Design a keyspace and table optimised for audit log queries (partition by source, cluster by timestamp). Write a schema.cql migration script. Run Cassandra locally via docker-compose.

[TICKET-6] Cassandra persistence — Integrate gocql/gocql. Wire the Kafka consumer to write each consumed event to Cassandra. Handle retries and idempotency (use event id as primary key component).

[TICKET-7] Query API — Add GET /events?source=X&from=T&limit=N endpoint that reads from Cassandra and returns a paginated JSON response using Cassandra's paging token.

## Week 4 — Kubernetes
[TICKET-8] Dockerise — Write a multi-stage Dockerfile producing a minimal Go binary image. Add a docker-compose.yml that wires together the service, Kafka (+ Zookeeper), and Cassandra.

[TICKET-9] Kubernetes manifests — Write Deployment, Service, ConfigMap, and Secret manifests. Deploy to a local cluster (e.g. kind or minikube). Include liveness/readiness probes on the /healthz endpoint.

[TICKET-10] Kafka & Cassandra on Kubernetes — Deploy Kafka using the Strimzi operator and Cassandra using the K8ssandra operator (or Helm charts). Update the service ConfigMap to point at in-cluster addresses.

## Week 5 — Envoy
[TICKET-11] Envoy as ingress — Deploy Envoy as a standalone ingress in the cluster using a ConfigMap-backed envoy.yaml. Route external traffic to the audit service. Understand listener → filter chain → cluster config.

[TICKET-12] Envoy sidecar + observability — Add Envoy as a sidecar container in the service Deployment (manual injection, no Istio). Configure request tracing headers (x-request-id) and access logs. Wire up to a local Jaeger instance.

## Week 6 — Polish
[TICKET-13] Structured logging & metrics — Add slog structured logging (noting your familiarity with LoggerRequestContextMiddleware patterns) and expose a GET /metrics Prometheus endpoint for event throughput and Cassandra latency.

[TICKET-14] Integration tests — Write Go integration tests using testcontainers-go to spin up real Kafka and Cassandra in tests, covering the full ingest-to-query path.
