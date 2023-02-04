package controllers

import (
	"net/http"
	"project/tes-teknis/structs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (idb *InDB) GetProduct(c *gin.Context) {
	var (
		products structs.Product
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&products).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": products,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetProducts(c *gin.Context) {
	var (
		products []structs.Product
		result   gin.H
	)
	idb.DB.Find(&products)
	if len(products) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": products,
			"count":  len(products),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateProduct(c *gin.Context) {
	var (
		product structs.Product
		result  gin.H
	)
	ID := uuid.New()
	product_id := uuid.New()
	product_name := c.PostForm("product_name")
	product_description := c.PostForm("product_description")
	product.ID = ID
	product.Product_id = product_id
	product.Product_name = product_name
	product.Product_description = product_description
	idb.DB.Create(&product)
	result = gin.H{
		"result": product,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateProduct(c *gin.Context) {
	id := c.Query("id")
	product_name := c.PostForm("product_name")
	product_description := c.PostForm("product_description")
	var (
		product    structs.Product
		newProduct structs.Product
		result     gin.H
	)

	err := idb.DB.First(&product, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newProduct.Product_name = product_name
	newProduct.Product_description = product_description
	err = idb.DB.Model(&product).Updates(newProduct).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteProduct(c *gin.Context) {
	var (
		product structs.Product
		result  gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&product, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&product).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)

}
