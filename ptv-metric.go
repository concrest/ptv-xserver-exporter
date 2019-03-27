package main

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// PTVMetric models a single PTV metric
type PTVMetric struct {
	Desc *prometheus.Desc
}

// NewPtvMetric creates 1 new metric
func NewPtvMetric(name string, help string, variableLabels ...string) *PTVMetric {

	log.WithFields(log.Fields{
		"name":           name,
		"help":           help,
		"variableLabels": variableLabels,
	}).Debug("Creating PTVMetric")

	return &PTVMetric{
		Desc: prometheus.NewDesc(metricName(name), help, variableLabels, nil),
	}
}

// CreateMetric builds a new constant metric
func (m *PTVMetric) CreateMetric(valueType prometheus.ValueType, value float64, labelValues ...string) prometheus.Metric {
	log.WithFields(log.Fields{
		"name":        m.Desc,
		"valueType":   valueType,
		"value":       value,
		"labelValues": labelValues,
	}).Debug("CreateMetric")

	return prometheus.MustNewConstMetric(m.Desc, valueType, value, labelValues...)
}
