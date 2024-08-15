package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var spectrum_archive_tape_state = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_archive_tape_state",
	Help: "Spectrum Archive Tape State (eeadm tape list)",
}, []string{"state"})

var spectrum_archive_tape_status = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_archive_tape_status",
	Help: "Spectrum Archive Tape Status (eeadm tape list)",
}, []string{"status"})

func tape_status() {
	cmd := exec.Command("eeadm", "tape", "list")
	output, err := cmd.Output()
	unassigned_tape := 0
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
		// Extract the tape status/number of unassigned tapes
		if line != "" {
			//library_name := strings.Fields(line)[8]
			tape_state := strings.Fields(line)[2]
			status := strings.Fields(line)[1]

			if tape_state == "unassigned" {
				unassigned_tape++
			}
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

	spectrum_archive_tape_state.WithLabelValues("unassigned").Set(float64(unassigned_tape))
	spectrum_archive_tape_status.WithLabelValues("ok").Set(float64(len(lines) - status_nonok - 1))
	spectrum_archive_tape_status.WithLabelValues("non-ok").Set(float64(status_nonok))
	spectrum_archive_tape_status.WithLabelValues("error").Set(float64(status_error))
	spectrum_archive_tape_status.WithLabelValues("degraded").Set(float64(status_degraded))

}
