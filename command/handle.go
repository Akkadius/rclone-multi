package command

import (
	"fmt"
	"github.com/Akkadius/rclone-multi/rclone"
	"os"
)

func Run(args []string) {
	command := os.Args[1]
	if command == "upload" {
		if len(args) < 3 {
			fmt.Println("Usage: upload [source-file] [destination-path]")
			os.Exit(0)
		}

		sourceFile := args[2]
		destPath := ""
		if len(args) > 3 {
			destPath = args[3]
		}

		rclone.Upload(sourceFile, destPath)
	}
	if command == "trim" {
		if len(args) < 3 {
			fmt.Println("Usage: trim [after-days] Example: trim 10 would trim anything older than 10 days")
			os.Exit(0)
		}

		rclone.Trim(args[2])
	}
}
