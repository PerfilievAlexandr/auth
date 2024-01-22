package http

import (
	"github.com/PerfilievAlexandr/auth/internal/api/http/dto"
	"github.com/PerfilievAlexandr/auth/internal/api/http/mapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.userService.GetAll(c.Request.Context())

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var dtoStudents []*dto.UserApi
	for _, user := range users {
		mappedStudent := mapper.MapUserToApiDto(user)
		dtoStudents = append(dtoStudents, mappedStudent)
	}

	c.JSON(http.StatusOK, dtoStudents)
}
