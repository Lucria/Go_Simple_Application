package services

import (
	"backend/database"
	. "backend/models"
	"errors"
	"fmt"
)

func VerifyLogin(authRequest AuthenticationRequest) (*User, error) {
	for i := range database.UserList {
		if database.UserList[i].Username == authRequest.Username {
			user := database.UserList[i]

			err := user.ValidatePassword(authRequest.Password)

			if err != nil {
				return nil, err
			}
		}
	}
	return nil, errors.New("no user found")
}

func SaveUser(newUser *User) {
	test := append(database.UserList, *newUser)
	fmt.Printf("%v\n", test)
}
