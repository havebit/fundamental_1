package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var port = ":8081"

func main() {
	defer func() {
		log.Println("ok")
	}()

	http.HandleFunc("/hello", helloHandler)
	log.Println("listening on", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Panic(err)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello world!!",
	})
}
