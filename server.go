package main

import (
	"log"
	"net/http"
	"os"
)

type Hook struct{}

func (h *Hook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Request received")
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func main() {
	var h Hook
	log.Printf("Listening on %s", os.Args[1])
	err := http.ListenAndServe(os.Args[1], &h)
	if err != http.ErrServerClosed {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
