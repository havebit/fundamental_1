package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello world!!",
	})
}
