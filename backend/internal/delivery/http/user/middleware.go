package user_routes

import (
	"context"
	"strings"

	delivery_http "github.com/acool-kaz/simple-marketplace/internal/delivery/http"
	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

type authCtx string

const curUserId authCtx = "cur_user_id"

func (h *UserHandler) authMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader(authorizationHeader)

	if token == "" {
		delivery_http.ErrorHandler(ctx, models.ErrInvalidAuthToken)
		return
	}

	headerParts := strings.Split(token, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		delivery_http.ErrorHandler(ctx, models.ErrInvalidAuthToken)
		return
	}

	accessToken := headerParts[1]

	userId, err := h.services.Auth.ParseToken(ctx.Request.Context(), accessToken)
	if err != nil {
		delivery_http.ErrorHandler(ctx, err)
		return
	}

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), curUserId, userId))

	ctx.Next()
}
