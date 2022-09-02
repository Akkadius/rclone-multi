package command

import (
	"github.com/Akkadius/rclone-multi/internal/rclone"
	"github.com/pterm/pterm"
	"os"
)

func Run(args []string) error {
	command := ""
	if len(args) > 1 {
		command = os.Args[1]
	}

	// upload
	if command == "upload" {
		if len(args) < 3 {
			pterm.Info.Println("Usage: upload [source-file] [destination-path]")
			os.Exit(0)
		}

		sourceFile := args[2]
		destPath := ""
		if len(args) > 3 {
			destPath = args[3]
		}

		return rclone.Upload(sourceFile, destPath)
	}

	// trim
	if command == "trim" {
		if len(args) < 3 {
			pterm.Info.Println("Usage: trim [path] [duration] Delete files older than this in seconds or ms|s|m|h|d|w|M|y Ex: 10d or 10s")
			os.Exit(0)
		}

		destinationPath := args[2]
		duration := args[3]
		if destinationPath == "." {
			destinationPath = ""
		}

		return rclone.Trim(destinationPath, duration)
	}

	// exist
	// checks for if backups or files exist within this time in location
	if command == "exist" {
		if len(args) < 3 {
			pterm.Info.Println("Usage: exist [destination-path] [duration] Check for existence of files newer than this in seconds or alert. ms|s|m|h|d|w|M|y Ex: 10d or 10s")
			os.Exit(0)
		}

		duration := args[3]
		destinationPath := args[2]
		if destinationPath == "." {
			destinationPath = ""
		}

		return rclone.Exist(destinationPath, duration)
	}

	pterm.NewRGB(15, 199, 209).Println("")
	pterm.NewRGB(15, 199, 209).Println("[rclone-multi] A simple wrapper for rclone for multi-remote backup operations")
	pterm.NewRGB(15, 199, 209).Println("")
	pterm.NewRGB(15, 199, 209).Println("> upload [source-file] [destination-path]")
	pterm.NewRGB(15, 199, 209).Println("> trim [destination-path (. for current dir)] [duration] Delete files older than this in seconds or ms|s|m|h|d|w|M|y Ex: 10d or 10s")
	pterm.NewRGB(15, 199, 209).Println("> exist [destination-path (. for current dir)] [duration] Check for existence of files newer than this in seconds or alert. ms|s|m|h|d|w|M|y Ex: 10d or 10s")
	pterm.NewRGB(15, 199, 209).Println("")

	return nil
}
