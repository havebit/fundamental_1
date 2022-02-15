package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"hello/foobar"

	"github.com/gorilla/mux"
)

var port = ":8081"

func main() {
	defer func() {
		log.Println("ok")
	}()

	r := mux.NewRouter()

	r.HandleFunc("/foobar/{param}", foobarHandler).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: r,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("listening on", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func foobarHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(map[string]string{
		"foobar": foobar.SayAny(vars["param"]),
	})
}
