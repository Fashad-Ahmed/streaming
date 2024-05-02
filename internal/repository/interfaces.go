package repository

import "github.com/ddiogoo/ddiogoo/auth_server/internal/model"

type ILiveStreamingKeyRepository interface {
	FindStreamKey(name, key string) (*model.LiveStreamingKey, error)
}
