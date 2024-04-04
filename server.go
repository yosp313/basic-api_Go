package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db = dbConfig()

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db.AutoMigrate(&Product{})

	router := http.NewServeMux()

	router.HandleFunc("/", nameController)
	router.HandleFunc("/ip", ipController)
	router.HandleFunc("POST /product", createProductController)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
	fmt.Printf("Server is listening on %s", server.Addr)
}

func nameController(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path[1:] == "" || len(r.URL.Path[1:]) < 3 {
		fmt.Fprintf(w, "Hello, Dumbass!")
		return
	}

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func ipController(w http.ResponseWriter, r *http.Request) {
	var ip string = ReadUserIP(r)

	fmt.Fprintf(w, "Your IP is: %s", ip)
}

func createProductController(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createProductService(&product)

	fmt.Fprintf(w, "The product have been created successfully!: %+v", product)
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func dbConfig() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connected to db")

	return db
}

func createProductService(product *Product) {
	db.Create(product)
}

func findProductByCodeService(product Product, productCode string) any {
	result := db.First(&product, "code = ?", productCode)

	return result
}
