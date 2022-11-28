package main

import (
	"backend/controllers"
	"backend/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	loadEnv()

	// Init mock db
	database.Initialize()

	server := gin.Default()

	// TODO to be removed!
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Appointment
	server.GET("/appointments", controllers.GetAllAppointments)
	server.POST("/appointments", controllers.SearchForAvailableAppointments)
	server.GET("/appointment", controllers.GetAppointmentById)
	server.POST("/appointment", controllers.CreateAppointment)
	server.DELETE("/appointment", controllers.DeleteAppointment)
	server.PUT("/appointment", controllers.UpdateAppointment)

	// Register
	server.POST("/auth/register", controllers.Register)
	server.POST("/auth/login", controllers.Login)

	err := server.Run(":8080")
	if err != nil {
		log.Fatalf("Cannot start Gin server: %s", err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}
