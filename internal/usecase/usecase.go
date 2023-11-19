package usecase

import (
	"bpm-wrapper/internal/adapter"
	"bpm-wrapper/internal/config"
	"bpm-wrapper/internal/data/dto"
	dtopassenger "bpm-wrapper/internal/data/dto_passenger"
	"bpm-wrapper/internal/repository"
	"sync"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

var mu sync.Mutex

type usecase struct {
	adapter adapter.Adapter
	cfg     *config.BonitaConfig
	redis   *redis.Client
	pub     message.Publisher
	repo    repository.Repository
}

type Usecase interface {
	// StartProcess
	StartProcess(processName string, version string, body interface{}) (string, error)
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
	// SaveWorkflow
	SaveWorkflow(body *dto.SaveWorkflowRequest) error

	// Passenger
	UpdatePassenger(body dtopassenger.Passenger) error
	StartProcessPassenger(processName string, version string, body dto.StartProcessPassengerRequest) (string, error)
}

func New(adapter *adapter.Adapter, cfg *config.BonitaConfig, redis *redis.Client, pub message.Publisher, repo repository.Repository) Usecase {
	return &usecase{
		adapter: *adapter,
		cfg:     cfg,
		redis:   redis,
		pub:     pub,
		repo:    repo,
	}
}
