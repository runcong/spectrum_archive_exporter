package main

import "github.com/prometheus/client_golang/prometheus"

var (
	myGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_gauge",
		Help: "This is my gauge",
	},
	)

	myCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_counter",
			Help: "This is my counter",
		},
		[]string{"label1", "label2"},
	)

	diskSizeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "disk_size",
		Help: "Current disk size",
	})
)

func updateMyCounter() {
	myCounter.WithLabelValues("AAA", "BBB").Inc()
}

func updateDiskSizeMetric() {
	diskSize := 100 * 1024 * 1024 * 1024
	diskSizeGauge.Set(float64(diskSize))
}

// type SaCollector struct {
// 	myGauge			prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "my_gauge",
// 		Help: "This is my gauge",
// 	},
// 	)
// 	myCounter
// 	diskSizeGauge
// }
