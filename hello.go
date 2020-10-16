package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Hello returns hello string
func Hello() string {
	return "Hello, world."
}

// HelloHandler is a handler that prints hello message
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, Hello())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloHandler)
	r.Path("/metrics").Handler(promhttp.Handler())

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
