package middleware

import (
	"backend/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthCookieChecker(c *gin.Context) {
	// Retrieve session id from cookie
	cookie, err := c.Cookie("sessionCookie")
	if err != nil {
		log.Println("Session cookie not set")
		c.Redirect(http.StatusMovedPermanently, "/auth/login")
		c.Abort()
		return
	}
	_, isPresent := database.SessionMap[cookie]
	if !isPresent {
		log.Println("Session id is not present")
		c.Redirect(http.StatusMovedPermanently, "/auth/login")
		c.Abort()
		return
	}
	// Session is valid and user is present
	c.Next()
}
