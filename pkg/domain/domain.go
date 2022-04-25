package domain

type WorkflowResponse struct {
	ID         int64  `json:"workflowId"`
	Status     string `json:"status"`
	Response   string `json:"Response"`
	ErrDetails string `json:"ErrorMessage"`
}

type TaskHistory struct {
	ID       int64  `json:"taskId"`
	Name     string `json:"taskName"`
	Type     string `json:"taskType"`
	Status   string `json:"taskStatus"`
	Duration string `json:"taskDuration"`
}

type LogData struct {
	URL      string `json:"url"`
	Type     string `json:"type"`
	Request  string `json:"request"`
	Response string `json:"response"`
	Duration string `json:"duration"`
}

type TaskLog struct {
	TaskHistory
	Logs []LogData `json:"logs"`
}

type WorkflowHistory struct {
	ID       int64         `json:"id"`
	Name     string        `json:"name"`
	Type     string        `json:"type"`
	Status   string        `json:"status"`
	Duration string        `json:"duration"`
	APICalls int           `json:"totalAPIs"`
	TaskList []TaskHistory `json:"tasks"`
}
