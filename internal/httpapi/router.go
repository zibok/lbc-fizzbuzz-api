package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(cfg Config, logger *slog.Logger) http.Handler {
	return NewRouterWithStatisticsRecorder(cfg, logger, NewInMemoryStatisticsRecorder())
}

func NewRouterWithStatisticsRecorder(cfg Config, logger *slog.Logger, statisticsRecorder StatisticsRecorder) http.Handler {
	cfg = cfg.WithDefaults()

	if statisticsRecorder == nil {
		statisticsRecorder = NewInMemoryStatisticsRecorder()
	}

	if logger == nil {
		logger = slog.Default()
	}

	mux := http.NewServeMux()
	api := API{
		config:             cfg,
		logger:             logger,
		statisticsRecorder: statisticsRecorder,
	}

	mux.HandleFunc("GET /healthz", api.health)
	mux.HandleFunc("GET /v1/fizzbuzz", api.fizzbuzz)
	mux.HandleFunc("GET /v1/statistics", api.statistics)
	mux.Handle("GET /metrics", promhttp.Handler())

	return requestMetrics(requestLogger(logger, mux))
}
