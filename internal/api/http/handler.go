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
		lists := api.Group("/users")
		{
			lists.GET("/all", h.getUsers)
		}
	}

	return router
}
