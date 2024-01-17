package main

import (
	"github.com/MarcoVitangeli/admin-dash/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles(
		"./templates/html/home.html",
		"./templates/html/error.html",
		"./templates/html/success.html",
		"./templates/html/create_product.html",
		"./templates/html/create_category.html")

	r.Static("/static", "./static")

	r.GET("/", controller.Home)

	r.GET("/create-category", controller.GetCreateCategory)
	r.POST("/create-category", controller.PostCreateCategory)

	r.GET("/create-product", controller.GetCreateProduct)
	r.POST("/create-product", controller.PostCreateProduct)

	r.Run(":8080")
}
