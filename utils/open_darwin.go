//go:build darwin
// +build darwin

package utils

import (
	"os/exec"
)

func OpenUri(uri string) {
	exec.Command(`open`, uri).Start()
}
