package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wrestler094/httpserver"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	srv := httpserver.NewServer(
		httpserver.Port(":8080"),
		httpserver.Handler(r),
	)

	if err := srv.Start(); err != nil {
		log.Fatal("failed to start server:", err)
	}

	select {
	case err := <-srv.Notify():
		log.Fatal("server error:", err)
	}
}
