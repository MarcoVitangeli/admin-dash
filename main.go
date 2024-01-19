package main

import (
	"github.com/MarcoVitangeli/admin-dash/controller"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"FormatDate": func(t time.Time) string {
			return t.Format("2006-01-02 15:04:05")
		},
	})
	r.LoadHTMLFiles(
		"./templates/html/home.html",
		"./templates/html/error.html",
		"./templates/html/success.html",
		"./templates/html/all_categories.html",
		"./templates/html/create_product.html",
		"./templates/html/search_category.html",
		"./templates/html/create_category.html")

	r.Static("/static", "./static")

	r.GET("/", controller.Home)

	r.GET("/create-category", controller.GetCreateCategory)
	r.POST("/create-category", controller.PostCreateCategory)

	r.GET("/create-product", controller.GetCreateProduct)
	r.POST("/create-product", controller.PostCreateProduct)

	r.GET("/categories", controller.GetListCategories)
	r.GET("/categories/search", controller.SearchCategories)

	log.Panic(r.Run(":8080"))
}
