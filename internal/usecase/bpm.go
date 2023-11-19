package usecase

import (
	"bpm-wrapper/internal/data/dto"
	"bpm-wrapper/internal/entity"
	"context"
	"fmt"
	"log"
)

func (u *usecase) loginUser() (dto.LoginResponse, error) {
	// mu.Lock()
	// defer mu.Unlock()

	// if u.redis.Exists(ctx, "token").Val() != 1 && u.redis.Exists(ctx, "auth").Val() != 1 {
	token, err := u.adapter.Login(u.cfg.Username, u.cfg.Password)
	if err != nil {
		fmt.Println("Error Login", err)
		return dto.LoginResponse{}, err
	}

	// 	u.redis.Set(ctx, "token", token.BonitaToken, 1*time.Hour)
	// 	u.redis.Set(ctx, "auth", token.BonitaAuth, 1*time.Hour)
	// }

	// return dto.LoginResponse{
	// 	BonitaToken: u.redis.Get(ctx, "token").Val(),
	// 	BonitaAuth:  u.redis.Get(ctx, "auth").Val(),
	// }, nil

	return dto.LoginResponse{
		BonitaToken: token.BonitaToken,
		BonitaAuth:  token.BonitaAuth,
	}, nil
}

// GetTaskID(taskName string) (string, error)
func (u *usecase) GetTaskID(taskName string, caseID int64) (string, error) {
	token, err := u.loginUser()
	if err != nil {
		return "", err
	}

	task, err := u.adapter.FindTaskByName(&token, caseID, taskName)
	if err != nil {
		return "", err
	}

	if task.ID == "" {
		return "", fmt.Errorf("task not found")
	}

	return task.ID, nil
}

// ExecuteHumanTask implements Usecase
func (u *usecase) ExecuteHumanTask(taskID string, caseID int64, variables interface{}) error {
	token, err := u.loginUser()
	if err != nil {
		return err
	}

	err = u.adapter.ExecuteTask(&token, taskID, variables)
	if err != nil {
		return err
	}

	return nil
}

// AssignHumanTask implements Usecase
func (u *usecase) AssignHumanTask(taskID string, caseID int64, actorName string) error {
	token, err := u.loginUser()
	if err != nil {
		return err
	}

	result, err := u.adapter.FindUser(&token, actorName)
	if err != nil {
		return err
	}

	err = u.adapter.AssignTask(&token, taskID, result[0].ID)
	if err != nil {
		return err
	}

	return nil

}

var ctx = context.Background()

// UpdateHumanProcess implements Usecase
func (u *usecase) UpdateHumanProcess(taskID string, variables interface{}) error {
	token, err := u.loginUser()
	if err != nil {
		return err
	}

	err = u.adapter.ExecuteTask(&token, taskID, variables)
	if err != nil {
		return err
	}

	return nil
}

// StartProcess implements Usecase
func (u *usecase) StartProcess(processName string, version string, body interface{}) (string, error) {
	token, err := u.loginUser()
	if err != nil {
		log.Println(err)
		return "", err
	}

	processId, err := u.adapter.FindProcess(&token, processName, version)
	if err != nil {
		log.Println(err)
		return "", err
	}

	fmt.Println("Body", body)

	result, err := u.adapter.CreateProcessInstance(&token, processId, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println("case id", result)
	return result, nil
}

// SaveWorkflow implements Usecase.
func (u *usecase) SaveWorkflow(body *dto.SaveWorkflowRequest) error {
	entity := entity.Workflow{
		CaseID:      body.CaseID,
		TaskName:    body.TaskName,
		ProcessName: body.ProcessName,
	}
	err := u.repo.SaveWorkflow(&entity)
	if err != nil {
		return err
	}
	return nil
}

// StopProcess implements Usecase
func (*usecase) StopProcess(token string, processId string) error {
	panic("unimplemented")
}
