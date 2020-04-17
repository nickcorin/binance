package binance

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(httpRequestLatencyHist)
	prometheus.MustRegister(httpResponseCodesCounter)
	prometheus.MustRegister(clientErrorsCounter)
}

var (
	httpRequestLatencyHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "binance",
		Subsystem: "http",
		Name:      "request_latency",
		Help:      "HTTP request latency in seconds",
	}, []string{"path"})

	httpResponseCodesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "binance",
		Subsystem: "http",
		Name:      "response_codes_count",
		Help:      "HTTP response code counter",
	}, []string{"path", "code"})

	clientErrorsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "binance",
		Subsystem: "client",
		Name:      "errors_count",
		Help:      "Client errors counter",
	}, []string{"error"})
)
