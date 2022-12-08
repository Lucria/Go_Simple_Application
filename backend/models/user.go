package models

import (
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"` // Consider removing from here?
	Password string    `json:"password"` // Consider removing from struct?
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	IsAdmin  bool      `json:"isAdmin"`
}

type RegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age,string"`
}

type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) MapRegistrationRequestToUser(request RegistrationRequest) error {
	newUuid, err := uuid.NewV4()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Id = newUuid
	user.Username = request.Username
	user.Password = string(passwordHash)
	user.Name = request.Name
	user.Age = request.Age
	if request.Username == "Admin" {
		user.IsAdmin = true
	} else {
		user.IsAdmin = false
	}
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
