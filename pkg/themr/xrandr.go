package themr

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const xrandrCommand string = "xrandr"

func executeXrandr(args ...string) (string, error) {
	_, err := exec.LookPath(xrandrCommand)

	if err != nil {
		return "", err
	}

	command := exec.Command(xrandrCommand, args...)
	output, err := command.Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}

func parseMonitor(match []string) (Monitor, error) {
	// Convert Width Height X and Y to int
	indices := []int{4, 5, 6, 7}
	values := make([]int, len(indices))

	for i, idx := range indices {
		v, err := strconv.Atoi(match[idx])
		if err != nil {
			return Monitor{}, err
		}
		values[i] = v
	}

	monitor := Monitor { 
		Output: match[3],
		Primary: len(match[1]) == 1,
		Enabled: len(match[2]) == 1,
		Mode: MonitorMode {
			Width: values[0],
			Height: values[1],
		},
		Position: Position{
			X: values[2],
			Y: values[3],
		},
	}

	return monitor, nil
}

func GetMonitors() ([]Monitor, error) {
	output, err := executeXrandr("--listmonitors")
	monitors := []Monitor{}

	if err != nil {
		return  nil, err
	}

	lines := strings.Split(output, "\n")
	re := regexp.MustCompile(`\d+:\s+(\+?)(\*?)(\S+)\s+(\d+)\/\d+x(\d+)\/\d+\+(\d+)\+(\d+)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)

		if len(match) <= 1 {
			continue
		}

		monitor, err := parseMonitor(match)

		if err != nil {
			return nil, err
		}

		monitors = append(monitors, monitor)
	}
	return monitors, nil
}

func isOutputInMonitors(output string, monitors []Monitor) bool {
	for _, monitor := range monitors {
		if monitor.Output == output {
			return true
		}
	}
	return false
}

func SetMonitors(monitors []Monitor) error {
	var args []string

	connectedMonitors, err := GetMonitors()

	if err != nil {
		return err
	}

	for _, monitor := range monitors {
		if !isOutputInMonitors(monitor.Output, connectedMonitors) {
			return errors.New("Output: " + monitor.Output + " is not connected")
		}

		args = append(args, "--output", monitor.Output)

		if !monitor.Enabled {
			args = append(args, "--off")
			continue
		}

		if monitor.Primary {
			args = append(args, "--primary")
		}

		args = append(args, "--mode", monitor.Mode.String())
		args = append(args, "--pos", monitor.Position.String())
		args = append(args, "--rotate", string(monitor.Rotation))
	}

	for _, monitor := range connectedMonitors {
		if !isOutputInMonitors(monitor.Output, monitors) {
			args = append(args, "--output", monitor.Output, "--off")
		}
	}

	_, err= executeXrandr(args...)

	return err
}
