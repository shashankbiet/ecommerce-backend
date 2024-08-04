package metric

import "github.com/prometheus/client_golang/prometheus"

var RequestCounter *prometheus.CounterVec
var ErrorCounter *prometheus.CounterVec
var RequestLatency *prometheus.HistogramVec
var TimeoutCounter *prometheus.CounterVec

func InitMetricStore(prometheusServer *PrometheusServer) {
	RequestCounter = prometheusServer.createCounterVec("total_request_count", "received request", []string{"layer_name", "method_name"})
	ErrorCounter = prometheusServer.createCounterVec("total_error_count", "received request error", []string{"layer_name", "method_name", "error_code"})
	RequestLatency = prometheusServer.createHistogramVec("latency_histogram", "received request latency", []string{"client_name", "layer_name", "method_name"})
	TimeoutCounter = prometheusServer.createCounterVec("total_timeout_count", "number of time client timed out by different api", []string{"api_name"})
}
