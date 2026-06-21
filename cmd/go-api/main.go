package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tanzid64/go-api/internal/config"
	"github.com/tanzid64/go-api/internal/http/handlers/student"
)

func main() {
	// Load Configuration
	cfg := config.MustLoad()

	// Database setup
	// Router setup
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())
	// Start server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	log.Printf("Server started at %s", cfg.Addr)
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err.Error())
		}
	}()

	<-done

	slog.Info("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Error occurred while shutting down server", "error", err.Error())
	}

	slog.Info("Server stopped gracefully")
}
