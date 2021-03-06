package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
	"github.com/webapp/controllers"
	"github.com/webapp/db"
	"github.com/webapp/models"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.Get().DropTableIfExists(&models.Product{}, &models.Category{})
	db.Get().CreateTable(&models.Product{}, &models.Category{})

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/products", controllers.GetAllProducts)
	router.POST("/admin/add_product", controllers.AddProduct)
	router.POST("/admin/add_category", controllers.AddCategory)

	router.Run(":" + port)
}
