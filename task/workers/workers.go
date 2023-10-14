package main 

import(
	
	"github.com/hibiken/asynq"
	"log"
	"github.com/imnulhaqueruman/quickstart/task"
)

func main() {
    srv := asynq.NewServer(
        asynq.RedisClientOpt{Addr: "localhost:6379"},
        asynq.Config{Concurrency: 10},
    )

	mux := asynq.NewServeMux()
    mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
    mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

    if err := srv.Run(mux); err != nil {
        log.Fatal(err)
    }
}

