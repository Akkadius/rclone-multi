package command

import (
	"fmt"
	"github.com/Akkadius/rclone-multi/rclone"
	"os"
)

// handles upload command
func handleUpload(args []string) {
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
