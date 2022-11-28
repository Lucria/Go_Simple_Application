package controllers

import (
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllAppointments
// Get all appointments for specific user
func GetAllAppointments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting all appointments for user!",
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

	// TODO implement get by id
	fmt.Println(id)

	c.Status(http.StatusOK)
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

	// TODO Append to database

	c.JSON(http.StatusCreated, newAppointment)
}

// DeleteAppointment
// Deletes an appointment by id
func DeleteAppointment(c *gin.Context) {
	id := c.Query("id")

	// TODO implement delete
	fmt.Println(id)

	c.Status(http.StatusNoContent)
}

// UpdateAppointment
// Updates an appointment by id
func UpdateAppointment(c *gin.Context) {
	id := c.Query("id")

	// TODO implement update
	fmt.Println(id)

	c.Status(http.StatusOK)
}
