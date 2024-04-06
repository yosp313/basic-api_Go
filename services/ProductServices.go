package services

import (
	"basic-api/configs"
	"basic-api/models"

	"gorm.io/gorm"
)

var db *gorm.DB = configs.DbConfig()

func CreateProductService(product *models.Product) {
	db.Create(product)
}

func FindProductByCodeService(product *models.Product, productId string) any {

	result := db.First(&product, productId)

	return result
}

func ListAllProductsService(product *[]models.Product) any {
	result := db.Select("Code", "Price").Find(&product)
	return result
}

func DeleteProductService(product *models.Product, productId string) {
	db.Delete(&product, productId)
}

func UpdateProductService(updatedProduct *models.Product, product *models.Product, productId string) {
	db.First(&product, productId)

	product.Code = updatedProduct.Code
	product.Price = updatedProduct.Price

	db.Save(&product)
}
