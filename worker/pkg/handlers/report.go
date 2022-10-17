package handlers

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/verasthiago/quiz/worker/pkg/builder"
)

type CreateReportAPI interface {
	Handler(context context.Context, task *asynq.Task) error
}

type CreateReportHandler struct {
	builder.Builder
}

func (c *CreateReportHandler) InitFromBuilder(builder builder.Builder) *CreateReportHandler {
	c.Builder = builder
	return c
}

func (c *CreateReportHandler) Handler(context context.Context, task *asynq.Task) error {
	submissionID := string(task.Payload()[:])
	return c.GetRepository().CreateReportBySubmissionID(submissionID)
}
