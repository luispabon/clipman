package main

import (
	"fmt"
	"os/exec"
	"time"
)

// notify sends a desktop notification to Linux users
// Urgency: low, normal, critical
func notify(app string, message string, urgency string, timeout time.Duration) error {

	// notify-send only accepts milliseconds
	millisec := fmt.Sprintf("%v", timeout.Seconds()*1000)

	args := []string{"-a", app, "-u", urgency, "-t", millisec, message}

	return exec.Command("notify-send", args...).Run()

}
