package command

import "os"

func Run(args []string) {
	command := os.Args[1]
	if command == "upload" {
		handleUpload(os.Args)
	}
	if command == "trim" {
		handleTrim(os.Args)
	}
}
