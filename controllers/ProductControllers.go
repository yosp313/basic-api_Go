package controllers

import (
	"basic-api/configs"
	"basic-api/models"
	"basic-api/services"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB = configs.DbConfig()

func CreateProductController(w http.ResponseWriter, r *http.Request) {
  var product models.Product

  err := json.NewDecoder(r.Body).Decode( &product)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Invalid request body")
    return
  }

  services.CreateProductService(db, &product)

  fmt.Fprintf(w, "The product have been created successfully!: %+v", product)
}

func FindProductsController(w http.ResponseWriter, r *http.Request) {
  var product models.Product

  id := r.PathValue("id")

  result := services.FindProductByCodeService(db, &product, id)
  
  if result == nil {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Product not found")
    return
  }

  jsonData, err := json.Marshal(product)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Product not found")
    return
  }

  fmt.Fprintf(w, "%s", jsonData)
}

func ListAllProductsController(w http.ResponseWriter, r *http.Request) {
  var products []models.Product

  services.ListAllProductsService(db, &products)

  jsonData, err := json.Marshal(products)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Internal Server Error")
    return
  }

  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", jsonData)
}

func DeleteProductController(w http.ResponseWriter, r *http.Request) {
  var product models.Product

  id := r.PathValue("id")

  services.DeleteProductService(db, &product, id)

  fmt.Fprintf(w, "Deleted Product with ID: %s", id)
}

func UpdateProductController(w http.ResponseWriter, r *http.Request) {
  var product models.Product

  var updatedProduct models.Product

  id := r.PathValue("id")

  if 	err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Bad Request")
    return
  }


  services.UpdateProductService(db, &updatedProduct, &product, id)

  fmt.Fprintf(w, "The product has been updated successfully!")
}
