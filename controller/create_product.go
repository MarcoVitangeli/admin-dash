package controller

import (
	"errors"
	"fmt"
	"github.com/MarcoVitangeli/admin-dash/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCreateProduct(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create_product.html", nil)
}

func PostCreateProduct(ctx *gin.Context) {
	pdb, err := database.GetRepository()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

	name := ctx.Request.FormValue("product-name")
	if name == "" {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]any{
			"Message": "missing name parameter",
		})
		return
	}

	description := ctx.Request.FormValue("product-description")
	if description == "" {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]any{
			"Message": "missing description parameter",
		})
		return
	}

	categoryId := ctx.Request.FormValue("product-category")
	if categoryId == "" {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]any{
			"Message": "missing categoryId parameter",
		})
		return
	}

	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]any{
			"Message": fmt.Sprintf("invalid category id: %s", err),
		})
		return
	}

	newId, err := pdb.CreateProduct(ctx, name, description, categoryIdInt)

	if err != nil {
		var (
			msg        string
			statusCode int
		)

		msg = err.Error()

		//TODO: maybe build some kind of error mapper?
		if errors.Is(err, database.ErrCategoryNotFound) {
			statusCode = http.StatusNotFound
		} else if errors.Is(err, database.ErrDuplicatedProduct) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		ctx.HTML(statusCode, "error.html", map[string]any{
			"Message": msg,
		})
		return
	}

	ctx.HTML(http.StatusCreated, "success.html", map[string]any{
		"Message": fmt.Sprintf("successfully created product with ID: %d", newId),
	})
}
