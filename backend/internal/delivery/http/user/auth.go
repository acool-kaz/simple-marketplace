package user_routes

import (
	"net/http"

	delivery_http "github.com/acool-kaz/simple-marketplace/internal/delivery/http"
	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) signUpHandler(ctx *gin.Context) {
	var user models.UserSignUp
	err := ctx.BindJSON(&user)
	if err != nil {
		delivery_http.ErrorHandler(ctx, err)
		return
	}

	id, err := h.services.Auth.SignUp(ctx.Request.Context(), user)
	if err != nil {
		delivery_http.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": id,
	})
}

func (h *UserHandler) signInHandler(ctx *gin.Context) {
	var user models.UserSignIn
	err := ctx.BindJSON(&user)
	if err != nil {
		delivery_http.ErrorHandler(ctx, err)
		return
	}

	access, refresh, err := h.services.Auth.SignIn(ctx.Request.Context(), user)
	if err != nil {
		delivery_http.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
