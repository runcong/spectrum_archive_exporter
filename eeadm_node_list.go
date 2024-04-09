package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var specturm_archive_node_status = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_archive_node_status",
	Help: "Spectrum Archive Node Status (eeadm node list)",
}, []string{"status", "node_name", "node_ip"})

func node_status() {
	cmd := exec.Command("eeadm", "node", "list")
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
			// Extract the tape status
			status := strings.Fields(line)[1]
			node_name := strings.Fields(line)[7]
			node_ip := strings.Fields(line)[2]
			specturm_archive_node_status.WithLabelValues(status, node_name, node_ip).Set(0)
		}
	}

}
