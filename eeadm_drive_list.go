package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var specturm_archive_drive_status = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_archive_drive_status",
	Help: "Spectrum Archive Drive Status (eeadm drive list)",
}, []string{"status"})

func drive_status() {
	cmd := exec.Command("cat", "eeadm_drive_list.txt")
	// cmd := exec.Command("eeadm", "drive", "list")
	output, err := cmd.Output()
	status_nonok := 0

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
			if status != "ok" {
				status_nonok++
			}
		}
	}

	specturm_archive_drive_status.WithLabelValues("ok").Set(float64(len(lines) - status_nonok - 1))
	specturm_archive_drive_status.WithLabelValues("non-ok").Set(float64(status_nonok))
	// fmt.Print(len(lines))
}
