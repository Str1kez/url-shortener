package handler

import (
	"net/http"

	"github.com/Str1kez/url-shortener/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) redirect(ctx *gin.Context) {
	shortUrl := ctx.Param("short_code")

	longUrl, err := h.Model.Get(shortUrl)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schema.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, longUrl)
}
