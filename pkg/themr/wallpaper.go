package themr

import (
	"errors"
	"fmt"
)

func ListWallpapers() {
	for _,v := range config.Wallpapers {
		fmt.Println("Name:", v.Name)
		for _, file := range v.Wallpapers {
			orientation := ""
			if file.Orientation == HorizontalOrientation {
				orientation = "horizontal"
			} else {
				orientation = "vertical"
			}

			fmt.Printf("\t[%s] %s %s\n", orientation, file.Mode, file.Path)
		}
	}
}

func selectWallpaperGroupFromConfig(wallpaperGroup string) (*WallpaperGroup, error) {
	for i := range config.Wallpapers {
		if config.Wallpapers[i].Name == wallpaperGroup {
			return &config.Wallpapers[i], nil
		}
	}
	return new(WallpaperGroup), errors.New("WallpaperGroup: " + wallpaperGroup + " is not in the config")
}

func SelectWallpaper(wallpaperGroup string) error {
	wallpaper, err := selectWallpaperGroupFromConfig(wallpaperGroup)
	_ = wallpaper

	if err != nil {
		return err
	}

	// TODO: implement

	return nil
}

func RenameWallpaper(oldName, newName string) error {
	wallpaper, err := selectWallpaperGroupFromConfig(oldName)

	if err != nil {
		return err
	}

	_, err = selectWallpaperGroupFromConfig(newName)

	if err == nil {
		return errors.New("Allready a WallpaperGroup with the name: " + newName + " in the config")
	}

	wallpaper.Name = newName

	return writeConfig()
}
