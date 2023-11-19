package repository

import (
	"bpm-wrapper/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SaveWorkflow(workflow *entity.Workflow) error
}

type repository struct {
	db *sqlx.DB
}

// SaveWorkflow implements Repository.
func (r *repository) SaveWorkflow(workflow *entity.Workflow) error {
	query := `INSERT INTO workflow (case_id, task_name, process_name, is_active, created_at) VALUES ($1, $2, $3, $4, NOW())`

	_, err := r.db.Exec(query, workflow.CaseID, workflow.TaskName, workflow.ProcessName, workflow.IsActive)
	if err != nil {
		return err
	}

	return nil
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
