package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path[1:] == "" || len(r.URL.Path[1:]) < 3 {
			fmt.Fprintf(w, "Hello, Dumbass!")
			return
		}

		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	http.ListenAndServe(":8080", nil)
}
