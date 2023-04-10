package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initFronendRoutes(basePath string, router *gin.Engine) {
	log.Println("init frontend routes")

	// router.LoadHTMLFiles(
	// 	"./public/404.html",
	// 	"./public/addProduct.html",
	// 	"./public/index.html",
	// 	"./public/product.html",
	// 	"./public/search.html",
	// 	"./public/signin.html",
	// 	"./public/signup.html",
	// )

	router.Static("/css", "./public/css")
	router.Static("/img", "./public/img")
	router.Static("/js", "./public/js")

	fronendRoutes := router.Group(basePath)
	{
		fronendRoutes.GET("", func(ctx *gin.Context) {
			// ctx.HTML(http.StatusOK, "index.html", nil)
			h.tmpl.ExecuteTemplate(ctx.Writer, "index.html", nil)
		})

		auth := fronendRoutes.Group("/auth")
		{
			auth.GET("/sign-up", func(ctx *gin.Context) {
				// ctx.HTML(http.StatusOK, "signup.html", nil)
				h.tmpl.ExecuteTemplate(ctx.Writer, "signup.html", nil)
			})

			auth.GET("/sign-in", func(ctx *gin.Context) {
				// ctx.HTML(http.StatusOK, "signin.html", nil)
				h.tmpl.ExecuteTemplate(ctx.Writer, "signin.html", nil)
			})
		}

		fronendRoutes.GET("/search", func(ctx *gin.Context) {
			// ctx.HTML(http.StatusOK, "search.html", nil)
			h.tmpl.ExecuteTemplate(ctx.Writer, "search.html", nil)
		})

		fronendRoutes.GET("/product", func(ctx *gin.Context) {
			// ctx.HTML(http.StatusOK, "product.html", nil)
			h.tmpl.ExecuteTemplate(ctx.Writer, "product.html", nil)
		})

		fronendRoutes.GET("/product/add", func(ctx *gin.Context) {
			// ctx.HTML(http.StatusOK, "addProduct.html", nil)
			h.tmpl.ExecuteTemplate(ctx.Writer, "addProduct.html", nil)
		})
	}
}
