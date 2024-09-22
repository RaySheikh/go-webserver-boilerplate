package server

import (
	"go-webserver-boilerplate/internal/logger"
	"go-webserver-boilerplate/internal/telemetry"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func StartWebServer(port int) {
	mux := http.NewServeMux()
	mux.Handle("/", telemetry.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.String(),
		}).Info("Received request for /")
		w.Write([]byte("Hello, World!"))
	})))

	logger.Logger.Infof("Starting web server on port %d\n", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), mux); err != nil {
		logger.Logger.Fatalf("Could not start web server: %v", err)
	}
}
