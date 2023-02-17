package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/acool-kaz/simple-marketplace/internal/service"
)

type Handler struct {
	services *service.Service
}

func InitHandler(svc *service.Service) *Handler {
	log.Println("init handler")
	return &Handler{
		services: svc,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	log.Println("init routes")
	engine := gin.Default()

	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	auth := engine.Group("/auth")
	{
		auth.POST("/sign-up", h.signUpHandler)
		auth.POST("/sign-in", h.signInHandler)
	}

	users := engine.Group("/users", h.authMiddleware)
	{
		users.GET("", h.getAllUsersHandler)
	}

	product := engine.Group("/product", h.authMiddleware)
	{
		product.POST("", h.productCreateHandler)
	}

	return engine
}
