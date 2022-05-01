package domain

type WorkflowResponse struct {
	ID         int64  `json:"workflowId"`
	Status     string `json:"status"`
	Response   string `json:"Response"`
	ErrDetails string `json:"ErrorMessage"`
}
