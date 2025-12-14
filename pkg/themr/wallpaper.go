package themr

import (
	"math/rand/v2"
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

func getRandomWallpaper(group *WallpaperGroup, rotation Rotation) (Wallpaper, error) {
	var wallpaper []Wallpaper
	orientation := HorizontalOrientation

	if rotation == RotationLeft || rotation == RotationRight {
		orientation = VerticalOrientation
	}

	for _, v := range group.Wallpapers {
		if v.Orientation == orientation {
			wallpaper = append(wallpaper, v)
		}
	}

	if len(wallpaper) == 0 {
		return Wallpaper{}, errors.New("No Wallpaper with Orientation " + string(orientation) + " in Group: " + group.Name)
	}

	return wallpaper[rand.IntN(len(wallpaper))], nil
}

func SelectWallpaper(wallpaperGroupName string) error {
	wallpaperGroup, err := selectWallpaperGroupFromConfig(wallpaperGroupName)

	if err != nil {
		return err
	}

	monitors, err := GetActiveMonitors()

	if err != nil {
		return err
	}

	wallpapers := []Wallpaper{};

	for _, monitor := range monitors {
		wallpaper, err := getRandomWallpaper(wallpaperGroup, monitor.Rotation)

		if err != nil {
			return err
		}

		wallpapers = append(wallpapers, wallpaper)
	}

	return SetWallpaper(wallpapers, monitors)
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
