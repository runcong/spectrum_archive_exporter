package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var listenAddress = flag.String(
	"listen-address",
	":8080",
	"The address to listen on for HTTP requests.")

var reg = prometheus.NewRegistry()

func registerMetrics() {
	// reg.MustRegister(NewTestCollector())
	// reg.MustRegister(NewTest2Collector())
	// reg.MustRegister(myGauge)
	// reg.MustRegister(myCounter)
	// reg.MustRegister(diskSizeGauge)
	reg.MustRegister(specturm_archive_tape_status)
	reg.MustRegister(specturm_archive_drive_status)
	reg.MustRegister(specturm_archive_pool_used)
	reg.MustRegister(specturm_archive_pool_available)
	reg.MustRegister(specturm_archive_task_status)
	// reg.MustRegister()
}

func main() {
	registerMetrics()
	flag.Parse()
	log.Printf("Starting Server: %s", *listenAddress)

	// http.Handle("/metrics", promhttp.Handler())
	http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// updateDiskSizeMetric()
		// updateMyCounter()
		tape_status()
		drive_status()
		pool_status()
		task_status()
		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
	}))
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
