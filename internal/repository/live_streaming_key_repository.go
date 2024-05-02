package repository

import (
	"errors"

	"github.com/ddiogoo/ddiogoo/auth_server/internal/model"
	"gorm.io/gorm"
)

type LiveStreamingKeyRepository struct {
	*gorm.DB
}

func NewLiveStreamingKeyRepository(db *gorm.DB) ILiveStreamingKeyRepository {
	return &LiveStreamingKeyRepository{
		db,
	}
}

func (r *LiveStreamingKeyRepository) FindStreamKey(name, key string) (*model.LiveStreamingKey, error) {
	liveStreamingKey := &model.LiveStreamingKey{
		Name:      name,
		StreamKey: key,
	}
	result := r.DB.First(liveStreamingKey)
	if result.RowsAffected != 1 {
		return nil, errors.New("live streaming key not found")
	}
	if result.Error != nil {
		return nil, errors.New("error on query")
	}
	return liveStreamingKey, nil
}
