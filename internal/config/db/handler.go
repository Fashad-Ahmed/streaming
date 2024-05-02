package db

import (
	"github.com/ddiogoo/ddiogoo/auth_server/internal/model"
)

func MigrationHandler() {
	pg, err := ConfigDatabaseHandler()
	if err != nil {
		panic("failed to connect database")
	}
	pg.AutoMigrate(&model.LiveStreamingKey{})
	pg.Create(model.NewLiveStreamingKey("minhalive"))
}
