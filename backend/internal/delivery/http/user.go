package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) adminCreateUserHandler(ctx *gin.Context) {
	var user models.UserSignUp
	err := ctx.BindJSON(&user)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	id, err := h.services.User.Create(ctx.Request.Context(), user)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user_id": id,
	})
}

func (h *Handler) adminGetAllUsersHandler(ctx *gin.Context) {
	users, err := h.services.User.GetAll(ctx.Request.Context())
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *Handler) adminGetUserByIdHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorHandler(ctx, models.ErrUserNotFound)
		return
	}

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), models.UserId, uint(id)))
	user, err := h.services.User.GetOneBy(ctx.Request.Context())
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *Handler) adminDeleteUserHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorHandler(ctx, models.ErrUserNotFound)
		return
	}

	err = h.services.User.Delete(ctx.Request.Context(), uint(id))
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "user deleted",
	})
}

func (h *Handler) adminUpdateUserHanler(ctx *gin.Context) {
	var info models.UserUpdate
	err := ctx.BindJSON(&info)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorHandler(ctx, models.ErrUserNotFound)
		return
	}

	user, err := h.services.User.Update(ctx.Request.Context(), uint(id), info)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
