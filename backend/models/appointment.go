package models

import (
	"github.com/gofrs/uuid"
)

type Appointment struct {
	Id            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Owner         string    `json:"owner"`
	StartDateTime int64     `json:"startDateTime"`
	EndDateTime   int64     `json:"endDateTime"`
}
