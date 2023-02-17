package admin_routes

import (
	"log"
	"net/http"

	"github.com/acool-kaz/simple-marketplace/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	services *service.Service
}

func InitAdminHandler(svc *service.Service) *AdminHandler {
	log.Println("init admin handler")
	return &AdminHandler{
		services: svc,
	}
}

func (h *AdminHandler) InitAdminRoutes() http.Handler {
	log.Println("init admin routes")
	engine := gin.Default()

	return engine
}
