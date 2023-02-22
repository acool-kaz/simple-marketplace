package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
)

func errorHandler(ctx *gin.Context, err error) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	status := http.StatusInternalServerError

	switch {
	case errors.Is(err, models.ErrInvalidProduct):
		fallthrough
	case errors.Is(err, models.ErrUserUsernameExist):
		fallthrough
	case errors.Is(err, models.ErrUserEmailExist):
		status = http.StatusBadRequest
	case errors.Is(err, models.ErrProductNotFound):
		fallthrough
	case errors.Is(err, models.ErrUserNotFound):
		status = http.StatusNotFound
	case errors.Is(err, models.ErrNotAdmin):
		fallthrough
	case errors.Is(err, models.ErrInvalidAuthToken):
		status = http.StatusUnauthorized
	}

	info := gin.H{
		"status": status,
		"msg":    err.Error(),
	}

	ctx.AbortWithStatusJSON(status, info)

	log.Printf("\n%+v\n", info)
}
