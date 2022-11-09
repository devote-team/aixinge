//go:build windows
// +build windows

package utils

import (
	"os/exec"
	"syscall"
)

func OpenUri(uri string) {
	cmd := exec.Command(`cmd`, `/c`, `start`, uri)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
