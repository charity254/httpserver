package main

import (
	
	"fmt"
	"net/http"
	
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/health", getHealth)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/status", getStatus)

	fmt.Println("Server starting on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
