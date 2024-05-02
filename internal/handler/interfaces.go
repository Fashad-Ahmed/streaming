package handler

import "github.com/gin-gonic/gin"

type ILiveStreamingKeyHandler interface {
	AuthStreamingKey(ctx *gin.Context)
}
