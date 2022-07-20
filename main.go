package main

import (
	"github.com/Akkadius/rclone-multi/command"
	"os"
)

func main() {
	// process command line args
	if len(os.Args) > 1 {
		command.Run(os.Args)
	}
}
