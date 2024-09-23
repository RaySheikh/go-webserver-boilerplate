package server

import (
	"go-webserver-boilerplate/internal/logger"
	"go-webserver-boilerplate/internal/telemetry"
	"net/http"

	"github.com/sirupsen/logrus"
)

// StartWebServer initializes the web server
func StartWebServer(port int, mux *http.ServeMux) {
	// Handle the root path
	mux.Handle("/", telemetry.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.String(),
		}).Info("Received request for /")
		w.Write([]byte("Hello, World!"))
	})))

	// Handle GET /user/{id}
	mux.Handle("/user/", telemetry.HTTPHandler(http.HandlerFunc(GetUserHandler)))

	logger.Logger.Infof("Starting web server on port %d\n", port)
}
