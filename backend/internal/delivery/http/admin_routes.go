package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminRoutes(basePath string, router *gin.Engine) {
	log.Println("init admin routes")

	adminRoutes := router.Group(basePath)
	{
		adminRoutes.GET("/", func(c *gin.Context) {
			c.Writer.Write([]byte("hello world"))
		})
	}
}
