package controllers

import (
	"backend/database"
	"backend/models"
	"backend/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
)

// Register
// User can sign up for a new account with credentials
func Register(context *gin.Context) {
	// TODO check already logged in and redirect

	var input models.RegistrationRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	newUser := models.User{}
	err := newUser.MapRegistrationRequestToUser(input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	services.SaveUser(&newUser)

	// Issue cookies and save session
	sessionId, err := uuid.NewV4()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	userCookie := &http.Cookie{
		Name:  "patientCookie",
		Value: sessionId.String(),
	}
	http.SetCookie(context.Writer, userCookie)
	database.SessionMap[userCookie.Value] = newUser.Username

	context.JSON(http.StatusCreated, gin.H{
		"newUser": newUser,
	})
	// TODO redirect back to normal page
}

// Login
// User attempts to log in with username and password
// Returns User if successful. Returns error if not
func Login(context *gin.Context) {
	// TODO if already logged in, redirect

	var authRequest models.AuthenticationRequest
	if err := context.ShouldBindJSON(&authRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user, err := services.VerifyLogin(authRequest)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
	}

	// Issue cookies and save session
	sessionId, err := uuid.NewV4()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	userCookie := &http.Cookie{
		Name:  "patientCookie",
		Value: sessionId.String(),
	}
	http.SetCookie(context.Writer, userCookie)
	database.SessionMap[userCookie.Value] = user.Username

	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	// TODO redirect back to normal page
}
