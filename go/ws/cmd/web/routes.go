package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/wolftsao/learning_notes/go/internal/handlers"
)

// routes defines the application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
