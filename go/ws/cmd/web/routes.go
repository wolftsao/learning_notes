package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/wolftsao/learning_notes/go/ws/internal/handlers"
)

// routes defines the application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
