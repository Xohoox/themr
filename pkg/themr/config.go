package themr

import (
	"os"
	"gopkg.in/yaml.v3"
)

type WallpaperGroup struct {
	Name string				`yaml:"name"`
	Wallpapers []Wallpaper	`yaml:"wallpapers"`
	InitScript string		`yaml:initscript"`
}

type ScreenProfile struct {
	Name string			`yaml:"name"`
	Monitors []Monitor	`yaml:"monitors"`
	InitScript string	`yaml:initscript"`
}

type Config struct {
	Wallpapers []WallpaperGroup		`yaml:"wallpapers"`
	ScreenProfiles []ScreenProfile	`yaml:"screen_profiles"`
	DefaultWallpaper string			`yaml:"defaultWallpapers"`
}

func getConfigPath() (string, error) {
	configFile := ""
	themrEnvVar := os.Getenv("THEMR_CONFIG_DIR")
	xdgEnvVar := os.Getenv("XDG_CONFIG_HOME")
	homeEnvVar := os.Getenv("HOME")

	if themrEnvVar != "" {
		if err := os.MkdirAll(themrEnvVar, os.ModePerm); err != nil {
			return  "", err
		}
		configFile = themrEnvVar + "/themr.yml"
	} else if xdgEnvVar != "" {
		if _, err := os.Stat(xdgEnvVar); err != nil {
			return "", err
		}
		if err := os.MkdirAll(xdgEnvVar + "/themr", os.ModePerm); err != nil {
			return  "", err
		}
		configFile = xdgEnvVar + "/themr/themr.yml"
	} else {
		if _, err := os.Stat(homeEnvVar + "/.config"); err != nil {
			configFile = homeEnvVar + "/.themr.yml"
		} else {
			if err := os.MkdirAll(homeEnvVar + "/.config/themr", os.ModePerm); err != nil {
				return  "", err
			}
			configFile = homeEnvVar + "/.config/themr/themr.yml"
		}
	}

	return configFile, nil
}

func ReadConfig() error {
	configFile, err := getConfigPath()

	if err != nil {
		return err
	}

	if _, err := os.Stat(configFile); err != nil {
		config = Config{}
		return nil
	}

	data, err := os.ReadFile(configFile)

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

	configFile, err := getConfigPath()

	if err != nil {
		return err
	}

	// TODO: Pretty Indentation
	return os.WriteFile(configFile, yml, 0644)
}

var config = Config{}
