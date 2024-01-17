package controller

import (
	"errors"
	"fmt"
	"github.com/MarcoVitangeli/admin-dash/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCreateCategory(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create_category.html", nil)
}

func PostCreateCategory(ctx *gin.Context) {
	name := ctx.Request.FormValue("name")
	if name == "" {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]any{
			"Message": "missing name in request",
		})
		return
	}
	pr, err := database.GetRepository()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

	newId, err := pr.CreateCategory(ctx, name)
	if errors.Is(err, database.ErrDuplicatedCategory) {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]any{
			"Message": "product category already present in DB",
		})
		return
	}

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "success.html", map[string]any{
		"Message": fmt.Sprintf("sucessfully created category with ID %d", newId),
	})
}
