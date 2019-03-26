package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	log "github.com/sirupsen/logrus"
)

var scrapeDurationHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: metricName("scrape_duration_seconds"),
	Help: "Histogram for the times taken to scrape the Metrics API",
})

// Scraper will scrape the metrics API
type Scraper struct {
	MetricsAPIURL string
	Caller        APICaller
}

// NewScraper creates a basic Scraper for a metrics API
func NewScraper(api string, caller APICaller) *Scraper {
	return &Scraper{
		MetricsAPIURL: api,
		Caller:        caller,
	}
}

// Scrape calls the metrics API parses the result
func (s *Scraper) Scrape() (*RawMetrics, error) {
	start := time.Now()
	defer timeTrack(start, "Metrics scrape")
	defer observeHistogramTiming(start, scrapeDurationHistogram)

	bytes, err := s.Caller.GetBytes(s.MetricsAPIURL)
	if err != nil {
		return &RawMetrics{}, err
	}

	metrics, err := NewRawMetrics(bytes)
	if err != nil {
		return &RawMetrics{}, err
	}

	log.WithFields(log.Fields{
		"rawMetrics": fmt.Sprintf("%+v", metrics),
	}).Debug("Successful scrape")

	return metrics, nil
}
