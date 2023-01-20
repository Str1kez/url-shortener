package handler

import (
	"net/http"

	"github.com/Str1kez/url-shortener/schema"
	"github.com/gin-gonic/gin"
)


func (h *Handler) shortener(ctx *gin.Context) {
  data, err := h.Model.Create("https://ya.ru")
  if err != nil {
    ctx.JSON(http.StatusUnprocessableEntity, schema.ErrorResponse{Message: err.Error()})
    return
  }
  ctx.JSON(http.StatusOK, data)
}
