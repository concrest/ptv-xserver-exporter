package main

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// InstanceMetrics models 1 of RawMetrics Instances collection
type InstanceMetrics struct {
	InstanceSuffix             string
	RestartCounter             int
	UserRestartCounter         int
	Uptime                     uint64
	InUse                      bool
	UseCounter                 int
	ModuleStatus               string
	ProcessName                string
	CommittedVirtualMemorySize uint64
	ProcessCPUTime             uint64
	HeapCommittedMemorySize    uint64
	HeapUsedMemorySize         uint64
	NonHeapCommittedMemorySize uint64
	NonHeapUsedMemorySize      uint64
}

// RawMetrics models the result of the metrics request to the PTV server
type RawMetrics struct {
	ServiceName                string
	MinPoolSize                float64 // TODO: Make all numbers compatible with Prometheus types
	MaxPoolSize                int
	NumSuccess                 uint64
	NumFailure                 uint64
	NumRejected                uint64
	NumService                 int
	AvgInnerTime               int
	MinInnerTime               int
	MaxInnerTime               int
	AvgComputationTime         int
	AvgOuterTime               int
	MinOuterTime               int
	MaxOuterTime               int
	ProcessName                string
	CommittedVirtualMemorySize uint64
	ProcessCPUTime             uint64
	HeapCommittedMemorySize    uint64
	HeapUsedMemorySize         uint64
	NonHeapCommittedMemorySize uint64
	NonHeapUsedMemorySize      uint64
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
