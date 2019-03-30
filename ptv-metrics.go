package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PTVMetrics models all the exported PTV metrics
type PTVMetrics struct {
	PTVUp       *PTVMetric
	PTVInfo     *PTVMetric
	MinPoolSize *PTVMetric
	MaxPoolSize *PTVMetric

	NumSuccess  *PTVMetric
	NumFailure  *PTVMetric
	NumRejected *PTVMetric
	NumService  *PTVMetric

	CommittedVirtualMemoryBytes *PTVMetric
	HeapCommittedMemoryBytes    *PTVMetric
	HeapUsedMemoryBytes         *PTVMetric
	NonHeapCommittedMemoryBytes *PTVMetric
	NonHeapUsedMemoryBytes      *PTVMetric

	ProcessCPUTimeSeconds *PTVMetric

	AvgInnerTimeSeconds       *PTVMetric
	MinInnerTimeSeconds       *PTVMetric
	MaxInnerTimeSeconds       *PTVMetric
	AvgComputationTimeSeconds *PTVMetric
	AvgOuterTimeSeconds       *PTVMetric
	MinOuterTimeSeconds       *PTVMetric
	MaxOuterTimeSeconds       *PTVMetric

	TimeQuantile50InnerSeconds *PTVMetric
	TimeQuantile50OuterSeconds *PTVMetric
	TimeQuantile75InnerSeconds *PTVMetric
	TimeQuantile75OuterSeconds *PTVMetric
	TimeQuantile90InnerSeconds *PTVMetric
	TimeQuantile90OuterSeconds *PTVMetric
	TimeQuantile95InnerSeconds *PTVMetric
	TimeQuantile95OuterSeconds *PTVMetric
	TimeQuantile98InnerSeconds *PTVMetric
	TimeQuantile98OuterSeconds *PTVMetric

	InstanceRestartCounter     *PTVMetric
	InstanceUserRestartCounter *PTVMetric
	InstanceUseCounter         *PTVMetric

	InstanceUptimeSeconds         *PTVMetric
	InstanceProcessCPUTimeSeconds *PTVMetric

	InstanceCommittedVirtualMemoryBytes *PTVMetric
	InstanceHeapCommittedMemoryBytes    *PTVMetric
	InstanceHeapUsedMemoryBytes         *PTVMetric
	InstanceNonHeapCommittedMemoryBytes *PTVMetric
	InstanceNonHeapUsedMemoryBytes      *PTVMetric
}

// NewPtvMetrics creates the container of metrics to export
func NewPtvMetrics() *PTVMetrics {
	return &PTVMetrics{
		PTVUp:   NewPtvMetric("up", "1 if the last metrics scrape was successful"),
		PTVInfo: NewPtvMetric("info", "A gauge with constant value 1 showing PTV info in labels", "service_name", "process_name"),

		MinPoolSize: NewPtvMetric("min_pool_size", "Value of minPoolSize"),
		MaxPoolSize: NewPtvMetric("max_pool_size", "Value of maxPoolSize"),

		NumSuccess:  NewPtvMetric("num_success", "Value of numSuccess"),
		NumFailure:  NewPtvMetric("num_failure", "Value of numFailure"),
		NumRejected: NewPtvMetric("num_rejected", "Value of numRejected"),
		NumService:  NewPtvMetric("num_service", "Value of numService"),

		CommittedVirtualMemoryBytes: NewPtvMetric("committed_virtual_memory_bytes", "Value of committedVirtualMemorySize in bytes"),
		HeapCommittedMemoryBytes:    NewPtvMetric("heap_committed_memory_bytes", "Value of heapCommittedMemorySize in bytes"),
		HeapUsedMemoryBytes:         NewPtvMetric("heap_used_memory_bytes", "Value of heapUsedMemorySize in bytes"),
		NonHeapCommittedMemoryBytes: NewPtvMetric("non_heap_committed_memory_bytes", "Value of nonHeapCommittedMemorySize in bytes"),
		NonHeapUsedMemoryBytes:      NewPtvMetric("non_heap_used_memory_bytes", "Value of nonHeapUsedMemorySize in bytes"),

		ProcessCPUTimeSeconds: NewPtvMetric("process_cpu_time_seconds", "Value of processCpuTime in seconds"),

		AvgInnerTimeSeconds:       NewPtvMetric("avg_inner_time_seconds", "Value of avgInnerTime in seconds"),
		MinInnerTimeSeconds:       NewPtvMetric("min_inner_time_seconds", "Value of minInnerTime in seconds"),
		MaxInnerTimeSeconds:       NewPtvMetric("max_inner_time_seconds", "Value of maxInnerTime in seconds"),
		AvgComputationTimeSeconds: NewPtvMetric("avg_computation_time_seconds", "Value of avgComputationTime in seconds"),
		AvgOuterTimeSeconds:       NewPtvMetric("avg_outer_time_seconds", "Value of avgOuterTime in seconds"),
		MinOuterTimeSeconds:       NewPtvMetric("min_outer_time_seconds", "Value of minOuterTime in seconds"),
		MaxOuterTimeSeconds:       NewPtvMetric("max_outer_time_seconds", "Value of maxOuterTime in seconds"),

		TimeQuantile50InnerSeconds: NewPtvMetric("time_quantile_50_inner_seconds", "Value of innerTime in timeQuantiles for q=0.5 in seconds"),
		TimeQuantile50OuterSeconds: NewPtvMetric("time_quantile_50_outer_seconds", "Value of outerTime in timeQuantiles for q=0.5 in seconds"),
		TimeQuantile75InnerSeconds: NewPtvMetric("time_quantile_75_inner_seconds", "Value of innerTime in timeQuantiles for q=0.75 in seconds"),
		TimeQuantile75OuterSeconds: NewPtvMetric("time_quantile_75_outer_seconds", "Value of outerTime in timeQuantiles for q=0.75 in seconds"),
		TimeQuantile90InnerSeconds: NewPtvMetric("time_quantile_90_inner_seconds", "Value of innerTime in timeQuantiles for q=0.90 in seconds"),
		TimeQuantile90OuterSeconds: NewPtvMetric("time_quantile_90_outer_seconds", "Value of outerTime in timeQuantiles for q=0.90 in seconds"),
		TimeQuantile95InnerSeconds: NewPtvMetric("time_quantile_95_inner_seconds", "Value of innerTime in timeQuantiles for q=0.95 in seconds"),
		TimeQuantile95OuterSeconds: NewPtvMetric("time_quantile_95_outer_seconds", "Value of outerTime in timeQuantiles for q=0.95 in seconds"),
		TimeQuantile98InnerSeconds: NewPtvMetric("time_quantile_98_inner_seconds", "Value of innerTime in timeQuantiles for q=0.98 in seconds"),
		TimeQuantile98OuterSeconds: NewPtvMetric("time_quantile_98_outer_seconds", "Value of outerTime in timeQuantiles for q=0.98 in seconds"),

		InstanceRestartCounter:     NewPtvMetric("instance_restart_counter", "Value of restartCounter for an instance", "instance_name"),
		InstanceUserRestartCounter: NewPtvMetric("instance_user_restart_counter", "Value of userRestartCounter for an instance", "instance_name"),
		InstanceUseCounter:         NewPtvMetric("instance_use_counter", "Value of useCounter for an instance", "instance_name"),

		InstanceUptimeSeconds:         NewPtvMetric("instance_uptime_seconds", "Value of uptime in seconds for an instance", "instance_name"),
		InstanceProcessCPUTimeSeconds: NewPtvMetric("instance_process_cpu_time_seconds", "Value of processCpuTime in seconds for an instance", "instance_name"),

		InstanceCommittedVirtualMemoryBytes: NewPtvMetric("instance_committed_virtual_memory_bytes", "Value of committedVirtualMemorySize for an instance", "instance_name"),
		InstanceHeapCommittedMemoryBytes:    NewPtvMetric("instance_heap_committed_memory_bytes", "Value of heapCommittedMemorySize for an instance", "instance_name"),
		InstanceHeapUsedMemoryBytes:         NewPtvMetric("instance_heap_used_memory_bytes", "Value of heapUsedMemorySize for an instance", "instance_name"),
		InstanceNonHeapCommittedMemoryBytes: NewPtvMetric("instance_non_heap_committed_memory_bytes", "Value of nonHeapCommittedMemorySize for an instance", "instance_name"),
		InstanceNonHeapUsedMemoryBytes:      NewPtvMetric("instance_non_heap_used_memory_bytes", "Value of nonHeapUsedMemorySize for an instance", "instance_name"),
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
	ch <- m.CommittedVirtualMemoryBytes.Desc
	ch <- m.HeapCommittedMemoryBytes.Desc
	ch <- m.HeapUsedMemoryBytes.Desc
	ch <- m.NonHeapCommittedMemoryBytes.Desc
	ch <- m.NonHeapUsedMemoryBytes.Desc
	ch <- m.ProcessCPUTimeSeconds.Desc

	ch <- m.AvgInnerTimeSeconds.Desc
	ch <- m.MinInnerTimeSeconds.Desc
	ch <- m.MaxInnerTimeSeconds.Desc
	ch <- m.AvgComputationTimeSeconds.Desc
	ch <- m.AvgOuterTimeSeconds.Desc
	ch <- m.MinOuterTimeSeconds.Desc
	ch <- m.MaxOuterTimeSeconds.Desc

	ch <- m.TimeQuantile50InnerSeconds.Desc
	ch <- m.TimeQuantile50OuterSeconds.Desc
	ch <- m.TimeQuantile75InnerSeconds.Desc
	ch <- m.TimeQuantile75OuterSeconds.Desc
	ch <- m.TimeQuantile90InnerSeconds.Desc
	ch <- m.TimeQuantile90OuterSeconds.Desc
	ch <- m.TimeQuantile95InnerSeconds.Desc
	ch <- m.TimeQuantile95OuterSeconds.Desc
	ch <- m.TimeQuantile98InnerSeconds.Desc
	ch <- m.TimeQuantile98OuterSeconds.Desc

	ch <- m.InstanceRestartCounter.Desc
	ch <- m.InstanceUserRestartCounter.Desc
	ch <- m.InstanceUseCounter.Desc

	ch <- m.InstanceUptimeSeconds.Desc
	ch <- m.InstanceProcessCPUTimeSeconds.Desc

	ch <- m.InstanceCommittedVirtualMemoryBytes.Desc
	ch <- m.InstanceHeapCommittedMemoryBytes.Desc
	ch <- m.InstanceHeapUsedMemoryBytes.Desc
	ch <- m.InstanceNonHeapCommittedMemoryBytes.Desc
	ch <- m.InstanceNonHeapUsedMemoryBytes.Desc
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

	ch <- m.CommittedVirtualMemoryBytes.CreateMetric(prometheus.GaugeValue, rawMetrics.CommittedVirtualMemorySize)
	ch <- m.HeapCommittedMemoryBytes.CreateMetric(prometheus.GaugeValue, rawMetrics.HeapCommittedMemorySize)
	ch <- m.HeapUsedMemoryBytes.CreateMetric(prometheus.GaugeValue, rawMetrics.HeapUsedMemorySize)
	ch <- m.NonHeapCommittedMemoryBytes.CreateMetric(prometheus.GaugeValue, rawMetrics.NonHeapCommittedMemorySize)

	ch <- m.ProcessCPUTimeSeconds.CreateMetric(prometheus.CounterValue, cpuTimeToSeconds(rawMetrics.ProcessCPUTime))

	ch <- m.AvgInnerTimeSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.AvgInnerTime))
	ch <- m.MinInnerTimeSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.MinInnerTime))
	ch <- m.MaxInnerTimeSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.MaxInnerTime))
	ch <- m.AvgComputationTimeSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.AvgComputationTime))
	ch <- m.AvgOuterTimeSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.AvgOuterTime))
	ch <- m.MinOuterTimeSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.MinOuterTime))
	ch <- m.MaxOuterTimeSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.MaxOuterTime))

	ch <- m.TimeQuantile50InnerSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.5).InnerTime))
	ch <- m.TimeQuantile50OuterSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.5).OuterTime))
	ch <- m.TimeQuantile75InnerSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.75).InnerTime))
	ch <- m.TimeQuantile75OuterSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.75).OuterTime))
	ch <- m.TimeQuantile90InnerSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.9).InnerTime))
	ch <- m.TimeQuantile90OuterSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.9).OuterTime))
	ch <- m.TimeQuantile95InnerSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.95).InnerTime))
	ch <- m.TimeQuantile95OuterSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.95).OuterTime))
	ch <- m.TimeQuantile98InnerSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.98).InnerTime))
	ch <- m.TimeQuantile98OuterSeconds.CreateMetric(prometheus.GaugeValue, millisToSeconds(rawMetrics.GetQuantile(0.98).OuterTime))

	for _, rawInstance := range rawMetrics.Instances {
		ch <- m.InstanceRestartCounter.CreateMetric(prometheus.CounterValue, rawInstance.RestartCounter, rawInstance.InstanceSuffix)
		ch <- m.InstanceUserRestartCounter.CreateMetric(prometheus.CounterValue, rawInstance.UserRestartCounter, rawInstance.InstanceSuffix)
		ch <- m.InstanceUseCounter.CreateMetric(prometheus.CounterValue, rawInstance.UseCounter, rawInstance.InstanceSuffix)

		// "upTime" looks like it's in milliseconds, while cpuTime is in nanoseconds?
		ch <- m.InstanceUptimeSeconds.CreateMetric(prometheus.CounterValue, millisToSeconds(rawInstance.Uptime), rawInstance.InstanceSuffix)
		ch <- m.InstanceProcessCPUTimeSeconds.CreateMetric(prometheus.CounterValue, cpuTimeToSeconds(rawInstance.ProcessCPUTime), rawInstance.InstanceSuffix)

		ch <- m.InstanceCommittedVirtualMemoryBytes.CreateMetric(prometheus.GaugeValue, rawInstance.CommittedVirtualMemorySize, rawInstance.InstanceSuffix)
		ch <- m.InstanceHeapCommittedMemoryBytes.CreateMetric(prometheus.GaugeValue, rawInstance.HeapCommittedMemorySize, rawInstance.InstanceSuffix)
		ch <- m.InstanceHeapUsedMemoryBytes.CreateMetric(prometheus.GaugeValue, rawInstance.HeapUsedMemorySize, rawInstance.InstanceSuffix)
		ch <- m.InstanceNonHeapCommittedMemoryBytes.CreateMetric(prometheus.GaugeValue, rawInstance.NonHeapCommittedMemorySize, rawInstance.InstanceSuffix)
		ch <- m.InstanceNonHeapUsedMemoryBytes.CreateMetric(prometheus.GaugeValue, rawInstance.NonHeapUsedMemorySize, rawInstance.InstanceSuffix)
	}
}

// cpuTimeToSeconds converts PTV processor time (appears to be nanoseconds) to seconds
func cpuTimeToSeconds(fromPtv float64) float64 {
	return fromPtv / 1000000000
}

// millisToSeconds converts non-zero `millis` values to seconds
func millisToSeconds(millis float64) float64 {
	if millis > 0 {
		return millis / 1000.0
	}

	return 0
}
