package rclone

import (
	"fmt"
	"github.com/Akkadius/rclone-multi/internal/notify"
	"github.com/pterm/pterm"
	"strings"
)

// Exist will check for files that exist within a timeframe or alert
func Exist(destPath string, duration string) error {
	pterm.Info.Printf("Checking for existence of files duration [%v] via remotes to path [%v]\n", duration, destPath)

	for _, remote := range getRemotes() {
		cmd := fmt.Sprintf("rclone lsl %v:%v --max-age=%v", remote, destPath, duration)
		pterm.Info.Println(cmd)

		output, _ := execCmd("bash", "-c", cmd)
		pterm.Info.Println(output)

		pterm.Info.Printf("Remote [%v]\n", remote)

		// no files
		if len(strings.Split(output, "\n")) == 1 {
			notify.Alert(
				fmt.Sprintf(
					"No files or backups found for remote [%v] path [%v] duration [%v]",
					remote,
					destPath,
					duration,
				),
			)
		}
	}

	return nil
}
