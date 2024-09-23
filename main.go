package main

import (
	"go-webserver-boilerplate/config"
	_ "go-webserver-boilerplate/docs" // Import the generated docs
	"go-webserver-boilerplate/internal/logger"
	"go-webserver-boilerplate/internal/server"
	"go-webserver-boilerplate/internal/telemetry"
	"net/http"
	"strconv"
	"sync"

	httpSwagger "github.com/swaggo/http-swagger" // Swagger middleware
)

// @title Go Webserver API
// @version 1.0
// @description This is a sample server for a Go web server boilerplate.
// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Logger.Fatalf("Could not load config: %v", err)
	}

	// Initialize logger
	logger.InitLogger()

	// Initialize telemetry
	telemetry.InitTelemetry(cfg.Otel.ServiceName)

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve Swagger UI
	mux.Handle("/swagger/", httpSwagger.WrapHandler)         // Serve Swagger UI
	mux.Handle("/swagger/doc.json", httpSwagger.WrapHandler) // Serve the Swagger documentation

	// Use WaitGroup to start servers
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		server.StartWebServer(cfg.Server.Port, mux) // Pass the mux to the web server
	}()

	go func() {
		defer wg.Done()
		telemetry.StartMetricsServer(cfg.Otel.MetricsPort)
	}()

	// Start the HTTP server with the mux that includes Swagger
	logger.Logger.Infof("Starting server on port %d\n", cfg.Server.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), mux); err != nil {
		logger.Logger.Fatalf("Could not start server: %v", err)
	}

	wg.Wait()
}
