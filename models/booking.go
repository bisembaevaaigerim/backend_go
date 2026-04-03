package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model `json:"-"`
	User       string
	EventID    uint
	Quantity   int
}
