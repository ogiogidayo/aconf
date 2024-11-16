package utils

import (
	"bytes"
	"os/exec"
)

func RunShellCommand(cmd []string) (string, error) {
	c := exec.Command(cmd[0], cmd[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &out
	c.Stderr = &stderr

	if err := c.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}
