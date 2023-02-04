package main

import (
	"project/tes-teknis/config"
	"project/tes-teknis/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/product/:id", inDB.GetProduct)
	router.GET("/products", inDB.GetProducts)
	router.POST("/product", inDB.CreateProduct)
	router.PUT("/product/:id", inDB.UpdateProduct)
	router.DELETE("/product/:id", inDB.DeleteProduct)
	router.Run(":3000")
}
