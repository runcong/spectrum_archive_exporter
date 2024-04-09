package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var specturm_archive_tape_status = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_archive_tape_status",
	Help: "Spectrum Archive Tape Status (eeadm tape list)",
}, []string{"status"})

func tape_status() {
	// Execute the command "cat eeadm_tape_list.txt"
	cmd := exec.Command("cat", "eeadm_tape_list.txt")
	// cmd := exec.Command("eeadm", "tape", "list")
	output, err := cmd.Output()
	status_nonok := 0
	status_error := 0
	status_degraded := 0

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
			if status == "error" {
				status_error++
			}
			if status == "degraded" {
				status_degraded++
			}
		}

		// status := strings.Fields(line)[1]

		// Print the tape status
		// fmt.Println("Tape status:", status)
	}

	specturm_archive_tape_status.WithLabelValues("ok").Set(float64(len(lines) - status_nonok - 1))
	specturm_archive_tape_status.WithLabelValues("non-ok").Set(float64(status_nonok))
	specturm_archive_tape_status.WithLabelValues("error").Set(float64(status_error))
	specturm_archive_tape_status.WithLabelValues("degraded").Set(float64(status_degraded))

	// fmt.Print("Number of tapes with status non-OK: ", status_nonok, "\n")
	// fmt.Print("Number of tapes with status error: ", status_error, "\n")
	// fmt.Print("Number of tapes with status degraded: ", status_degraded, "\n")
}
