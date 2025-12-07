package themr

type WallpaperGroup struct {
	Name string
	Wallpapers []Wallpaper
}

type ScreenProfile struct {
	Name string
	Monitors []Monitor
}

type Config struct {
	Wallpapers []WallpaperGroup
	ScreenProfiles []ScreenProfile
}

func readConfig() error {
	// TODO: read config from filesystem
	return nil
}

func writeConfig() error {
	// TODO: write config to filesystem
	return nil
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
