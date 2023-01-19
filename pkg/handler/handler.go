package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Handler struct {
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	apiRouter := router.Group(viper.GetString("prefix"))
	{
		apiRouter.POST("/make_shortener", h.shortenerHandler)
		apiRouter.GET("/:short_code", h.redirectHandler)
		apiRouter.GET("/health_check", h.healthCheckHandler)

		adminRouter := apiRouter.Group("/admin")
		{
			adminRouter.GET("/:secret_key", h.infoHandler)
			adminRouter.DELETE("/:secret_key", h.deleteHandler)
		}
	}

	return router
}
