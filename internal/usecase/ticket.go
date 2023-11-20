package usecase

import (
	dtoticket "bpm-wrapper/internal/data/dto_ticket"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

// UpdateTicket implements Usecase.
func (u *usecase) UpdateTicket(body *dtoticket.RequestUpdateTicket) error {
	json, err := json.Marshal(body)
	if err != nil {
		return err
	}

	id := watermill.NewUUID()

	// publish to message broker
	err = u.pub.Publish("update_ticket_from_bpm", message.NewMessage(id, json))
	if err != nil {
		return err
	}

	return nil
}

// RedeemTicket implements Usecase.
func (u *usecase) RedeemTicket(body *dtoticket.RequestRedeemTicketToBPM) error {
	token, err := u.loginUser()
	if err != nil {
		log.Println("Error Login", err)
		return err
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Error Marshal", err)
		return err
	}

	latestTask, err := u.repo.FindLatestTaskByCaseID(body.CaseID)
	if err != nil {
		log.Println("Error FindLatestTaskByCaseID", err)
		return err
	}

	task, err := u.adapter.FindTaskByName(&token, body.CaseID, latestTask.TaskName)
	if err != nil {
		log.Println("Error FindTaskByName", err)
		return err
	}

	fmt.Println("task", task.ID)
	fmt.Println("json Body", jsonBody)
	err = u.adapter.ExecuteTask(&token, task.ID, jsonBody)
	if err != nil {
		log.Println("Error ExecuteTask", err)
		return err
	}

	return nil
}
