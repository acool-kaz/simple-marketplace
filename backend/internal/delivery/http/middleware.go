package http

import (
	"context"
	"strings"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

type authCtx string

const curUserId authCtx = "cur_user_id"

func (h *Handler) authMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader(authorizationHeader)

	if token == "" {
		errorHandler(ctx, models.ErrInvalidAuthToken)
		return
	}

	headerParts := strings.Split(token, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		errorHandler(ctx, models.ErrInvalidAuthToken)
		return
	}

	accessToken := headerParts[1]

	userId, err := h.services.Auth.ParseToken(ctx.Request.Context(), accessToken)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), curUserId, userId))

	ctx.Next()
}
