package main

import (
	"basic-api/configs"
	"basic-api/controllers"
	"basic-api/models"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB = configs.DbConfig()

func main() {
	db.AutoMigrate(&models.Product{})

	router := http.NewServeMux()

	router.HandleFunc("/", controllers.NameController)
	router.HandleFunc("/ip", controllers.IpController)
	router.HandleFunc("POST /product", controllers.CreateProductController)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
	fmt.Printf("Server is listening on %s", server.Addr)
}
