package rclone

import (
	"fmt"
	"github.com/Akkadius/rclone-multi/notify"
	"github.com/dustin/go-humanize"
	"github.com/pterm/pterm"
	"log"
	"os"
)

// Upload will upload a file to multiple destinations
func Upload(sourceFile string, destPath string) {
	pterm.Info.Printf("Uploading file [%v] via remotes to path [%v]\n", sourceFile, destPath)

	for _, remote := range getRemotes() {

		execArgs := []string{
			"copy",
			sourceFile,
		}

		// build destination
		destination := fmt.Sprintf("%v:%v", remote, destPath)
		execArgs = append(execArgs, destination)

		pterm.Info.Printf("-- Remote [%v] Uploading [%v] to [%v]\n", remote, sourceFile, destination)

		_, err := execCmd("rclone", execArgs...)
		if err != nil {
			pterm.Error.Printf(
				"[Error] Uploading [%v] to [%v] Error (%v)\n",
				sourceFile,
				destination,
				err.Error(),
			)
		}

		pterm.Success.Printf("-- Remote [%v] Uploaded [%v] to [%v]\n", remote, sourceFile, destination)

		fi, err := os.Stat(sourceFile)
		if err != nil {
			log.Println(err)
		}

		// get the size
		size := uint64(fi.Size())

		// send to notifiers
		notify.Send(
			fmt.Sprintf(
				"Uploaded [%v] (%v) to [%v]\n",
				sourceFile,
				humanize.Bytes(size),
				destination,
			),
		)
	}
}