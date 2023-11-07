package usecase

import (
	"bpm-wrapper/internal/config"
	"bpm-wrapper/internal/data/model"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type usecase struct {
	cache *redis.Client
	cfg   *config.Config
}

// ClaimTaskToUserPoll implements Usecase.
func (u *usecase) ClaimTaskToUserPoll(ctx context.Context, taskID int64, caseID int64, actorName string) error {
	// Check Login first to bpm
	_, err := u.Login(ctx, u.cfg.Bonita.Username, u.cfg.Bonita.Password)
	if err != nil {
		return err
	}

	// err = u.adapter.ClaimTaskToUserPoll(&token, taskID, caseID, actorName)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// ExecuteHumanTask implements Usecase.
func (u *usecase) ExecuteHumanTask(ctx context.Context, taskID int64, caseID int64, variables interface{}) error {
	// Check Login first to bpm
	_, err := u.Login(ctx, u.cfg.Bonita.Username, u.cfg.Bonita.Password)
	if err != nil {
		return err
	}

	// err = u.adapter.ExecuteHumanTask(&token, taskID, caseID, variables)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// GetTaskID implements Usecase.
func (u *usecase) GetTaskID(ctx context.Context, taskName string, caseID int64) (taskID int64, err error) {
	// Check Login first to bpm
	_, err = u.Login(ctx, u.cfg.Bonita.Username, u.cfg.Bonita.Password)
	if err != nil {
		return 0, err
	}

	// task, err := u.adapter.FindTaskByName(&token, caseID, taskName)
	// if err != nil {
	// 	return "", err
	// }
	return 0, nil
}

// Login implements Usecase.
func (u *usecase) Login(ctx context.Context, username string, password string) (loginModel interface{}, err error) {
	if u.cache.Exists(ctx, "token_bonita").Val() != 1 && u.cache.Exists(ctx, "auth_bonita").Val() != 1 {
		loginResult := model.LoginBonitaBPM{}
		// token, err := u.adapter.Login(u.cfg.Username, u.cfg.Password)
		// if err != nil {
		// 	return "", err
		// }

		u.cache.Set(ctx, "token", loginResult.BonitaToken, time.Duration(u.cfg.Bonita.LoginCacheDuration)*time.Hour)
		u.cache.Set(ctx, "auth", loginResult.BonitaAuth, time.Duration(u.cfg.Bonita.LoginCacheDuration)*time.Hour)
	}

	loginModel = model.LoginBonitaBPM{
		BonitaToken: u.cache.Get(ctx, "token_bonita").Val(),
		BonitaAuth:  u.cache.Get(ctx, "auth_bonita").Val(),
	}

	return loginModel, nil

}

// Logout implements Usecase.
func (*usecase) Logout(ctx context.Context) error {
	panic("unimplemented")
}

// RecordCurrentTask implements Usecase.
func (*usecase) RecordCurrentTask(ctx context.Context, taskID int64, caseID int64) error {
	panic("unimplemented")
}

// StartProcess implements Usecase.
func (u *usecase) StartProcess(ctx context.Context, version string) (caseID int64, err error) {
	// Check Login first to bpm
	_, err = u.Login(ctx, u.cfg.Bonita.Username, u.cfg.Bonita.Password)
	if err != nil {
		return 0, err
	}

	// processID, err := u.adapter.FindProcessIDByVersion(&token, version)
	// if err != nil {
	// 	return "", err
	// }

	// caseID, err = u.adapter.StartProcess(&token, processID)
	// if err != nil {
	// 	return "", err
	// }
	return 0, nil
}

type Usecase interface {
	Login(ctx context.Context, username string, password string) (loginModel interface{}, err error)
	Logout(ctx context.Context) error
	ExecuteHumanTask(ctx context.Context, taskID int64, caseID int64, variables interface{}) error
	StartProcess(ctx context.Context, version string) (caseID int64, err error)
	GetTaskID(ctx context.Context, taskName string, caseID int64) (taskID int64, err error)
	ClaimTaskToUserPoll(ctx context.Context, taskID int64, caseID int64, actorName string) error
	RecordCurrentTask(ctx context.Context, taskID int64, caseID int64) error
}

func NewBonitaBPM(cache *redis.Client, cfg *config.Config) Usecase {
	return &usecase{
		cache: cache,
		cfg:   cfg,
	}
}
