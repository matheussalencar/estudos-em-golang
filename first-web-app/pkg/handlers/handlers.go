package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

func About(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "this is the about page")
}
