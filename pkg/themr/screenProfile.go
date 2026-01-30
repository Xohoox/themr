package themr

import (
	"errors"
	"fmt"
)

func ListScreenProfile() {
	for _,v := range config.ScreenProfiles {
		fmt.Println("Name:", v.Name)

		for _, monitor := range v.Monitors {
			fmt.Printf("\tOutput: %s, Enabled: %t, Primary: %t, Mode: %s, Position: %s, Rotation: %s\n",
				monitor.Output,
				monitor.Enabled,
				monitor.Primary,
				monitor.Mode.String(),
				monitor.Position.String(),
				monitor.Rotation)
		}
	}
}

func selectScreenProfileFromConfig(screenProfile string) (*ScreenProfile, error) {
	for i := range config.ScreenProfiles {
		if config.ScreenProfiles[i].Name == screenProfile {
			return  &config.ScreenProfiles[i], nil
		}
	}
	return new(ScreenProfile), errors.New("ScreenProfile: " + screenProfile + " is not in the config")
}

func SelectScreenProfile(screenProfileName string) error {
	screenProfile, err := selectScreenProfileFromConfig(screenProfileName)

	if err != nil {
		return err
	}

	err = SetMonitors(screenProfile.Monitors)

	if err != nil {
		return err
	}

	return SelectWallpaper(config.DefaultWallpaper)
}

func AddCurrentScreenProfile(name string) error {
	montors, err := GetActiveMonitors()

	if err != nil {
		return err
	}

	screenProfile := ScreenProfile{
		Name: name,
		Monitors: montors,
	}

	config.ScreenProfiles = append(config.ScreenProfiles, screenProfile)

	return writeConfig()
}

func RenameScreenProfile(oldName, newName string) error {
	screenProfile, err := selectScreenProfileFromConfig(oldName)

	if err != nil {
		return err
	}

	_, err = selectScreenProfileFromConfig(newName)

	if err == nil {
		return errors.New("Allready a ScreenProfile with the name: " + newName + " in the config")
	}

	screenProfile.Name = newName

	return writeConfig()
}
