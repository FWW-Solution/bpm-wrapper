package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required" example:"walter.bates"`
	Password string `json:"password" validate:"required" example:"bpm"`
}

type ProcessInstanceRequest struct {
	TicketNumber string `json:"ticket_number" validate:"required" example:"TMS-123456789"`
}

type StartProcessRequest struct {
	Version string `json:"version" validate:"required" example:"1.0"`
}

type AssignTaskRequest struct {
	AssignedID string `json:"assigned_id" validate:"required" example:"1"`
}

type ExecuteTaskRequest struct {
	Variables interface{} `json:"variables" validate:"required" example:"1"`
}

type GetTaskIDRequest struct {
	CaseID   int64  `json:"case_id" validate:"required" example:"1"`
	TaskName string `json:"task_name" validate:"required" example:"Review and analysis email"`
}

type ExecuteHumanTaskRequest struct {
	Variables interface{} `json:"variable" validate:"required" example:"1"`
	TaskID    string      `json:"task_id" validate:"required" example:"1235322345"`
	CaseID    int64       `json:"case_id" validate:"required" example:"1"`
}

type UpdateHumanProcessRequest struct {
	TaskID    string      `json:"task_id" validate:"required" example:"1"`
	Variables interface{} `json:"variables" validate:"required" example:"1"`
}

type AssignHumanTaskRequest struct {
	TaskID string `json:"task_id" validate:"required" example:"1"`
	CaseID int64  `json:"case_id" validate:"required" example:"1"`
	Actor  string `json:"actor" validate:"required" example:"L0"`
}
