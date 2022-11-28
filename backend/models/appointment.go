package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Appointment struct {
	Id            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Owner         string    `json:"owner"`
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
}

type AppointmentSearchRequest struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
