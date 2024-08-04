package metric

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusServer struct {
	Namespace string
	Subsystem string
	Registry  *prometheus.Registry
}

func (ps *PrometheusServer) NewPrometheusServer(namespace string, subsystem string) {
	ps.Namespace = namespace
	ps.Subsystem = subsystem
	ps.Registry = prometheus.NewRegistry()
}

func (ps *PrometheusServer) CreateGrpcServerMetrics() *grpc_prometheus.ServerMetrics {
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	ps.Registry.MustRegister(grpcMetrics)
	return grpcMetrics
}

// Namespace, Subsystem, and Name are components of the fully-qualified  name of the Metric (created by joining these components with "_"). '
// Only Name is mandatory, the others merely help structuring the name. Note that the fully-qualified name of the metric must be a valid Prometheus metric name.
func (ps *PrometheusServer) createCounterVec(name string, help string, labels []string) *prometheus.CounterVec {
	metrics := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: ps.Namespace,
		Subsystem: ps.Subsystem,
		Name:      name,
		Help:      help,
	}, labels)
	ps.Registry.MustRegister(metrics)
	return metrics
}

func (ps *PrometheusServer) createHistogramVec(name string, help string, labels []string) *prometheus.HistogramVec {
	metrics := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: ps.Namespace,
		Subsystem: ps.Subsystem,
		Name:      name,
		Help:      help,
		Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
	}, labels)
	ps.Registry.MustRegister(metrics)
	return metrics
}
