package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PTVMetrics models all the exported PTV metrics
type PTVMetrics struct {
	PTVUp                      *PTVMetric
	PTVInfo                    *PTVMetric
	MinPoolSize                *PTVMetric
	MaxPoolSize                *PTVMetric
	NumSuccess                 *PTVMetric
	NumFailure                 *PTVMetric
	NumRejected                *PTVMetric
	NumService                 *PTVMetric
	CommittedVirtualMemorySize *PTVMetric
	HeapCommittedMemorySize    *PTVMetric
	HeapUsedMemorySize         *PTVMetric
	NonHeapCommittedMemorySize *PTVMetric
	NonHeapUsedMemorySize      *PTVMetric
}

// NewPtvMetrics creates the container of metrics to export
func NewPtvMetrics() *PTVMetrics {
	return &PTVMetrics{
		PTVUp:                      NewPtvMetric("up", "1 if the last metrics scrape was successful"),
		PTVInfo:                    NewPtvMetric("info", "A gauge with constant value 1 showing PTV info in labels", "service_name", "process_name"),
		MinPoolSize:                NewPtvMetric("min_pool_size", "Value of minPoolSize"),
		MaxPoolSize:                NewPtvMetric("max_pool_size", "Value of maxPoolSize"),
		NumSuccess:                 NewPtvMetric("num_success", "Value of numSuccess"),
		NumFailure:                 NewPtvMetric("num_failure", "Value of numFailure"),
		NumRejected:                NewPtvMetric("num_rejected", "Value of numRejected"),
		NumService:                 NewPtvMetric("num_service", "Value of numService"),
		CommittedVirtualMemorySize: NewPtvMetric("committed_virtual_memory_bytes", "Value of committedVirtualMemorySize"),
		HeapCommittedMemorySize:    NewPtvMetric("heap_committed_memory_bytes", "Value of heapCommittedMemorySize"),
		HeapUsedMemorySize:         NewPtvMetric("heap_used_memory_bytes", "Value of heapUsedMemorySize"),
		NonHeapCommittedMemorySize: NewPtvMetric("non_heap_committed_memory_bytes", "Value of nonHeapCommittedMemorySize"),
		NonHeapUsedMemorySize:      NewPtvMetric("non_heap_used_memory_bytes", "Value of nonHeapUsedMemorySize"),
	}
}

// Describe pipes all metric descriptions to the `ch` channel
func (m *PTVMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.PTVUp.Desc
	ch <- m.PTVInfo.Desc
	ch <- m.MinPoolSize.Desc
	ch <- m.MaxPoolSize.Desc
	ch <- m.NumSuccess.Desc
	ch <- m.NumFailure.Desc
	ch <- m.NumRejected.Desc
	ch <- m.NumService.Desc
	ch <- m.CommittedVirtualMemorySize.Desc
	ch <- m.HeapCommittedMemorySize.Desc
	ch <- m.HeapUsedMemorySize.Desc
	ch <- m.NonHeapCommittedMemorySize.Desc
	ch <- m.NonHeapUsedMemorySize.Desc
}

// SetPtvDown exports metrics indicating that the PTV instance is down
func (m *PTVMetrics) SetPtvDown(ch chan<- prometheus.Metric) {
	ch <- m.PTVUp.CreateMetric(prometheus.GaugeValue, 0)
}

// SetPtvUp exports PTV metrics
func (m *PTVMetrics) SetPtvUp(ch chan<- prometheus.Metric, rawMetrics *RawMetrics) {
	ch <- m.PTVUp.CreateMetric(prometheus.GaugeValue, 1)
	ch <- m.PTVInfo.CreateMetric(prometheus.GaugeValue, 1, rawMetrics.ServiceName, rawMetrics.ProcessName)
	ch <- m.MinPoolSize.CreateMetric(prometheus.GaugeValue, rawMetrics.MinPoolSize)
	ch <- m.MaxPoolSize.CreateMetric(prometheus.GaugeValue, rawMetrics.MaxPoolSize)
	ch <- m.NumSuccess.CreateMetric(prometheus.CounterValue, rawMetrics.NumSuccess)
	ch <- m.NumFailure.CreateMetric(prometheus.CounterValue, rawMetrics.NumFailure)
	ch <- m.NumRejected.CreateMetric(prometheus.CounterValue, rawMetrics.NumRejected)
	ch <- m.NumService.CreateMetric(prometheus.CounterValue, rawMetrics.NumService)

	ch <- m.CommittedVirtualMemorySize.CreateMetric(prometheus.GaugeValue, rawMetrics.CommittedVirtualMemorySize)
	ch <- m.HeapCommittedMemorySize.CreateMetric(prometheus.GaugeValue, rawMetrics.HeapCommittedMemorySize)
	ch <- m.HeapUsedMemorySize.CreateMetric(prometheus.GaugeValue, rawMetrics.HeapUsedMemorySize)
	ch <- m.NonHeapCommittedMemorySize.CreateMetric(prometheus.GaugeValue, rawMetrics.NonHeapCommittedMemorySize)
	ch <- m.NonHeapUsedMemorySize.CreateMetric(prometheus.GaugeValue, rawMetrics.NonHeapUsedMemorySize)
}
