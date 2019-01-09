package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonStruct struct {
	Test string
}

func parsePost(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var j jsonStruct
	err := decoder.Decode(&j)

	if err != nil {
		panic(err)
	}

	fmt.Println(j.Test)
}

func main() {
	// Test with: curl -X POST -d "{"jsonTestKey": "jsonTestValue"}" http://localhost:8080
	http.HandleFunc("/", parsePost)
	http.ListenAndServe(":8080", nil)
}
