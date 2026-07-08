package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(cfg Config, logger *slog.Logger) http.Handler {
	if logger == nil {
		logger = slog.Default()
	}

	mux := http.NewServeMux()
	api := API{
		config: cfg,
		logger: logger,
	}

	mux.HandleFunc("GET /healthz", api.health)
	mux.HandleFunc("GET /v1/fizzbuzz", api.fizzbuzz)
	mux.Handle("GET /metrics", promhttp.Handler())

	return requestMetrics(requestLogger(logger, mux))
}
