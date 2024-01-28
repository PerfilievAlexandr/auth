package http

import (
	"encoding/json"
	"github.com/PerfilievAlexandr/auth/internal/api/http/dtoHttpUser"
	"github.com/PerfilievAlexandr/auth/internal/api/http/httpMapper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAll(ctx *gin.Context) {
	users, err := h.userService.GetAll(ctx.Request.Context())

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var dtoStudents []*dtoHttpUser.UserResponse
	for _, user := range users {
		mappedStudent := httpMapper.MapUserToApiDto(user)
		dtoStudents = append(dtoStudents, mappedStudent)
	}

	ctx.JSON(http.StatusOK, dtoStudents)
}

func (h *Handler) create(ctx *gin.Context) {
	var signUp dtoHttpUser.SignUpRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&signUp); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.userService.Create(ctx.Request.Context(), signUp)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, userId)
}

func (h *Handler) getById(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userService.Get(ctx.Request.Context(), userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, httpMapper.MapUserToApiDto(user))
}

func (h *Handler) updateById(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var req dtoHttpUser.UpdateRequest
	if err = json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.userService.Update(ctx, userId, req)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) deleteById(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.userService.Delete(ctx, userId)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
