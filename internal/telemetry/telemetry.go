package telemetry

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0" // Ensure the correct version is imported
	"go.opentelemetry.io/otel/trace"
)

// SetupTracer initializes the OpenTelemetry tracer
func SetupTracer(serviceName, collectorURL string, sampleRatio float64) (*sdktrace.TracerProvider, trace.Tracer, error) {
	// If collectorURL is empty, skip telemetry setup and return a no-op tracer
	if collectorURL == "" {
		fmt.Println("Collector URL is empty, disabling telemetry.")
		return nil, otel.Tracer(serviceName), nil
	}

	// Attempt to connect to the OpenTelemetry Collector
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(collectorURL))
	if err != nil {
		log.Printf("Failed to create OTLP exporter: %v. Running without telemetry.", err)
		return nil, otel.Tracer(serviceName), nil // Return a no-op tracer on failure
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(sampleRatio)),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	// Set the TracerProvider globally
	otel.SetTracerProvider(tp)

	fmt.Println("Telemetry setup complete. Tracing enabled.")
	return tp, otel.Tracer(serviceName), nil
}

func ShutdownTracer(tp *sdktrace.TracerProvider) {
	if tp == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := tp.Shutdown(ctx); err != nil {
		log.Printf("Failed to shut down TracerProvider: %v", err)
	}
}
