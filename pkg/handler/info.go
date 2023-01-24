package handler

import (
	"errors"
	"net/http"

	"github.com/Str1kez/url-shortener/pkg/db"
	"github.com/Str1kez/url-shortener/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) info(ctx *gin.Context) {
	var correctKey schema.InfoRequest
	var errNoResult *db.NoResultFound

	if err := ctx.ShouldBindUri(&correctKey); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, schema.ErrorResponse{Message: err.Error()})
		return
	}

	json, err := h.Model.GetInfo(correctKey.SecretKey)

	if err != nil {
		status := http.StatusInternalServerError
		if errors.As(err, &errNoResult) {
			status = http.StatusNotFound
		}
		ctx.JSON(status, schema.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, json)
}
