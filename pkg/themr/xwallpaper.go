package themr

import (
	"errors"
	"os/exec"
)

type WallpaperMode string

const (
	CenterMode WallpaperMode = "--center"
	FocusMode WallpaperMode = "--focus"
	MaximizeMode WallpaperMode = "--maximize"
	StretchMode WallpaperMode = "--stretch"
	TileMode WallpaperMode = "--tile"
	ZoomMode WallpaperMode = "--zoom"
)

type WallpaperOrientation string

const (
	HorizontalOrientation WallpaperOrientation = "horizontal"
	VerticalOrientation WallpaperOrientation = "vertical"
)

type Wallpaper struct {
	Mode WallpaperMode					`yaml:wallpaperMode`
	Path string							`yaml:path`
	Orientation WallpaperOrientation	`yaml:orientation`
}

const xwallpaperCommand = "xwallpaper"

func executeXwallpaper(args ...string) error {
	_, err := exec.LookPath(xwallpaperCommand)

	if err != nil {
		return err
	}

	command := exec.Command(xwallpaperCommand, args...)
	_, err = command.Output()

	return err
}

func SetWallpaper(wallpapers []Wallpaper, monitors []Monitor) error {
	var args []string

	if len(wallpapers) != len(monitors) {
		return  errors.New("Lenght of wallpapers and monitors did not match")
	}

	for i, monitor := range monitors {
		if !monitor.Enabled {
			return errors.New("Monitor" + monitor.Output + " is not enabled")
		}

		args = append(args, "--output", monitor.Output, string(wallpapers[i].Mode), wallpapers[i].Path)
	}

	return executeXwallpaper(args...)
}
