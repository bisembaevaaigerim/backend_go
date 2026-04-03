package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model `json:"-"`
	Name       string
	Location   string
	Date       string
	Tickets    int
	Bookings   []Booking `gorm:"foreignKey:EventID"`
}
