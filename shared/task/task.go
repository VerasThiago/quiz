package task

type Task interface {
	CreateReportAync(submissionID string) error
}
