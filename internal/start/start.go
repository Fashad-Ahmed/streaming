package start

import (
	"net/http"

	"github.com/ddiogoo/ddiogoo/auth_server/internal/config/db"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/config/env"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/handler"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/repository"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/service"
	"github.com/gin-gonic/gin"
)

func StartApplication() *gin.Engine {
	env.EnvConfigurationHandler()
	conn, err := db.ConfigDatabaseHandler()
	if err != nil {
		panic("error while connecting database")
	}
	db.MigrationHandler()

	repo := repository.NewLiveStreamingKeyRepository(conn)
	service := service.NewLiveStreamingKeyService(repo)
	streaming := handler.NewLiveStreamingKeyHandler(service)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/auth", streaming.AuthStreamingKey)
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Working successfully!")
	})
	return router
}
