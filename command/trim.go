package command

import (
	"fmt"
	"github.com/Akkadius/rclone-multi/rclone"
	"os"
)

// handles trim command
func handleTrim(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: trim [after-days] Example: trim 10 would trim anything older than 10 days")
		os.Exit(0)
	}

	rclone.Trim(args[2])
}
