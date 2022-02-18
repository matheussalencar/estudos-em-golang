package main

import (
	"github.com/matheussalencar/estudos-em-golang/first-web-app/pkg/handlers"
	"net/http"
)

const portNumber = ":5000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(portNumber, nil)
}
