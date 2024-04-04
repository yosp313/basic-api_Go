package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", nameController)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}

func nameController(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path[1:] == "" || len(r.URL.Path[1:]) < 3 {
		fmt.Fprintf(w, "Hello, Dumbass!")
		return
	}

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
