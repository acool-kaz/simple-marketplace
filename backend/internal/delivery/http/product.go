package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) productCreateHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
	})
}
