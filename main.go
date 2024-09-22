package main

import (
	"go-webserver-boilerplate/config"
	"go-webserver-boilerplate/internal/logger"
	"go-webserver-boilerplate/internal/server"
	"go-webserver-boilerplate/internal/telemetry"
	"sync"
)

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

	// Use WaitGroup to start servers
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		server.StartWebServer(cfg.Server.Port)
	}()

	go func() {
		defer wg.Done()
		telemetry.StartMetricsServer(cfg.Otel.MetricsPort)
	}()

	wg.Wait()
}
