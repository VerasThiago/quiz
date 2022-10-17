package asynctask

import (
	"fmt"

	"github.com/hibiken/asynq"
	shared "github.com/verasthiago/quiz/shared/flags"
	"github.com/verasthiago/quiz/shared/task"
)

const (
	TypeReportCreate = "report:create"
)

type AsyncQueue struct {
	t *asynq.Client
}

func (a *AsyncQueue) InitFromFlags(sharedFlags *shared.SharedFlags) task.Task {
	redisConnection := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%+v:%+v", sharedFlags.QueueHost, sharedFlags.QueuePort),
	}

	t := asynq.NewClient(redisConnection)
	return &AsyncQueue{
		t,
	}
}

func (a *AsyncQueue) CreateReportAync(submissionID string) error {
	task := asynq.NewTask(TypeReportCreate, []byte(submissionID))
	_, err := a.t.Enqueue(task, asynq.Queue("critical"))
	return err
}
