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

func getInt(str string) int {
	v, err := strconv.Atoi(str)

	if err != nil {
		return 0
	}

	return v
}

func GetAllMonitors() ([]Monitor, error) {
	monitors := []Monitor{}
	output, err := executeXrandr()

	if err != nil {
		return  nil, err
	}

	lines := strings.Split(output, "\n")
	re := regexp.MustCompile(`^(\S+)\s+(dis)?connected(\s+(primary)?\s*(\d+)x(\d+)\+(\d+)\+(\d+)\s+(left|right|inverted)?)?`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)

		if len(match) <= 1 {
			continue
		}

		// Skip not connected monitors
		if match[2] == "dis" {
			continue
		}

		rotation := Rotation(match[9])

		if rotation == "" {
			rotation = RotationNormal
		}

		monitor := Monitor {
			Output: match[1],
			Primary: match[4] == "primary",
			Enabled: len(match[3]) != 0,
			Rotation: rotation,
			Mode: MonitorMode {
				Width: getInt(match[5]),
				Height: getInt(match[6]),
			},
			Position: Position{
				X: getInt(match[7]),
				Y: getInt(match[8]),
			},
		}

		monitors = append(monitors, monitor)
	}

	return monitors, nil
}

func GetActiveMonitors() ([]Monitor, error) {
	activeMonitors := []Monitor{}
	allMonitors, err := GetAllMonitors()

	if err != nil {
		return nil, err
	}

	for _, monitor := range allMonitors {
		if monitor.Enabled {
			activeMonitors = append(activeMonitors, monitor)
		}
	}

	return activeMonitors, nil
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

	connectedMonitors, err := GetAllMonitors()

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
