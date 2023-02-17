package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUsersHandler(ctx *gin.Context) {
	users, err := h.services.User.GetAll(ctx.Request.Context())
	if err != nil {
		h.errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
