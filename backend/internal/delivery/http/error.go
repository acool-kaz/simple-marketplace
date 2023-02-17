package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func (h *Handler) errorHandler(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError

	switch {
	case errors.Is(err, models.ErrUserEmailExist):
		fallthrough
	case errors.Is(err, models.ErrUserEmailExist):
		fallthrough
	case errors.Is(err, models.ErrUserNotFound):
		status = http.StatusBadRequest
	case errors.Is(err, models.ErrInvalidAuthToken):
		status = http.StatusUnauthorized
	}

	info := gin.H{
		"status": status,
		"msg":    err.Error(),
	}

	ctx.AbortWithStatusJSON(status, info)

	jsonInfo, _ := json.MarshalIndent(info, "", " ")
	log.Println("\n", string(jsonInfo))
}
