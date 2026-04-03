package handlers

import (
	"backend-go/database"
	"backend-go/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var newEvent models.Event

	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := database.DB.Create(&newEvent)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add event"})
		return
	}
	c.JSON(http.StatusCreated, newEvent)
}

func GetEvents(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	offset := (pageInt - 1) * limitInt

	var events []models.Event
	var total int64

	database.DB.Model(&models.Event{}).Count(&total)

	result := database.DB.Model(&models.Event{}).
		Limit(limitInt).
		Offset(offset).
		Find(&events)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       events,
		"total":      total,
		"page":       pageInt,
		"limit":      limitInt,
		"totalPages": int(math.Ceil(float64(total) / float64(limitInt))),
	})
}

func GetEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event

	result := database.DB.First(&event, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")

	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var input models.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	event.Name = input.Name
	event.Location = input.Location
	event.Date = input.Date
	event.Tickets = input.Tickets

	result := database.DB.Save(&event)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Event{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
