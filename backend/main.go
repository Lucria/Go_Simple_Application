package main

import (
	"backend/controllers"
	"backend/database"
	"backend/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()

	// Init mock db
	database.Initialize()

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// Private Routes (Appointment)
	private := server.Group("/")
	private.Use(middleware.AuthCookieChecker)
	privateRoutes(private)

	// Public Routes (Register and Login)
	public := server.Group("/")
	publicRoutes(public)

	err := server.Run(":8080")
	if err != nil {
		log.Fatalf("Cannot start Gin server: %s", err)
	}
}

func publicRoutes(server *gin.RouterGroup) {
	server.POST("/auth/register", controllers.Register)
	server.POST("/auth/login", controllers.Login)
}

func privateRoutes(server *gin.RouterGroup) {
	server.GET("/appointments", controllers.GetAllAppointments)
	server.POST("/appointments", controllers.SearchForAvailableAppointments)
	server.GET("/appointment", controllers.GetAppointmentById)
	server.POST("/appointment", controllers.CreateAppointment)
	server.DELETE("/appointment", controllers.DeleteAppointment)
	server.PUT("/appointment", controllers.UpdateAppointment)
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}
