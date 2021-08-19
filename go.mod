module asynq-task-ctx-examples

go 1.17

require (
	github.com/hibiken/asynq v0.18.4
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.23.0
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.22.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
	go.opentelemetry.io/otel/trace v1.0.0-RC2
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/go-redis/redis/v7 v7.4.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	go.opentelemetry.io/otel/metric v0.22.0 // indirect
	go.opentelemetry.io/otel/sdk/export/metric v0.22.0 // indirect
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/hibiken/asynq => github.com/crossworth/asynq v0.18.3-0.20210819144548-b6e254f82b6b
