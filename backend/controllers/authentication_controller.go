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
	services.CheckForSession(context)

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
		Name:   "sessionCookie",
		Value:  sessionId.String(),
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(context.Writer, userCookie)
	database.SessionMap[userCookie.Value] = newUser.Username

	context.JSON(http.StatusOK, gin.H{
		"name": newUser.Name,
	})
}

// Login
// User attempts to log in with username and password
// Returns User if successful. Returns error if not
func Login(context *gin.Context) {
	services.CheckForSession(context)

	var authRequest models.AuthenticationRequest
	if err := context.ShouldBindJSON(&authRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := services.VerifyLogin(authRequest)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Issue cookies and save session
	sessionId, err := uuid.NewV4()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	userCookie := &http.Cookie{
		Name:   "sessionCookie",
		Value:  sessionId.String(),
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(context.Writer, userCookie)
	database.SessionMap[userCookie.Value] = user.Username

	context.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}
