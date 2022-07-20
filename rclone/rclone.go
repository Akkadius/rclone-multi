package rclone

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// gets the configured remotes for rclone
func getRemotes() []string {
	remotes, _ := execCmd("rclone", "listremotes")
	remotes = strings.TrimSpace(strings.ReplaceAll(remotes, ":", ""))
	return strings.Split(remotes, "\n")
}

// crude function that reduces lines for repeated command fetches
func execCmd(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New(stderr.String())
	}

	return stdout.String(), nil
}
