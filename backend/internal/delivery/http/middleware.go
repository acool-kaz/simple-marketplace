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

const curUser authCtx = "cur_user"

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

	user, err := h.services.User.GetOneBy(context.WithValue(ctx.Request.Context(), models.UserId, claims.Id))
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), curUser, user))

	ctx.Next()
}

func (h *Handler) checkIfAdminMiddleware(ctx *gin.Context) {
	user := ctx.Request.Context().Value(curUser)
	user = user.(models.User)

	fmt.Println(user.Role)

	ctx.Next()
}
