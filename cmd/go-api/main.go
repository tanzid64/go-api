package main

import (
	"log"
	"net/http"

	"github.com/tanzid64/go-api/internal/config"
)

func main() {
	// Load Configuration
	cfg := config.MustLoad()

	// Database setup
	// Router setup
	router := http.NewServeMux()
	router.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go API"))
	})
	// Start server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}

	log.Printf("Server started at %s", cfg.Addr)
}
