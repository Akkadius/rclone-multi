package rclone

import (
	"fmt"
	"github.com/Akkadius/rclone-multi/notify"
	"github.com/dustin/go-humanize"
	"github.com/pterm/pterm"
	"log"
	"strconv"
	"strings"
)

// Trim will trim files older than specified time
// example 10d will trim files older than 10 days
func Trim(time string) {
	pterm.Info.Printf("Deleting files older than [%v]\n", time)

	for _, remote := range getRemotes() {
		cmd := fmt.Sprintf("rclone lsl %v: --min-age=%v", remote, time)
		output, _ := execCmd("bash", "-c", cmd)

		pterm.Info.Printf("Remote [%v]\n", remote)

		for _, line := range strings.Split(output, "\n") {
			cols := strings.Split(line, " ")
			if len(cols) > 3 {
				size, err := strconv.ParseUint(cols[1], 10, 32)
				if err != nil {
					log.Println(err)
				}

				fileName := cols[4]

				pterm.Info.Printf("Deleting [%v] to [%v]\n", fileName, remote)

				del := fmt.Sprintf("rclone deletefile %v:%v", remote, fileName)
				execCmd("bash", "-c", del)

				pterm.Success.Printf("[DONE] Deleting [%v] from [%v]\n", fileName, remote)

				// send to notifiers
				notify.Send(
					fmt.Sprintf(
						"Trimmed (older than %v) [%v] (%v) via remote [%v]\n",
						time,
						fileName,
						humanize.Bytes(size),
						remote,
					),
				)
			}
		}
	}
}