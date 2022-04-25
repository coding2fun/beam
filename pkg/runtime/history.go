package runtime

import "coding2fun.in/beam/pkg/domain"

type BeamHistory interface {
	WorkflowHistory(resourceId string, onlyTasks bool) domain.WorkflowHistory

	PendingTask(resourceId string) domain.TaskHistory

	TaskHistory(resourceId, taskId string) domain.LogData
}
