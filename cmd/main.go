package main

import (
	"log"
	"net/http"

	"github.com/jrh3k5/freecaster/api"
)

func main() {
	http.Handle("/api/index", http.HandlerFunc(api.Handler))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to run HTTP server: %v", err)
	}
}
