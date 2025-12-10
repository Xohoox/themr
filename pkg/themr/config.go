package themr

import (
	"os"
	"gopkg.in/yaml.v3"
)

type WallpaperGroup struct {
	Name string				`yaml:"name"`
	Wallpapers []Wallpaper	`yaml:"wallpapers"`
}

type ScreenProfile struct {
	Name string			`yaml:"name"`
	Monitors []Monitor	`yaml:"monitors"`
}

type Config struct {
	Wallpapers []WallpaperGroup		`yaml:"wallpapers"`
	ScreenProfiles []ScreenProfile	`yaml:"screen_profiles"`
}

func getConfigPath() string {
	// TODO: implement
	return "config.yml"
}

func readConfig() error {
	// TODO: read config from filesystem
	return nil
}

func writeConfig() error {
	yml, err := yaml.Marshal(&config)

	if err != nil {
		return  err
	}

	// TODO: Pretty Indentation
	return os.WriteFile(getConfigPath(), yml, 0644)
}

var config = Config {
	ScreenProfiles: []ScreenProfile {
		ScreenProfile {
			Name: "laptop",
			Monitors: []Monitor {
				Monitor {
					Output: "eDP-1",
					Primary: true,
					Enabled: true,
					Rotation: RotationNormal,
					Position: Position {
						X: 0,
						Y: 0,
					},
					Mode: MonitorMode {
						Width: 2880,
						Height: 1920,
					},
				},
			},
		},
		ScreenProfile {
			Name: "desk",
			Monitors: []Monitor {
				Monitor {
					Output: "DP-10",
					Primary: true,
					Enabled: true,
					Rotation: RotationLeft,
					Position: Position {
						X: 0,
						Y: 0,
					},
					Mode: MonitorMode {
						Width: 2560,
						Height: 1440,
					},
				},
				Monitor {
					Output: "DP-11",
					Primary: false,
					Enabled: true,
					Rotation: RotationNormal,
					Position: Position {
						X: 1440,
						Y: 693,
					},
					Mode: MonitorMode {
						Width: 2560,
						Height: 1440,
					},
				},
			},
		},
	},
	Wallpapers: []WallpaperGroup {
		WallpaperGroup {
			Name: "Zima",
			Wallpapers: []Wallpaper {
				Wallpaper {
					Mode: CenterMode,
					Path: "/home/fynn/.config/wallpaper/zima.jpg",
				},
			},
		},
		WallpaperGroup {
			Name: "Test",
			Wallpapers: []Wallpaper {
				Wallpaper {
					Mode: CenterMode,
					Path: "/home/fynn/.config/wallpaper/test1.jpg",
				},
			},
		},
	},
}
