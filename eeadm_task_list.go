package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var specturm_archive_task_status = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_archive_task_status",
	Help: "Spectrum Archive Task Status (eeadm task list)",
}, []string{"status"})

func task_status() {
	cmd := exec.Command("cat", "eeadm_task_list.txt")
	// cmd := exec.Command("eeadm", "task", "list")
	output, err := cmd.Output()
	status_running := 0
	status_failed := 0
	status_completed := 0

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
			status := strings.Fields(line)[3]
			if status == "running" {
				status_running++
			}
			if status == "failed" {
				status_failed++
			}
			if status == "completed" {
				status_completed++
			}
		}
	}

	specturm_archive_task_status.WithLabelValues("non-running").Set(float64(len(lines) - status_running - 1))
	specturm_archive_task_status.WithLabelValues("running").Set(float64(status_running))
	specturm_archive_task_status.WithLabelValues("failed").Set(float64(status_failed))
	specturm_archive_task_status.WithLabelValues("completed").Set(float64(status_completed))
	// fmt.Print(len(lines))
}
