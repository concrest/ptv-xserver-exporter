package main

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// InstanceMetrics models 1 of RawMetrics Instances collection
type InstanceMetrics struct {
	InstanceSuffix             string
	RestartCounter             float64
	UserRestartCounter         float64
	Uptime                     float64
	InUse                      bool
	UseCounter                 float64
	ModuleStatus               string
	ProcessName                string
	CommittedVirtualMemorySize float64
	ProcessCPUTime             float64
	HeapCommittedMemorySize    float64
	HeapUsedMemorySize         float64
	NonHeapCommittedMemorySize float64
	NonHeapUsedMemorySize      float64
}

// RawMetrics models the result of the metrics request to the PTV server
type RawMetrics struct {
	ServiceName                string
	MinPoolSize                float64 // TODO: Make all numbers compatible with Prometheus types
	MaxPoolSize                float64
	NumSuccess                 float64
	NumFailure                 float64
	NumRejected                float64
	NumService                 float64
	AvgInnerTime               float64
	MinInnerTime               float64
	MaxInnerTime               float64
	AvgComputationTime         float64
	AvgOuterTime               float64
	MinOuterTime               float64
	MaxOuterTime               float64
	ProcessName                string
	CommittedVirtualMemorySize float64
	ProcessCPUTime             float64
	HeapCommittedMemorySize    float64
	HeapUsedMemorySize         float64
	NonHeapCommittedMemorySize float64
	NonHeapUsedMemorySize      float64
	Instances                  []InstanceMetrics
}

// NewRawMetrics creates a RawMetrics struct from some JSON bytes
func NewRawMetrics(bytes []byte) (*RawMetrics, error) {
	metrics := RawMetrics{}

	err := json.Unmarshal(bytes, &metrics)
	if err != nil {
		log.WithFields(log.Fields{
			"bytes": fmt.Sprintf("%s", bytes),
			"err":   err,
		}).Warn("Error parsing RawMetrics bytes")

		return &RawMetrics{}, err
	}

	return &metrics, nil
}
