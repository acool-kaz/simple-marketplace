package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initPublicRoutes(basePath string, router *gin.Engine) {
	log.Println("init public routes")

	publicRoutes := router.Group(basePath)
	{
		auth := publicRoutes.Group("/auth")
		{
			auth.POST("/sign-up", h.signUpHandler)
			auth.POST("/sign-in", h.signInHandler)
		}

		product := publicRoutes.Group("/product", h.authMiddleware)
		{
			product.POST("", h.createProductHandler)
			product.GET("", h.getAllProductsHandler)
			id := product.Group("/:id")
			{
				id.GET("", h.getProductByIdHandler)
				id.DELETE("", h.deleteProductHandler)
				id.PATCH("", h.updateProductHandler)
			}
		}
	}
}
