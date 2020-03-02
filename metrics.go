package binance

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(httpRequestLatency)
	prometheus.MustRegister(httpResponseCodes)
}

var (
	httpRequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "binance",
		Subsystem: "http",
		Name:      "request_latency",
		Help:      "HTTP request latency in seconds",
	}, []string{"path"})

	httpResponseCodes = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "binance",
		Subsystem: "http",
		Name:      "response_codes_count",
		Help:      "HTTP response code counter",
	}, []string{"path", "code"})
)
