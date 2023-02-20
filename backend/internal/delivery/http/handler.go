package http

import (
	"log"
	"net/http"

	"github.com/acool-kaz/simple-marketplace/internal/service"
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

	h.initPublicRoutes("", router)
	h.initAdminRoutes("/admin", router)

	return router
}
