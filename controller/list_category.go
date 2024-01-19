package controller

import (
    "github.com/MarcoVitangeli/admin-dash/database"
	"github.com/gin-gonic/gin"
    "net/http"
)

func GetListCategories(ctx *gin.Context) {
    pdb, err := database.GetRepository()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

    categories, err := pdb.GetCategories(ctx, 20)
 	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", map[string]any{
			"Message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "all_categories.html", map[string]any{
        "Categories" : categories,
    })
}
