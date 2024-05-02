package handler

import (
	"github.com/ddiogoo/ddiogoo/auth_server/internal/config/db/connection"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/model"
)

func MigrationHandler() {
	pg, err := connection.ConfigDatabaseHandler()
	if err != nil {
		panic("failed to connect database")
	}
	pg.AutoMigrate(&model.LiveStreaming{})
	pg.Create(&model.LiveStreaming{Name: "LIVE DO CASSIMIRO", StreamKey: "19ef6fab-d328-455c-a313-fdb3d6474793"})
	pg.Create(&model.LiveStreaming{Name: "LIVE DO GAULES", StreamKey: "e1f6b986-9211-40dd-b96f-f6a5026cfc5c"})
}
