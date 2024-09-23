package main

import (
	"os"
	"web-starter/cmd/lib"
	"web-starter/handlers"

	// "handlers"
	"log"
	"net/http"
	"time"
)

var ntfy_token string

func init() {
	// load env file
	lib.LoadEnv(".env")
	ntfy_token = os.Getenv("NTFY_token")
}

func main() {
	mux := setupMux()

	server := setupServer(mux)
	log.Printf("Server starting on http://%s...\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func setupMux() *http.ServeMux {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Set up routes
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/error", handlers.ForceDirectError) // !for testing purpose only (not for production)
	mux.HandleFunc("/500", handlers.Force500Handler)    // !for testing purpose only (not for production)

	return mux
}

func setupServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              "localhost:8080",
		Handler:           handlers.WithErrorHandling(handler),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
		// ErrorLog: *log.Logger,
	}
}
