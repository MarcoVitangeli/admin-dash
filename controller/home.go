package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", nil)
}
