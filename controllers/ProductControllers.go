package controllers

import (
  "basic-api/models"
  "basic-api/services"
  "encoding/json"
  "fmt"
  "net/http"
)

func CreateProductController(w http.ResponseWriter, r *http.Request) {
  var product models.Product

  err := json.NewDecoder(r.Body).Decode(&product)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Invalid request body")
    return
  }

  services.CreateProductService(&product)

  fmt.Fprintf(w, "The product have been created successfully!: %+v", product)
}

func FindProductsController(w http.ResponseWriter, r *http.Request) {
  var product models.Product

  id := r.PathValue("id")

  services.FindProductByCodeService(&product, id)
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

  services.ListAllProductsService(&products)

  jsonData, err := json.Marshal(products)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Internal Server Error")
    return
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(products)
  fmt.Fprintf(w, "%s\n", jsonData)
}

func DeleteProductController(w http.ResponseWriter, r *http.Request) {
  var product models.Product

  id := r.PathValue("id")

  services.DeleteProductService(&product, id)

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


  services.UpdateProductService(&updatedProduct, &product, id)

  fmt.Fprintf(w, "The product has been updated successfully!")
}
