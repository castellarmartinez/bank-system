package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Account Service is running!")
	})

	fmt.Println("Starting Account Service on port 8081...")
	http.ListenAndServe(":8081", nil)
}
