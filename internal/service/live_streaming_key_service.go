package service

import (
	"github.com/ddiogoo/ddiogoo/auth_server/internal/model"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/repository"
)

type LiveStreamingKeyService struct {
	repository repository.ILiveStreamingKeyRepository
}

func NewLiveStreamingKeyService(repo repository.ILiveStreamingKeyRepository) ILiveStreamingKeyService {
	return &LiveStreamingKeyService{
		repo,
	}
}

func (r *LiveStreamingKeyService) AuthStreamingKey(name, key string) (*model.LiveStreamingKey, error) {
	return r.repository.FindStreamKey(name, key)
}
