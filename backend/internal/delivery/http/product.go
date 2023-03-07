package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createProductHandler(ctx *gin.Context) {
	user := h.getUserFromCtx(ctx)

	var product models.ProductCreate
	err := ctx.Bind(&product)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	id, err := h.services.Product.Create(ctx.Request.Context(), user, product)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"product_id": id,
	})
}

func (h *Handler) getAllProductsHandler(ctx *gin.Context) {
	sortBy := ctx.Query("sortBy")
	if sortBy != "" {
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), models.ProductSortBy, sortBy))
	} else {
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), models.ProductSortBy, "id.asc"))
	}

	products, err := h.services.Product.GetAll(ctx.Request.Context())
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func (h *Handler) getProductByIdHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorHandler(ctx, models.ErrProductNotFound)
		return
	}
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), models.ProductId, id))

	product, err := h.services.Product.GetOneBy(ctx.Request.Context())
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func (h *Handler) deleteProductHandler(ctx *gin.Context) {
	user := h.getUserFromCtx(ctx)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorHandler(ctx, models.ErrProductNotFound)
		return
	}

	err = h.services.Product.Delete(ctx.Request.Context(), user, uint(id))
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "product deleted",
	})
}

func (h *Handler) updateProductHandler(ctx *gin.Context) {
	user := h.getUserFromCtx(ctx)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorHandler(ctx, models.ErrProductNotFound)
		return
	}

	var info models.ProductUpdate
	err = ctx.BindJSON(&info)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	product, err := h.services.Product.Update(ctx.Request.Context(), user, uint(id), info)
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
