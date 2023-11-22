package usecase

import (
	dtonotification "bpm-wrapper/internal/data/dto_notification"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

// SendEmailNotification implements Usecase.
func (u *usecase) SendEmailNotification(body *dtonotification.SendEmailRequest) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}
	id := watermill.NewUUID()

	err = u.pub.Publish("send_email_notification_from_bpm", message.NewMessage(id, payload))
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) SendNotification(body *dtonotification.Request) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}
	id := watermill.NewUUID()

	err = u.pub.Publish("send_notification_from_bpm", message.NewMessage(id, payload))
	if err != nil {
		return err
	}

	return nil
}
