# go-webserver-boilerplate

A boilerplate project for building a RESTful web server in Go with structured logging, OpenTelemetry integration, and Prometheus metrics.

## Features

- **RESTful API**: Responds to HTTP requests.
- **Structured Logging**: Uses Logrus for JSON formatted logs.
- **OpenTelemetry**: Provides observability with tracing and metrics.
<<<<<<< Updated upstream
- **Prometheus Metrics**: Exposes application metrics for monitoring. (port 8081/metrics)
- **Configuration**: Utilizes a YAML configuration file with JSON tags in Go structs.

=======
- **Prometheus Metrics**: Exposes application metrics for monitoring (available at `/metrics` on port 8081).
- **Configuration**: Utilizes a YAML configuration file with JSON tags in Go structs.
- **Swagger Documentation**: API documentation is available at the `/swagger` route.
>>>>>>> Stashed changes

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/RaySheikh/go-webserver-boilerplate.git
cd go-webserver-boilerplate
make build
make run
