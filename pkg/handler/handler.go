package handler

import (
	"todoo/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.GetAllItems)
			lists.GET("/:id", h.GetListById)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteItem)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.GetAllItems)
				items.GET("/:item_id", h.GetItemById)
				items.PUT("/:item_id", h.UpdateItem)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}
	}
	return router
}
