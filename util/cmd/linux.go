// LINUX CMD
package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

type LinuxCmd struct {
	BaseCmd
}

func (c LinuxCmd) Exec(cmdStr string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(out)), nil
}

func (c LinuxCmd) ExecASync(cmdStr string) error {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	return cmd.Start()
}

func (c LinuxCmd) GetPidByPort(port int) int {
	return 0
}

func (c LinuxCmd) GetServerNameByPid(pid int) string {
	return ""
}

func (c LinuxCmd) GetPidByServerName(server string) int {
	return 0
}

func (c LinuxCmd) GetPidListByServerName(server string) []int {
	return nil
}

func (c LinuxCmd) GetTmpPath() string {
	return "/tmp/"
}

func (c LinuxCmd) GetSeparator() string {
	return "/"
}

func (c LinuxCmd) KillPID(pid int) bool {
	_, err := c.Exec(fmt.Sprintf("kill -9 %d", pid))
	return err == nil
}
