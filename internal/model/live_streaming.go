package model

import "gorm.io/gorm"

type LiveStreaming struct {
	gorm.Model
	Name      string `gorm:"not null"`
	StreamKey string `gorm:"primaryKey"`
}
