package database

import (
	"backend/models"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var (
	UserList        []models.User
	AppointmentList []models.Appointment
	SessionMap      map[string]string
)

func Initialize() {
	testPassword, err := bcrypt.GenerateFromPassword([]byte("testPassword"), bcrypt.DefaultCost)
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	hashedAdminPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	SessionMap = map[string]string{}

	if err != nil {
		panic("Initialization of mock db failed")
	}

	UserList = []models.User{
		{
			Id:       uuid.Must(uuid.NewV4()),
			Username: "testUsername",
			Password: string(testPassword),
			Name:     "Tester1",
			Age:      28,
		},
		{
			Id:       uuid.Must(uuid.NewV4()),
			Username: "Admin",
			Password: string(hashedAdminPassword),
			Name:     "Admin",
			Age:      99,
		},
	}

	AppointmentList = []models.Appointment{
		{
			Id:            uuid.Must(uuid.NewV4()),
			Title:         "Some Title",
			Owner:         "Some Owner",
			StartDateTime: time.Time{},
			EndDateTime:   time.Time{},
		},
	}
}
