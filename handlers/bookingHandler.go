package handlers

import (
	"backend-go/database"
	"backend-go/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var event models.Event
	if err := database.DB.First(&event, booking.EventID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event not found"})
		return
	}
	if event.Tickets < booking.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough tickets"})
		return
	}
	event.Tickets -= booking.Quantity
	database.DB.Save(&event)
	result := database.DB.Create(&booking)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add booking"})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func GetBookings(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	offset := (pageInt - 1) * limitInt

	var bookings []models.Booking
	var total int64

	database.DB.Model(&models.Booking{}).Count(&total)
	result := database.DB.Model(&models.Booking{}).
		Limit(limitInt).
		Offset(offset).
		Find(&bookings)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get bookings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       bookings,
		"total":      total,
		"page":       pageInt,
		"limit":      limitInt,
		"totalPages": int(math.Ceil(float64(total) / float64(limitInt))),
	})
}

func GetBooking(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking

	result := database.DB.First(&booking, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to get booking"})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking

	if err := database.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}
	var input models.Booking
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	booking.User = input.User
	booking.EventID = input.EventID
	booking.Quantity = input.Quantity
	result := database.DB.Save(&booking)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}
	c.JSON(http.StatusOK, booking)
}
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Booking{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
