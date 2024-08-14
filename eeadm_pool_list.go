package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var spectrum_archive_pool_used = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_archive_pool_used",
	Help: "Spectrum Archive Pool Used in TB (eeadm pool list)",
}, []string{"pool_name"})

var spectrum_archive_pool_available = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_archive_pool_available",
	Help: "Spectrum Archive Pool Available in TB (eeadm pool list)",
}, []string{"pool_name"})

func pool_status() {
	cmd := exec.Command("cat", "eeadm_pool_list.txt")
	// cmd := exec.Command("eeadm", "pool", "list")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Split the output into lines
	lines := strings.Split(string(output), "\n")

	// Remove the first line
	lines = lines[1:]

	// Process each line

	for _, line := range lines {
		// Extract the tape status
		if line != "" {
			pool_name := strings.Fields(line)[0]
			pool_used := strings.Fields(line)[2]
			pool_available := strings.Fields(line)[3]

			poolUsedFloat, err := strconv.ParseFloat(pool_used, 64)
			if err != nil {
				fmt.Println("Error converting pool_used to float64:", err)
				continue
			}

			poolAvailableFloat, err := strconv.ParseFloat(pool_available, 64)
			if err != nil {
				fmt.Println("Error converting pool_available to float64:", err)
				continue
			}

			spectrum_archive_pool_used.WithLabelValues(pool_name).Set(poolUsedFloat)
			spectrum_archive_pool_available.WithLabelValues(pool_name).Set(poolAvailableFloat)
		}
	}

}
