package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LiveStreamingKey struct {
	gorm.Model
	Name      string `gorm:"not null"`
	StreamKey string `gorm:"primaryKey"`
}

func NewLiveStreamingKey(name string) *LiveStreamingKey {
	return &LiveStreamingKey{
		Name:      name,
		StreamKey: uuid.New().String(),
	}
}
