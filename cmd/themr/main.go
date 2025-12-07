package main

import (
	"fmt"
	"os"
	"github.com/xohoox/themr/pkg/themr"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'screenProfile', 'theme' or 'wallpaper' as subcommand! Use --help for more information")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "screenProfile":
		if len(os.Args) == 2 {
			fmt.Println("Unknown command! Use --help for more information")
		}

		switch os.Args[2] {
		case "list":
			if len(os.Args) != 3 {
				fmt.Println("To many arguments 2 expected for screenProfile list! Use --help for more information")
				os.Exit(1)
			}
			themr.ListScreenProfile()
		case "select":
			if len(os.Args) != 4 {
				fmt.Println("Expected 3 arguments for screenProfile select! Use --help for more information")
				os.Exit(1)
			}

			err := themr.SelectScreenProfile(os.Args[3])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case "rename":
			if len(os.Args) != 5 {
				fmt.Println("Expected 4 arguments for screenProfile select! Use --help for more information")
				os.Exit(1)
			}
			
			err := themr.RenameScreenProfile(os.Args[3], os.Args[4])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			fmt.Printf("Unknown subcommand for wallpaper: %s! Use --help for more information\n", os.Args[1])
			os.Exit(1)
		}
	case "wallpaper":
		if len(os.Args) == 2 {
			fmt.Println("Unknown command! Use --help for more information")
		}

		switch os.Args[2] {
		case "list":
			if len(os.Args) != 3 {
				fmt.Println("To many arguments 2 expected for wallpaper list! Use --help for more information")
				os.Exit(1)
			}
			themr.ListWallpapers()
		case "select":
			if len(os.Args) != 4 {
				fmt.Println("Expected 3 arguments for wallpaper select! Use --help for more information")
				os.Exit(1)
			}

			err := themr.SelectWallpaper(os.Args[3])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case "rename":
			if len(os.Args) != 5 {
				fmt.Println("Expected 4 arguments for wallpaper select! Use --help for more information")
				os.Exit(1)
			}
			
			err := themr.RenameWallpaper(os.Args[3], os.Args[4])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			fmt.Printf("Unknown subcommand for wallpaper: %s! Use --help for more information\n", os.Args[1])
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown subcommand: %s! Use --help for more information\n", os.Args[1])
		os.Exit(1)
	}
}
