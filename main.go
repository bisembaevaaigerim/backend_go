package main

import (
	"backend-go/database"
	"backend-go/handlers"
	"backend-go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	database.DB.AutoMigrate(&models.Event{}, &models.Booking{})

	r := gin.Default()

	r.POST("/events", handlers.CreateEvent)
	r.GET("/events", handlers.GetEvents)
	r.GET("/events/:id", handlers.GetEvent)
	r.PUT("/events/:id", handlers.UpdateEvent)
	r.DELETE("/events/:id", handlers.DeleteEvent)

	r.POST("/bookings", handlers.CreateBooking)
	r.GET("/bookings", handlers.GetBookings)
	r.GET("/bookings/:id", handlers.GetBooking)
	r.PUT("/bookings/:id", handlers.UpdateBooking)
	r.DELETE("/bookings/:id", handlers.DeleteBooking)
	
	r.Run(":8080")
}
