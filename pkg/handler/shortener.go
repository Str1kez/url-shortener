package handler

import (
	"net/http"

	"github.com/Str1kez/url-shortener/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) shortener(ctx *gin.Context) {
	var json schema.ShortenerRequest

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, schema.ErrorResponse{Message: err.Error()})
		return
	}

	data, err := h.Model.Create(json.Url)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schema.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
