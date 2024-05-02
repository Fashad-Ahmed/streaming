package handler

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ddiogoo/ddiogoo/auth_server/internal/model"
	"github.com/ddiogoo/ddiogoo/auth_server/internal/service"
	"github.com/gin-gonic/gin"
)

type LiveStreamingKeyHandler struct {
	service service.ILiveStreamingKeyService
}

func NewLiveStreamingKeyHandler(service service.ILiveStreamingKeyService) ILiveStreamingKeyHandler {
	return &LiveStreamingKeyHandler{
		service,
	}
}

func (r *LiveStreamingKeyHandler) AuthStreamingKey(ctx *gin.Context) {
	body := ctx.Request.Body
	defer body.Close()

	parameters, err := io.ReadAll(body)
	if err != nil {
		panic("error on getting parameters")
	}

	liveStreamingKey := getStreamingKey(string(parameters))
	key, err := r.service.AuthStreamingKey(liveStreamingKey.Name, liveStreamingKey.StreamKey)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, "problem with streaming key")
		return
	}
	if key.StreamKey != "" {
		log.Default().Println("user authenticated")
		ctx.IndentedJSON(http.StatusOK, "OK")
		return
	}
	ctx.IndentedJSON(http.StatusForbidden, "Forbidden")
}

func getStreamingKey(parameters string) model.LiveStreamingKey {
	var liveStreamingKey model.LiveStreamingKey
	pairs := strings.Split(parameters, "&")
	for _, pair := range pairs {
		otherPair := strings.Split(pair, "=")
		key := otherPair[0]
		value := otherPair[1]
		if key == "name" {
			streamingKey := strings.Split(value, "_")
			liveStreamingKey.Name = streamingKey[0]
			liveStreamingKey.StreamKey = streamingKey[1]
		}
	}
	return liveStreamingKey
}
