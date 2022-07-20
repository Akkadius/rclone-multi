package command

import (
	"github.com/Akkadius/rclone-multi/rclone"
	"github.com/pterm/pterm"
	"os"
)

func Run(args []string) error {
	command := os.Args[1]
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
	if command == "trim" {
		if len(args) < 3 {
			pterm.Info.Println("Usage: trim [after-days] Example: trim 10 would trim anything older than 10 days")
			os.Exit(0)
		}

		return rclone.Trim(args[2])
	}
	return nil
}
