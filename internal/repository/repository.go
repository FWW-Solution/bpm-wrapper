package repository

import (
	"bpm-wrapper/internal/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SaveWorkflow(workflow *entity.Workflow) error
	FindLatestTaskByCaseID(caseID int64) (entity.Workflow, error)
}

type repository struct {
	db *sqlx.DB
}

// FindLatestTaskByCaseID implements Repository.
func (r *repository) FindLatestTaskByCaseID(caseID int64) (entity.Workflow, error) {
	query := fmt.Sprintf("SELECT case_id, task_name, process_name, is_active, created_at FROM workflow WHERE case_id = '%d' ORDER BY created_at DESC LIMIT 1", caseID)

	var row entity.Workflow
	result, err := r.db.Queryx(query)
	if err != nil {
		return entity.Workflow{}, err
	}

	for result.Next() {
		err := result.StructScan(&row)
		if err != nil {
			return entity.Workflow{}, err
		}
	}

	return row, nil
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
