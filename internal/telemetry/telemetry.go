package telemetry

import (
	"go-webserver-boilerplate/internal/logger"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTelemetry(serviceName string) {
	tracerProvider := trace.NewTracerProvider()
	otel.SetTracerProvider(tracerProvider)

	// Additional telemetry setup can go here (e.g., exporters)
}

func HTTPHandler(handler http.Handler) http.Handler {
	return otelhttp.NewHandler(handler, "HTTPHandler")
}

func StartMetricsServer(port int) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	logger.Logger.Infof("Starting metrics server on port %d\n", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), mux); err != nil {
		logger.Logger.Fatalf("Could not start metrics server: %v", err)
	}
}
