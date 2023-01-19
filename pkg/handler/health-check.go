package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	Message string `json:"message"`
}

func (h *Handler) healthCheckHandler(ctx *gin.Context) {
	msg := message{"Hello, I'm fine"}
	ctx.JSON(http.StatusOK, msg)
}
