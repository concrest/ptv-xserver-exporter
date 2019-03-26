package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

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
	defer timeTrack(time.Now(), "NewScraper.Scrape")

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
