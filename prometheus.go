package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var buildInfoGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricName("build_info"),
	Help: "A gauge with constant value 1 showing build info in labels",
}, []string{"version", "build_date", "commit_hash"})

func metricName(name string) string {
	return "ptvexporter_" + name
}

// observeHistogramTiming registers elapsed time to the histogram when used with `defer`
func observeHistogramTiming(start time.Time, hist prometheus.Histogram) {
	hist.Observe(time.Since(start).Seconds())
}

func registerBuildInfo(v *VersionInfo) {
	buildInfoGauge.WithLabelValues(v.Version, v.BuildDate, v.CommitHash).Set(1)
}
