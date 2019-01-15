package cmd

import "runtime"

var IC ICmd

func init() {
	if IsWindows() {
		IC = WindowsCmd{}
	} else {
		IC = LinuxCmd{}
	}
}

func IsWindows() bool {
	return "windows" == runtime.GOOS
}
