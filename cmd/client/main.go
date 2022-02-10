package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8081/foobar/5")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	var m foobarMessage
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		log.Panic(err)
	}

	fmt.Println(m)
}

type foobarMessage struct {
	FooBar string `json:"foobar"`
}
