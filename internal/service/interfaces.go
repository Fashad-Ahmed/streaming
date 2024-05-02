package service

import "github.com/ddiogoo/ddiogoo/auth_server/internal/model"

type ILiveStreamingKeyService interface {
	AuthStreamingKey(name, key string) (*model.LiveStreamingKey, error)
}
