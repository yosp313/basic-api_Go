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

func FindProductByCodeService(product models.Product, productCode string) any {
	result := db.First(&product, "code = ?", productCode)

	return result
}
