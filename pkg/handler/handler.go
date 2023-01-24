package handler

import (
	"github.com/Str1kez/url-shortener/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Handler struct {
	Model *db.DbModel
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	apiRouter := router.Group(viper.GetString("prefix"))
	{
		apiRouter.POST("/make_shortener", h.shortener)
		apiRouter.GET("/:short_code", h.redirect)
		apiRouter.GET("/health_check", h.healthCheck)

		adminRouter := apiRouter.Group("/admin")
		{
			adminRouter.GET("/:secret_key", h.info)
			adminRouter.DELETE("/:secret_key", h.delete)
		}
	}

	return router
}
