package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! \n")
		fmt.Println("Console output...")
	})
	http.ListenAndServe(":8000", nil)
}
