package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationRequest struct {
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		log.Printf("[NotificationService] %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	r.POST("/notify", func(c *gin.Context) {
		var req NotificationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("[NotificationService] Sending notification to user %d: %s", req.UserID, req.Message)

		c.JSON(http.StatusOK, gin.H{
			"status":  "sent",
			"user_id": req.UserID,
			"message": req.Message,
		})
	})

	log.Println("NotificationService running on :8085")
	r.Run(":8085")
}
