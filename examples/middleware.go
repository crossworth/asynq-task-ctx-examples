package examples

import (
	"context"
	"fmt"
	stdLog "log"

	"github.com/hibiken/asynq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func PassValueToTask(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(t *asynq.Task) error {

		t = t.WithContext(context.WithValue(t.Context(), "myKey", "value from middleware"))

		return h.ProcessTask(t)
	})
}

func HandlePassValueToTask(t *asynq.Task) error {
	stdLog.Printf("middleware value: %v\n", t.Context().Value("myKey"))
	return nil
}

var ctxKey = struct{}{}

// https://github.com/hibiken/asynq/issues/285
// and allow external implementation of
// https://github.com/hibiken/asynq/issues/265
func GetValueFromTask(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(t *asynq.Task) error {
		err := h.ProcessTask(t)
		stdLog.Printf("task returned: %v\n", t.Context().Value(ctxKey))
		return err
	})
}

func HandleGetValueFromTask(t *asynq.Task) error {
	*t = *t.WithContext(context.WithValue(t.Context(), ctxKey, "value from task"))
	return nil
}

var tracer = otel.Tracer("example.org/tasks")

func TaskTracing(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(t *asynq.Task) error {
		ctx, span := tracer.Start(t.Context(), fmt.Sprintf("middleware-task-%s", t.Type()))
		defer span.End()

		return h.ProcessTask(t.WithContext(ctx))
	})
}

func HandleTaskWithTracing(t *asynq.Task) error {
	ctx, span := tracer.Start(t.Context(), fmt.Sprintf("handling-task-%s", t.Type()))
	defer span.End()

	taskID, _ := asynq.GetTaskID(t.Context())
	span.SetAttributes(attribute.String("task_id", taskID))

	span.AddEvent("started process")

	// do some work

	span.AddEvent("done process")

	t = t.WithContext(ctx)
	return nil
}

// https://github.com/hibiken/asynq/wiki/Monitoring-and-Alerting
var processedCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "processed_tasks_total",
		Help: "The total number of processed tasks",
	},
	[]string{"task_type"},
)

func TaskMetrics(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(t *asynq.Task) error {
		err := h.ProcessTask(t)
		processedCounter.WithLabelValues(t.Type()).Inc()
		return err
	})
}

// https://github.com/rs/zerolog#integration-with-nethttp
func TaskLoggingFromContext(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(t *asynq.Task) error {
		l := log.With().Logger()
		t = t.WithContext(l.WithContext(t.Context()))

		// inside the task handler we use

		return h.ProcessTask(t)
	})
}

func HandleTaskWithLogging(t *asynq.Task) error {
	l := log.Ctx(t.Context())
	taskID, _ := asynq.GetTaskID(t.Context())
	l.Info().Str("task_type", t.Type()).Str("task_id", taskID).Msgf("running task")
	return nil
}
