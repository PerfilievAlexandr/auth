package http

import (
	"github.com/PerfilievAlexandr/auth/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *Handler {
	return &Handler{userService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		lists := api.Group("/user")
		{
			lists.POST("/create", h.create)
			lists.GET("/all", h.getAll)
			lists.GET("/:id", h.getById)
			lists.PUT("/update/:id", h.updateById)
			lists.DELETE("/:id", h.deleteById)
		}
	}

	return router
}
