package main

import (
	"context"
	"go-webserver-boilerplate/config"
	"go-webserver-boilerplate/internal/logger"
	"go-webserver-boilerplate/internal/server"
	"go-webserver-boilerplate/internal/telemetry"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Setup logging
	logger := logger.SetupLogger()

	// Setup OpenTelemetry tracer
	tracerProvider, tracer, err := telemetry.SetupTracer(cfg.OTel.ServiceName, cfg.OTel.CollectorURL, cfg.OTel.SampleRatio)
	if err != nil {
		logger.Fatalf("Error setting up OpenTelemetry tracer: %v", err)
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			logger.Fatalf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Initialize and start the server
	srv := server.NewServer(cfg.Server.Port, logger, tracer)
	if err := srv.Start(); err != nil {
		logger.Fatalf("Error starting server: %v", err)
	}
}
