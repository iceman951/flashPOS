package models

import "time"

type Product struct {
	ID           uint `gorm:"primary_key"`
	Category     string
	SerialNumber string
	Status       string
	Name         string
	Stock        int64
	Price        float64
	Image        string
	CreatedAt    time.Time
}
