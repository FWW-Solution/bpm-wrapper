package usecase

import (
	"bpm-wrapper/internal/adapter"
	"bpm-wrapper/internal/config"
	"bpm-wrapper/internal/data/dto"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var mu sync.Mutex

type usecase struct {
	adapter adapter.Adapter
	cfg     *config.BonitaConfig
	redis   *redis.Client
}

func (u *usecase) loginUser() (*dto.LoginResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	if u.redis.Exists(ctx, "token").Val() != 1 && u.redis.Exists(ctx, "auth").Val() != 1 {
		token, err := u.adapter.Login(u.cfg.Username, u.cfg.Password)
		if err != nil {
			return nil, err
		}

		u.redis.Set(ctx, "token", token.BonitaToken, 1*time.Hour)
		u.redis.Set(ctx, "auth", token.BonitaAuth, 1*time.Hour)
	}

	return &dto.LoginResponse{
		BonitaToken: u.redis.Get(ctx, "token").Val(),
		BonitaAuth:  u.redis.Get(ctx, "auth").Val(),
	}, nil
}

// GetTaskID(taskName string) (string, error)
func (u *usecase) GetTaskID(taskName string, caseID int64) (string, error) {
	token, err := u.loginUser()
	if err != nil {
		return "", err
	}

	task, err := u.adapter.FindTaskByName(token, caseID, taskName)
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

	err = u.adapter.ExecuteTask(token, taskID, variables)
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

	result, err := u.adapter.FindUser(token, actorName)
	if err != nil {
		return err
	}

	err = u.adapter.AssignTask(token, taskID, result[0].ID)
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

	err = u.adapter.ExecuteTask(token, taskID, variables)
	if err != nil {
		return err
	}

	return nil
}

// StartProcess implements Usecase
func (u *usecase) StartProcess(version string) (string, error) {
	token, err := u.loginUser()
	if err != nil {
		return "", err
	}

	// check if variables if not empty
	// TODO: Broken need check further
	// if len(variables["process_id"].(string)) > 0 {
	// 	result, err := u.adapter.CreateProcessInstance(&token, variables["process_id"].(string), variables)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return result, nil
	// }

	processId, err := u.adapter.FindProcess(token, version)
	if err != nil {
		return "", err
	}

	result, err := u.adapter.CreateProcessInstance(token, processId, nil)
	if err != nil {
		return "", err
	}
	return result, nil
}

// StopProcess implements Usecase
func (*usecase) StopProcess(token string, processId string) error {
	panic("unimplemented")
}

type Usecase interface {
	// StartProcess
	StartProcess(version string) (string, error)
	// StopProcess
	StopProcess(token string, processId string) error
	// Update Human Process
	UpdateHumanProcess(processId string, variables interface{}) error
	// Assign Human Task
	AssignHumanTask(taskID string, caseID int64, userID string) error
	// ExecuteHumanTask
	ExecuteHumanTask(taskID string, caseID int64, variables interface{}) error
	// GetTaskID
	GetTaskID(taskName string, caseID int64) (string, error)
}

func NewUsecase(adapter *adapter.Adapter, cfg *config.BonitaConfig, redis *redis.Client) Usecase {
	return &usecase{
		adapter: *adapter,
		cfg:     cfg,
		redis:   redis,
	}
}
