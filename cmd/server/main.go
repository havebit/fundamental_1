package main

import (
	"encoding/json"
	"hello/foobar"
	"log"
	"net/http"
	"strings"
)

var port = ":8081"

func main() {
	defer func() {
		log.Println("ok")
	}()

	http.HandleFunc("/foobar/", foobarHandler)
	log.Println("listening on", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Panic(err)
	}
}

func foobarHandler(w http.ResponseWriter, req *http.Request) {
	param := strings.TrimPrefix(req.RequestURI, "/foobar/")

	json.NewEncoder(w).Encode(map[string]string{
		"foobar": foobar.SayAny(param),
	})
}
