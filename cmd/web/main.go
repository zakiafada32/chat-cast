package main

import (
	"log"
	"net/http"

	"github.com/zakiafada32/chat-cast/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("starting web server on port 8080")

	_ = http.ListenAndServe("127.0.0.1:8080", mux)

}
