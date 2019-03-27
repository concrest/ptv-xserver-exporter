package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// PTVMetricsCollector implements the `prometheus.Collector` interface
type PTVMetricsCollector struct {
	scraper *Scraper
	metrics *PTVMetrics
}

// NewPTVMetricsCollector creates a new custom Prometheus collector for PTV metrics
func NewPTVMetricsCollector(s *Scraper) *PTVMetricsCollector {
	log.Debug("Creating New PTVMetricsCollector")

	return &PTVMetricsCollector{
		scraper: s,
		metrics: NewPtvMetrics(),
	}
}

// Describe implements the `prometheus.Collector` interface
func (c *PTVMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	defer timeTrack(time.Now(), "PTVMetricsCollector.Describe")

	c.metrics.Describe(ch)
}

// Collect implements the `prometheus.Collector` interface
func (c *PTVMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	defer timeTrack(time.Now(), "PTVMetricsCollector.Collect")

	// External call to PTV here:
	rawMetrics, err := c.scraper.Scrape()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Scrape error")

		c.metrics.SetPtvDown(ch)
		return
	}

	c.metrics.SetPtvUp(ch, rawMetrics)
}
