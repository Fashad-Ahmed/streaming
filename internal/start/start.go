package start

import (
	"github.com/ddiogoo/ddiogoo/auth_server/internal/config/db"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/config/env"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/config/routes"
	"github.com/gin-gonic/gin"
)

func StartApplication() *gin.Engine {
	env.EnvConfigurationHandler()
	db.MigrationHandler()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.ConfigRoutesHandler(router)
	return router
}
