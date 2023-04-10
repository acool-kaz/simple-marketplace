package http

import (
	"context"
	"log"

	"github.com/acool-kaz/simple-marketplace/internal/models"
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
			new, err := h.services.Product.GetAllInfo(context.WithValue(ctx.Request.Context(), models.ProductSortBy, "id.desc"))
			if err != nil {
				errorHandler(ctx, err)
				return
			}

			men, err := h.services.Product.GetAllInfo(context.WithValue(ctx.Request.Context(), models.ProductTag, "men"))
			if err != nil {
				errorHandler(ctx, err)
				return
			}

			women, err := h.services.Product.GetAllInfo(context.WithValue(ctx.Request.Context(), models.ProductTag, "women"))
			if err != nil {
				errorHandler(ctx, err)
				return
			}

			info := struct {
				New   []models.ProductInfo
				Men   []models.ProductInfo
				Women []models.ProductInfo
			}{
				New:   new,
				Men:   men,
				Women: women,
			}

			h.tmpl.ExecuteTemplate(ctx.Writer, "index.html", info)
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
			searchBy := ctx.Query("search_by")
			if searchBy != "" {
				reqCtx := context.WithValue(ctx.Request.Context(), models.ProductSearchBy, searchBy)
				reqCtx = context.WithValue(reqCtx, models.ProductName, searchBy)
				reqCtx = context.WithValue(reqCtx, models.ProductShortDescription, searchBy)
				reqCtx = context.WithValue(reqCtx, models.ProductDescription, searchBy)
				reqCtx = context.WithValue(reqCtx, models.ProductTag, searchBy)
				reqCtx = context.WithValue(reqCtx, models.IsOrCtx, struct{}{})
				ctx.Request = ctx.Request.WithContext(reqCtx)
			}

			products, err := h.services.Product.GetAllInfo(ctx.Request.Context())
			if err != nil {
				errorHandler(ctx, err)
				return
			}

			info := struct {
				SearchBy string
				Search   []models.ProductInfo
			}{
				SearchBy: searchBy,
				Search:   products,
			}

			h.tmpl.ExecuteTemplate(ctx.Writer, "search.html", info)
		})

		fronendRoutes.GET("/product", func(ctx *gin.Context) {
			// ctx.HTML(http.StatusOK, "product.html", nil)
			ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), models.ProductId, ctx.Query("id")))

			curProduct, err := h.services.Product.GetAllInfo(ctx.Request.Context())
			if err != nil {
				errorHandler(ctx, err)
				return
			}

			men, err := h.services.Product.GetAllInfo(context.WithValue(ctx.Request.Context(), models.ProductTag, "men"))
			if err != nil {
				errorHandler(ctx, err)
				return
			}

			women, err := h.services.Product.GetAllInfo(context.WithValue(ctx.Request.Context(), models.ProductTag, "women"))
			if err != nil {
				errorHandler(ctx, err)
				return
			}

			info := struct {
				CurProduct models.ProductInfo
				Men        []models.ProductInfo
				Women      []models.ProductInfo
			}{
				CurProduct: curProduct[0],
				Men:        men,
				Women:      women,
			}

			h.tmpl.ExecuteTemplate(ctx.Writer, "product.html", info)
		})

		fronendRoutes.GET("/product/add", func(ctx *gin.Context) {
			// ctx.HTML(http.StatusOK, "addProduct.html", nil)
			h.tmpl.ExecuteTemplate(ctx.Writer, "addProduct.html", nil)
		})
	}
}
