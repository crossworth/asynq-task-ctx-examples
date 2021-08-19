package main

import (
	"log"

	"asynq-task-ctx-examples/examples"
	"asynq-task-ctx-examples/tasks"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	// pass value to task
	// mux := asynq.NewServeMux()
	// mux.Use(examples.PassValueToTask)
	// mux.HandleFunc(tasks.TypeEmailDelivery, examples.HandlePassValueToTask)

	// return value from task
	mux := asynq.NewServeMux()
	mux.Use(examples.GetValueFromTask)
	mux.HandleFunc(tasks.TypeEmailDelivery, examples.HandleGetValueFromTask)

	// tracing
	// mux := asynq.NewServeMux()
	// mux.Use(examples.TaskTracing)
	// mux.HandleFunc(tasks.TypeEmailDelivery, examples.HandleTaskWithTracing)

	// metrics
	// mux := asynq.NewServeMux()
	// mux.Use(examples.TaskMetrics)
	// mux.HandleFunc(tasks.TypeEmailDelivery, tasks.HandleEmailDeliveryTask)

	// logging
	// mux := asynq.NewServeMux()
	// mux.Use(examples.TaskLoggingFromContext)
	// mux.HandleFunc(tasks.TypeEmailDelivery, examples.HandleTaskWithLogging)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
