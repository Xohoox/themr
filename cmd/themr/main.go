package main

import (
	"fmt"
	"os"
	"github.com/xohoox/themr/pkg/themr"
)

func printUsageScreenProfile() {
	fmt.Println("Commands:")
	fmt.Printf("  screenProfile                   Manage ScreenProfiles\n")
	fmt.Printf("    list                          List all ScreenProfiles\n")
	fmt.Printf("    select <name>                 Apply the <name> xrandr configuration\n")
	fmt.Printf("    addCurrent <name>             Add current xrandr configuration as <name> to themr\n")
	fmt.Printf("    rename <oldName> <newName>    Rename a ScreenProfile from <oldName> to <newName>\n\n")
}

func printUsageWallpaper() {
	fmt.Printf("  wallpaper                       Manage wallpapers\n")
	fmt.Printf("    list                          List all wallpaper\n")
	fmt.Printf("    select <name>                 Apply the <name> wallpaper group\n")
	fmt.Printf("    rename <oldName> <newName>    Rename a wallpaper from <oldName> to <newName>\n\n")
}

func printUsage() {
	fmt.Printf("usage: themr [-h | --help] <command> [args]\n\n")
	printUsageScreenProfile()
	printUsageWallpaper()
}

func main() {

	if err := themr.ReadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) <= 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-h", "--help":
		if len(os.Args) == 2 {
			printUsage()
			os.Exit(1)
		}

		switch os.Args[2] {
		case "screenProfile":
			fmt.Printf("usage: themr [-h | --help] screenProfile [args]\n\n")
			printUsageScreenProfile()
		case "wallpaper":
			fmt.Printf("usage: themr [-h | --help] wallpaper [args]\n\n")
			printUsageWallpaper()
		default:
			printUsage()
		}
		os.Exit(1)

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

			if err := themr.SelectScreenProfile(os.Args[3]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case "addCurrent":
			if len(os.Args) != 4 {
				fmt.Println("Expected 3 arguments for screenProfile add! Use --help for more information")
				os.Exit(1)
			}

			if err := themr.AddCurrentScreenProfile(os.Args[3]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case "rename":
			if len(os.Args) != 5 {
				fmt.Println("Expected 4 arguments for screenProfile select! Use --help for more information")
				os.Exit(1)
			}

			if err := themr.RenameScreenProfile(os.Args[3], os.Args[4]); err != nil {
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

			if	err := themr.SelectWallpaper(os.Args[3]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case "rename":
			if len(os.Args) != 5 {
				fmt.Println("Expected 4 arguments for wallpaper select! Use --help for more information")
				os.Exit(1)
			}
			
			if err := themr.RenameWallpaper(os.Args[3], os.Args[4]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			fmt.Printf("Unknown subcommand for wallpaper: %s! Use --help for more information\n", os.Args[1])
			os.Exit(1)
		}
	default:
		printUsage()
		os.Exit(1)
	}
}
