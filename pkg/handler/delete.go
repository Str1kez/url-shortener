package handler

import (
	"net/http"

	"github.com/Str1kez/url-shortener/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) delete(ctx *gin.Context) {
	var correctKey schema.InfoRequest

	if err := ctx.ShouldBindUri(&correctKey); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, schema.ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.Model.Delete(correctKey.SecretKey); err != nil {
		ctx.JSON(http.StatusInternalServerError, schema.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
