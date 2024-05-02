package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ddiogoo/ddiogoo/auth_server/internal/migration"
	"github.com/gin-gonic/gin"
)

func main() {
	migration.MigrationHandler()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/auth", auth)
	router.GET("/healthcheck", healthCheck)
	router.Run(":8000")
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Ta funfando viu, pode ficar tranquilo!")
}

func auth(c *gin.Context) {
	log.Println("Running Auth...")
	body := c.Request.Body
	defer body.Close()

	fields, _ := io.ReadAll(body)
	fmt.Println(string(fields))

	c.IndentedJSON(http.StatusOK, "Autenticado com sucesso!")
}
