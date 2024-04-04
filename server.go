package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", nameController)
	router.HandleFunc("/ip", ipController)

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
