package rclone

import (
	"fmt"
	"github.com/Akkadius/rclone-multi/internal/notify"
	"github.com/dustin/go-humanize"
	"github.com/pterm/pterm"
	"log"
	"strconv"
	"strings"
)

// Trim will trim files older than specified time
// in the specified directory
// example 10d will trim files older than 10 days
func Trim(destPath string, time string) error {
	pterm.Info.Printf("Deleting files older than [%v]\n", time)

	for _, remote := range getRemotes() {
		cmd := fmt.Sprintf("rclone lsl %v:%v --min-age=%v", remote, destPath, time)
		output, _ := execCmd("bash", "-c", cmd)

		pterm.Info.Printf("Remote [%v]\n", remote)

		for _, line := range strings.Split(output, "\n") {
			line = strings.TrimSpace(line)

			// ex: 48720 2022-06-11 17:40:05.033892549 Menuetto.ttf
			cols := strings.Split(line, " ")
			if len(cols) > 3 {
				size, err := strconv.ParseUint(cols[0], 10, 32)
				if err != nil {
					log.Println(err)
				}

				fileName := cols[3]

				pterm.Info.Printf("Deleting [%v] via [%v]\n", fileName, remote)

				del := fmt.Sprintf("rclone deletefile %v:%v", remote, fileName)
				pterm.Info.Println(del)
				_, err = execCmd("bash", "-c", del)
				if err != nil {
					pterm.Error.Printf(
						"[Error] Failed to trim file [%v] via remote [%v] error [%v]\n",
						fileName,
						remote,
						err.Error(),
					)

					return err
				}

				pterm.Success.Printf("[DONE] Deleting [%v] from [%v]\n", fileName, remote)

				// send to notifiers
				notify.Info(
					fmt.Sprintf(
						"Trimmed [%v] (%v) via remote [%v] (older than %v)\n",
						fileName,
						humanize.Bytes(size),
						remote,
						time,
					),
				)
			}
		}
	}

	return nil
}
