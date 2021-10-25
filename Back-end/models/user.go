package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	Fullname  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
}
