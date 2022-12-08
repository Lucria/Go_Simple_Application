package services

import (
	"backend/database"
	. "backend/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func CheckForSession(context *gin.Context) {
	cookie, err := context.Cookie("sessionCookie")
	user, isPresent := database.SessionMap[cookie]
	if err == nil && isPresent {
		log.Printf("User already logged in as %s\n", user)
		context.Abort()
	}
}

func VerifyLogin(authRequest AuthenticationRequest) (*User, error) {
	for i := range database.UserList {
		if database.UserList[i].Username == authRequest.Username {
			user := database.UserList[i]

			err := user.ValidatePassword(authRequest.Password)

			if err != nil {
				return nil, err
			}
			return &user, err
		}
	}
	return nil, errors.New("no user found")
}

func SaveUser(newUser *User) {
	database.UserList = append(database.UserList, *newUser)
	fmt.Printf("%v\n", database.UserList)
}
