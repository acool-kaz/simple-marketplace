package http

import (
	"context"
	"fmt"
	"strings"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

type authCtx string

const curUserClaims authCtx = "cur_user_claims"

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

	claims, err := h.services.Auth.ParseToken(ctx.Request.Context(), accessToken)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), curUserClaims, claims))

	ctx.Next()
}

func (h *Handler) checkIfAdminMiddleware(ctx *gin.Context) {
	userClaims := ctx.Request.Context().Value(curUserClaims)
	userClaims = userClaims.(*models.Token)

	fmt.Println(userClaims.Id, userClaims.Role)

	ctx.Next()
}
