package main

import (
	"github.com/Akkadius/rclone-multi/internal/command"
	"github.com/pterm/pterm"
	"os"
)

func main() {
	// process command line args
	if len(os.Args) > 1 {
		err := command.Run(os.Args)
		if err != nil {
			pterm.Error.Println(err.Error())
		}
	}
}
