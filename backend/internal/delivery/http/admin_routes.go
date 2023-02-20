package http

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminRoutes(basePath string, router *gin.Engine) {
	log.Println("init admin routes")

	adminRoutes := router.Group(basePath)
	{
		auth := adminRoutes.Group("/auth")
		{
			auth.POST("/sign-in", h.signInHandler)
		}

		api := adminRoutes.Group("/api", h.authMiddleware, h.checkIfAdminMiddleware)
		{
			user := api.Group("/user")
			{
				user.POST("", h.adminCreateUserHandler)
				user.GET("", h.adminGetAllUsersHandler)
				id := user.Group("/:id")
				{
					id.GET("", h.adminGetUserByIdHandler)
					id.DELETE("", h.adminDeleteUserHandler)
					id.PATCH("", h.adminUpdateUserHanler)
				}
			}

			product := api.Group("/product")
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
}
