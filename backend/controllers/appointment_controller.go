package controllers

import (
	"backend/database"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetAllAppointments
// Get all appointments for specific user
func GetAllAppointments(c *gin.Context) {
	// Cookie authentication checked by middleware
	cookie, _ := c.Cookie("sessionCookie")
	usernameOfUser, _ := database.SessionMap[cookie]
	var nameOfUser string
	for _, user := range database.UserList {
		if user.Username == usernameOfUser {
			nameOfUser = user.Name
		}
	}

	var userAppointments []models.Appointment
	for _, appointment := range database.AppointmentList {
		if appointment.Owner == nameOfUser {
			userAppointments = append(userAppointments, appointment)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"appointments": userAppointments,
	})
}

// SearchForAvailableAppointments
// Search for available time slots within a time period
func SearchForAvailableAppointments(c *gin.Context) {
	var request models.AppointmentSearchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO search for appointments within time frame

	c.Status(http.StatusOK)
}

// GetAppointmentById
// Search for an Appointment by Id
func GetAppointmentById(c *gin.Context) {
	id := c.Query("id")

	for _, appointment := range database.AppointmentList {
		if appointment.Id.String() == id {
			c.JSON(http.StatusOK, gin.H{
				"appointment": appointment,
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Errorf("unable to find appointment by id %s", id),
	})
	log.Printf("Attempted to find missing appointment by id: %s\n", id)
}

// CreateAppointment
// Creates a new Appointment
func CreateAppointment(c *gin.Context) {
	var newAppointment models.Appointment

	// Validate input
	if err := c.ShouldBindJSON(&newAppointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Append to database
	database.AppointmentList = append(database.AppointmentList, newAppointment)

	c.JSON(http.StatusCreated, newAppointment)
}

// DeleteAppointment
// Deletes an appointment by id
func DeleteAppointment(c *gin.Context) {
	id := c.Query("id")

	fmt.Println(id)
	for idx, appointment := range database.AppointmentList {
		if appointment.Id.String() == id {
			database.AppointmentList = append(database.AppointmentList[:idx], database.AppointmentList[idx+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Errorf("unable to find appointment to delete. Id: %s", id),
	})
}

// UpdateAppointment
// Updates an appointment by id
func UpdateAppointment(c *gin.Context) {
	id := c.Query("id")

	// TODO implement update
	fmt.Println(id)

	c.Status(http.StatusOK)
}
