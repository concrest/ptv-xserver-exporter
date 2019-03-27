package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// PTVMetricsCollector implements the `prometheus.Collector` interface
type PTVMetricsCollector struct {
	scraper         *Scraper
	ptvUpDesc       *prometheus.Desc
	minPoolSizeDesc *prometheus.Desc
}

// NewPTVMetricsCollector creates a new custom Prometheus collector for PTV metrics
func NewPTVMetricsCollector(s *Scraper) *PTVMetricsCollector {
	log.Debug("Creating New PTVMetricsCollector")

	return &PTVMetricsCollector{
		scraper:         s,
		ptvUpDesc:       prometheus.NewDesc(metricName("up"), "1 if the last metrics scrape was successful", nil, nil),
		minPoolSizeDesc: prometheus.NewDesc(metricName("min_pool_size"), "Value of minPoolSize", nil, nil),
	}
}

// Describe implements the `prometheus.Collector` interface
func (c *PTVMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	defer timeTrack(time.Now(), "PTVMetricsCollector.Describe")

	ch <- c.ptvUpDesc
	ch <- c.minPoolSizeDesc
}

// Collect implements the `prometheus.Collector` interface
func (c *PTVMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	defer timeTrack(time.Now(), "PTVMetricsCollector.Collect")

	rawMetrics, err := c.scraper.Scrape()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Scrape error")

		ch <- prometheus.MustNewConstMetric(
			c.ptvUpDesc,
			prometheus.GaugeValue,
			0,
		)
		return
	}

	// Scrape successful - map metrics to prometheus ones
	ch <- prometheus.MustNewConstMetric(
		c.ptvUpDesc,
		prometheus.GaugeValue,
		1,
	)
	ch <- prometheus.MustNewConstMetric(
		c.minPoolSizeDesc,
		prometheus.GaugeValue,
		rawMetrics.MinPoolSize,
	)
}
