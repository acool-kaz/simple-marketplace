package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initFronendRoutes(basePath string, router *gin.Engine) {
	log.Println("init frontend routes")

	router.LoadHTMLFiles(
		"./public/404.html",
		"./public/addProduct.html",
		"./public/index.html",
		"./public/product.html",
		"./public/search.html",
		"./public/signin.html",
		"./public/signup.html",
	)

	router.Static("/css", "./public/css")
	router.Static("/img", "./public/img")
	router.Static("/js", "./public/js")

	fronendRoutes := router.Group(basePath)
	{
		fronendRoutes.GET("", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", nil)
		})

		auth := fronendRoutes.Group("/auth")
		{
			auth.GET("/sign-up", func(ctx *gin.Context) {
				ctx.HTML(http.StatusOK, "signup.html", nil)
			})

			auth.GET("/sign-in", func(ctx *gin.Context) {
				ctx.HTML(http.StatusOK, "signin.html", nil)
			})
		}
	}
}
