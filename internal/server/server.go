package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

// Server is a simple HTTP server struct
type Server struct {
	Port   int
	Logger *logrus.Logger
	Tracer trace.Tracer
}

// NewServer initializes a new server instance
func NewServer(port int, logger *logrus.Logger, tracer trace.Tracer) *Server {
	return &Server{
		Port:   port,
		Logger: logger,
		Tracer: tracer,
	}
}

// Start the HTTP server
func (s *Server) Start() error {
	http.HandleFunc("/", s.HelloHandler)
	addr := fmt.Sprintf(":%d", s.Port)
	s.Logger.Infof("Starting server on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// HelloHandler handles the root route
func (s *Server) HelloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.Tracer.Start(r.Context(), "HelloHandler")
	defer span.End()

	s.Logger.WithContext(ctx).Info("Received request")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, Go Server!")
}
