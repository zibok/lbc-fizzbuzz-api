package httpapi

import "github.com/prometheus/client_golang/prometheus"

var httpRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "HTTP request response time in seconds.",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"route", "status_code"},
)

func init() {
	prometheus.MustRegister(httpRequestDuration)
}
