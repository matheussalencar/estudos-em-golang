package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		n, err := fmt.Fprintf(writer, "Hello world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(":5000", nil)
}
