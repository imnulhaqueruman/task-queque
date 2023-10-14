package main 

import(
	
	"github.com/hibiken/asynq"
	"log"
	"time"
    "github.com/imnulhaqueruman/quickstart/task"

)

type EmailTaskPayload struct {
	UserID int
}

func main() {
    client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

    // Create a task with typename and payload.
	t1, err := task.NewWelcomeEmailTask(42)
    if err != nil {
        log.Fatal(err)
    }

    t2, err := task.NewReminderEmailTask(42)
    if err != nil {
        log.Fatal(err)
    }

    // Process the task immediately.
    info, err := client.Enqueue(t1)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(" [*] Successfully enqueued task: %+v", info)

    // Process the task 24 hours later.
    info, err = client.Enqueue(t2, asynq.ProcessIn(24*time.Hour))
    if err != nil {
        log.Fatal(err)
    }
    log.Println(" [*] Successfully enqueued task: %+v", info)
}