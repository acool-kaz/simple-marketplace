package http

import (
	"log"
	"net/http"
	"text/template"

	"github.com/acool-kaz/simple-marketplace/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	tmpl     *template.Template
}

func InitHandler(svc *service.Service) *Handler {
	log.Println("init handler")
	return &Handler{
		services: svc,
		tmpl:     template.Must(template.ParseGlob("./public/*.html")),
	}
}

func (h *Handler) InitRoutes() http.Handler {
	log.Println("init routes")

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "/web")
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	router.Static("/static", "./static")

	h.initFronendRoutes("/web", router)

	h.initPublicRoutes("", router)
	h.initAdminRoutes("/admin", router)

	return router
}
