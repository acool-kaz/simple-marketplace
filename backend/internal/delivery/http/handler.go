package http

import (
	"log"
	"net/http"

	"github.com/acool-kaz/simple-marketplace/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	router.Static("/static", "./static")

	h.initPublicRoutes("", router)
	h.initAdminRoutes("/admin", router)

	return router
}
