package handler

import (
	"test-zero-agency/package/service"

	_ "test-zero-agency/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.handleSingUp)
		auth.POST("/log-in", h.handleSingIn)
	}

	list := router.Group("/list", h.userIdentify)
	{
		list.GET("", h.handleGetAllNews)
		list.GET("/:Id", h.handleGetNewsById)
	}

	edit := router.Group("/edit", h.userIdentify)
	{
		edit.POST("/:Id", h.handleUpdateNewsById)
		edit.DELETE("/:Id", h.handleDeleteNewsById)
	}

	return router
}
