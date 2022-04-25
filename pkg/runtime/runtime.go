package runtime

import "coding2fun.in/beam/pkg/domain"

// BeamRuntime beam contract for workflow integration
type BeamRuntime interface {
	StartWorkflow(workflowName, resourceId string, variables map[string]any) domain.WorkflowResponse

	ResumeWorkflow(workflowName string, variables map[string]any) domain.WorkflowResponse
}
