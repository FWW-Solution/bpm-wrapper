package entity

type Workflow struct {
	CaseID      string `json:"case_id"`
	TaskName    string `json:"task_name"`
	ProcessName string `json:"process_name"`
	CreatedAt   string `json:"created_at"`
}
