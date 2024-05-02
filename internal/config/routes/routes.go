package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigRoutesHandler(router *gin.Engine) {
	router.POST("/auth", auth)
	router.GET("/healthcheck", healthCheck)
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
