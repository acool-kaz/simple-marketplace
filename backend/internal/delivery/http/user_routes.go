package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(basePath string, router *gin.Engine) {
	log.Println("init user routes")

	userRoutes := router.Group(basePath)
	{
		auth := userRoutes.Group("/auth")
		{
			auth.POST("/sign-up", h.signUpHandler)
			auth.POST("/sign-in", h.signInHandler)
		}

		/* product := userRoutes.Group("/product", h.authMiddleware)
		{
			product.POST("", h.createProductHandler)
			product.GET("", h.getAllProductsHandler)
			id := product.Group("/:id")
			{
				id.GET("", h.getProductByIdHandler)
				id.DELETE("", h.deleteProductHandler)
				id.PATCH("", h.updateProductHandler)
			}
		} */
	}
}
