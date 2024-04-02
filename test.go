package main

import (
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
)

type TestMetrics struct {
	alloc       float64
	idle        float64
	total       float64
	utilization float64
}

func TestGetMetrics() *TestMetrics {
	return ParseTestMetrics()
}

func ParseTestMetrics() *TestMetrics {
	var tm TestMetrics
	tm.alloc = 55
	tm.idle = rand.Float64()
	tm.total = rand.Float64()
	tm.utilization = rand.Float64()
	return &tm
}

func NewTestCollector() *TestCollector {
	return &TestCollector{
		alloc:       prometheus.NewDesc("Test_alloc", "Allocated Test", nil, nil),
		idle:        prometheus.NewDesc("Test_idle", "Idle Test", nil, nil),
		total:       prometheus.NewDesc("Test_total", "Total Test", nil, nil),
		utilization: prometheus.NewDesc("Test_utilization", "Total Test utilization", nil, nil),
	}
}

type TestCollector struct {
	alloc       *prometheus.Desc
	idle        *prometheus.Desc
	total       *prometheus.Desc
	utilization *prometheus.Desc
}

// Send all metric descriptions
func (cc *TestCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- cc.alloc
	ch <- cc.idle
	ch <- cc.total
	ch <- cc.utilization
}
func (cc *TestCollector) Collect(ch chan<- prometheus.Metric) {
	cm := TestGetMetrics()
	ch <- prometheus.MustNewConstMetric(cc.alloc, prometheus.GaugeValue, cm.alloc)
	ch <- prometheus.MustNewConstMetric(cc.idle, prometheus.GaugeValue, cm.idle)
	ch <- prometheus.MustNewConstMetric(cc.total, prometheus.GaugeValue, cm.total)
	ch <- prometheus.MustNewConstMetric(cc.utilization, prometheus.GaugeValue, cm.utilization)
}
