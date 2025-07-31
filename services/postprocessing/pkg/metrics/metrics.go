package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Namespace defines the namespace for the defines metrics.
	Namespace = "opencloud"

	// Subsystem defines the subsystem for the defines metrics.
	Subsystem = "postprocessing"

	buildInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "build_info",
		Help:      "Build information",
	}, []string{"version"})
	eventsOutstandingAcks = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "events_outstanding_acks",
		Help:      "Number of outstanding acks for events",
	})
	eventsUnprocessed = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "events_unprocessed",
		Help:      "Number of unprocessed events",
	})
	eventsRedelivered = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "events_redelivered",
		Help:      "Number of redelivered events",
	})
	inProgress = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "in_progress",
		Help:      "Number of postprocessing events in progress",
	})
	finished = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "finished",
		Help:      "Number of finished postprocessing events",
	}, []string{"status"})
	duration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      "duration_seconds",
		Help:      "Duration of postprocessing operations in seconds",
		Buckets:   []float64{0.1, 0.5, 1, 2.5, 5, 10, 30, 60, 120, 300, 600, 1200},
	}, []string{"status"})
)

// Metrics defines the available metrics of this service.
type Metrics struct {
	// Counter  *prometheus.CounterVec
	BuildInfo             *prometheus.GaugeVec
	EventsOutstandingAcks prometheus.Gauge
	EventsUnprocessed     prometheus.Gauge
	EventsRedelivered     prometheus.Gauge
	InProgress            prometheus.Gauge
	Finished              *prometheus.CounterVec
	Duration              *prometheus.HistogramVec
}

// New initializes the available metrics.
func New() *Metrics {
	m := &Metrics{
		BuildInfo:             buildInfo,
		EventsOutstandingAcks: eventsOutstandingAcks,
		EventsUnprocessed:     eventsUnprocessed,
		EventsRedelivered:     eventsRedelivered,
		InProgress:            inProgress,
		Finished:              finished,
		Duration:              duration,
	}

	return m
}
