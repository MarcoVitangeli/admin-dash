package controller

import (
	"fmt"
	"github.com/MarcoVitangeli/admin-dash/database"
	"github.com/MarcoVitangeli/admin-dash/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetListCategories(ctx *gin.Context) {
	pdb, err := database.GetRepository()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

	categories, err := pdb.GetCategories(ctx)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "all_categories.html", map[string]any{
		"Categories": categories,
	})
}

func SearchCategories(ctx *gin.Context) {
	time.Sleep(time.Second * 2) //delay to show loading spinner
	var (
		categories []models.ProductCategory
		err        error
	)
	pdb, err := database.GetRepository()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

	if sq := ctx.Query("search"); sq == "" {
		categories, err = pdb.GetCategories(ctx)
	} else {
		fmt.Printf("SEARCH: %s\n", sq)
		categories, err = pdb.SearchCategories(ctx, sq)
	}

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "search_category.html", map[string]any{
		"Categories": categories,
	})
}
