package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCreateCategory(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create_category.html", nil)
}

func PostCreateCategory(ctx *gin.Context) {
	name := ctx.Request.FormValue("name")
	if name == "" {
		ctx.HTML(http.StatusBadRequest, "error.html", map[string]any{})
	}
}
