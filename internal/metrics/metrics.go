package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)


type Metrics struct {
	requestsTotal *prometheus.CounterVec
	responseLatency *prometheus.HistogramVec
	createdPvzs *prometheus.CounterVec
	createdReceptions *prometheus.CounterVec
	addedProducts *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	metrics := &Metrics{
		requestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name:      "requests_total",
				Help:      "Total number of requests",
			},
			[]string{"method", "path", "status"},
		),
		responseLatency: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:      "response_latency",
				Help:      "Response latency in milliseconds",
				Buckets: []float64{50, 100, 300, 500, 1_000},
			},
			[]string{"method", "path"},
		),
		createdPvzs: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name:      "created_pvzs",
				Help:      "Total number of created Pvzs",
			},
			[]string{"status"},
		),
		createdReceptions: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name:      "created_receptions",
				Help:      "Total number of created receptions",
			},
			[]string{"status"},
		),
		addedProducts: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name:      "added_products",
				Help:      "Total number of added products",
			},
			[]string{"status"},
		),
	}
	return metrics
}

func (m *Metrics) Register() {
	prometheus.MustRegister(m.requestsTotal)
	prometheus.MustRegister(m.responseLatency)
}

func (m *Metrics) IncRequestsTotal(method, path, status string) {
	m.requestsTotal.With(prometheus.Labels{
		"method": method,
		"path":   path,
		"status": status,
	}).Inc()
}


func (m *Metrics) ResponseLatency(method, path string, milliseconds float64) {
	m.responseLatency.With(prometheus.Labels{
		"method": method,
		"path":   path,
	}).Observe(milliseconds)
}


func (m *Metrics) IncCreatedPvzs(status string) {
	m.createdPvzs.With(prometheus.Labels{
		"status": status,
	}).Inc()
}

func (m *Metrics) IncCreatedReceptions(status string) {
	m.createdReceptions.With(prometheus.Labels{
		"status": status,
	}).Inc()
}	

func (m *Metrics) IncAddedProducts(status string) {
	m.addedProducts.With(prometheus.Labels{
		"status": status,
	}).Inc()
}


