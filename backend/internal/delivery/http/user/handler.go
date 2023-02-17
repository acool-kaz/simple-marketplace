package user_routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/acool-kaz/simple-marketplace/internal/service"
)

type UserHandler struct {
	services *service.Service
}

func InitUserHandler(svc *service.Service) *UserHandler {
	log.Println("init user handler")
	return &UserHandler{
		services: svc,
	}
}

func (h *UserHandler) InitUserRoutes() http.Handler {
	log.Println("init user routes")
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUpHandler)
		auth.POST("/sign-in", h.signInHandler)
	}

	router.Group("/product", h.authMiddleware)
	{
		// todo product create
	}

	return router
}
