package usecase

import (
	"bpm-wrapper/internal/data/dto"
	dtopassenger "bpm-wrapper/internal/data/dto_passenger"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

// UpdatePassanger implements Usecase.
func (u *usecase) UpdatePassenger(body dtopassenger.Passenger) error {

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	ID := watermill.NewUUID()

	// publish to message broker
	err = u.pub.Publish("update_passanger_from_bpm", message.NewMessage(ID, payload))
	if err != nil {
		return err
	}

	return nil

}

// StartProcess implements Usecase
func (u *usecase) StartProcessPassenger(processName string, version string, body dto.StartProcessPassengerRequest) (string, error) {
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

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	result, err := u.adapter.CreateProcessInstance(&token, processId, jsonBody)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println("case id", result)
	return result, nil
}
