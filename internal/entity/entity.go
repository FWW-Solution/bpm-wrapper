package entity

type Workflow struct {
	CaseID      int64  `db:"case_id"`
	TaskName    string `db:"task_name"`
	ProcessName string `db:"process_name"`
	IsActive    bool   `db:"is_active"`
	CreatedAt   string `db:"created_at"`
}
