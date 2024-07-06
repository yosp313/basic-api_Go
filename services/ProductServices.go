package services

import (
	"basic-api/models"

	"gorm.io/gorm"
)


func CreateProductService(db *gorm.DB, product *models.Product) {
  db.Create(&product)
}

func FindProductByCodeService(db *gorm.DB, product *models.Product, productId string) any {

	result := db.First(&product, productId)

  if result.Error != nil {
    return nil
  }

	return result
}

func ListAllProductsService(db *gorm.DB, product *[]models.Product) any {
	result := db.Find(&product)
	return result
}

func DeleteProductService(db *gorm.DB, product *models.Product, productId string) {
	db.Delete(&product, productId)
}

func UpdateProductService(db *gorm.DB, updatedProduct *models.Product, product *models.Product, productId string) {
	db.First(&product, productId)

	product.Code = updatedProduct.Code
	product.Price = updatedProduct.Price

	db.Save(&product)
}
