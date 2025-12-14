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
	DefaultWallpaper string			`yaml:"defaultWallpapers"`
}

func getConfigPath() string {
	// TODO: implement
	return "config.yml"
}

func ReadConfig() error {
	data, err := os.ReadFile(getConfigPath())

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	return nil
}

func writeConfig() error {
	// TODO: Check config
	yml, err := yaml.Marshal(&config)

	if err != nil {
		return  err
	}

	// TODO: Pretty Indentation
	return os.WriteFile(getConfigPath(), yml, 0644)
}

var config = Config{}
