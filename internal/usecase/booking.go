package usecase

import (
	dtobooking "bpm-wrapper/internal/data/dto_booking"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

// UpdateBooking implements Usecase.
func (u *usecase) UpdateBooking(body *dtobooking.RequestUpdateBooking) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	id := watermill.NewUUID()

	// publish to message broker
	err = u.pub.Publish("update_booking_from_bpm", message.NewMessage(id, jsonBody))
	if err != nil {
		return err
	}

	return nil
}
