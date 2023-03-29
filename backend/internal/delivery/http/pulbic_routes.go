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

		product := publicRoutes.Group("/product")
		{
			product.GET("", h.getAllProductsInfoHandler)

			id := product.Group("/:id")
			{
				id.GET("", h.getProductInfoByIdHandler)
			}

			api := product.Group("/api", h.authMiddleware)
			{
				api.POST("", h.createProductHandler)

				id := api.Group("/:id")
				{
					id.DELETE("", h.deleteProductHandler)
					id.PATCH("", h.updateProductHandler)
				}
			}
		}
	}
}
