package http

import (
	"net/http"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUpHandler(ctx *gin.Context) {
	var user models.UserSignUp
	err := ctx.BindJSON(&user)
	if err != nil {
		h.errorHandler(ctx, err)
		return
	}

	id, err := h.services.Auth.SignUp(ctx.Request.Context(), user)
	if err != nil {
		h.errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": id,
	})
}

func (h *Handler) signInHandler(ctx *gin.Context) {
	var user models.UserSignIn
	err := ctx.BindJSON(&user)
	if err != nil {
		h.errorHandler(ctx, err)
		return
	}

	access, refresh, err := h.services.Auth.SignIn(ctx.Request.Context(), user)
	if err != nil {
		h.errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
