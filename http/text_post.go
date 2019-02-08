package main

import (
	"fmt"
	"net/http"
)

func post(rw http.ResponseWriter, request *http.Request) {
	fmt.Println("You hit me!")
}

func main() {
	// Test with: curl -X POST -d "ping" http://localhost:8080
	http.HandleFunc("/", post)
	http.ListenAndServe(":8080", nil)
}
