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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.CreateProductService(&product)

	fmt.Fprintf(w, "The product have been created successfully!: %+v", product)
}
