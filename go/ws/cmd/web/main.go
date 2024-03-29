package main

import (
	"log"
	"net/http"

	"github.com/wolftsao/learning_notes/go/ws/internal/handlers"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Staring web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
