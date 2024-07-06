package controllers

import (
	"basic-api/utils"
	"fmt"
	"net/http"
)

func IpController(w http.ResponseWriter, r *http.Request) {
	var ip string = utils.ReadUserIP(r)

	fmt.Fprintf(w, "Your IP is: %s", ip)
}

func NameController(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path[1:] == "" || len(r.URL.Path[1:]) < 3 {
		fmt.Fprintf(w, "Please enter a name")
		return
	}

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
