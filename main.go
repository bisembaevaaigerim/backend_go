package main

import (
	"backend-go/database"
	"backend-go/handlers"
	"backend-go/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/events", handlers.CreateEvent)
		protected.GET("/events", handlers.GetEvents)
		protected.GET("/events/:id", handlers.GetEvent)
		protected.PUT("/events/:id", handlers.UpdateEvent)
		protected.DELETE("/events/:id", handlers.DeleteEvent)

		protected.POST("/bookings", handlers.CreateBooking)
		protected.GET("/bookings", handlers.GetBookings)
		protected.GET("/bookings/:id", handlers.GetBooking)
		protected.PUT("/bookings/:id", handlers.UpdateBooking)
		protected.DELETE("/bookings/:id", handlers.DeleteBooking)
	}

	r.Run(":8080")
}
