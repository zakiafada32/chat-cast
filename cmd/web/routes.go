package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/zakiafada32/chat-cast/internal/handlers"
)

// routes defines the application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
