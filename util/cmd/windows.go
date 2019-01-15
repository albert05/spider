// WINDOWS CMD
package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

type WindowsCmd struct {
	BaseCmd
}

func (c WindowsCmd) Exec(cmdStr string) (string, error) {
	cmd := exec.Command("cmd.exe")

	w, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	out := &bytes.Buffer{}
	cmd.Stdout = out
	if err := cmd.Start(); err != nil {
		return "", err
	}

	io.WriteString(w, cmdStr)
	io.WriteString(w, "\n")
	w.Close()
	if err := cmd.Wait(); err != nil {
		return "", err
	}

	// filter head
	res := strings.Split(out.String(), cmdStr)
	if len(res) == 2 {
		// filter tail
		r := strings.Split(res[1], "\r\n\r\n")
		if len(r) == 2 {
			return strings.TrimSpace(r[0]), nil
		}
	}

	return "", errors.New(fmt.Sprintf("result[%s] exception", out.String()))
}

func (c WindowsCmd) ExecASync(cmdStr string) error {
	//_, err := os.StartProcess(fmt.Sprintf("start /c %s", cmdStr), nil, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	//return err

	cmd := exec.Command("cmd.exe", "/c", "start "+cmdStr)

	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}

func (c WindowsCmd) GetPidByPort(port int) int {
	r, err := c.Exec(fmt.Sprintf(`netstat -ano | findstr "%d" | findstr "LISTENING"`, port))
	if r == "" || err != nil {
		return 0
	}

	rs, err := c.Exec(fmt.Sprintf(`for /f "tokens=5" %s in ("%s") do echo %s`, "%i", strings.Split(r, "\r\n")[0], "%i"))
	if err != nil {
		return 0
	}

	t := strings.Split(rs, "\r\n")
	if len(t) == 2 {
		i, err := strconv.ParseInt(t[1], 10, 64)
		if err == nil {
			return int(i)
		}
	}

	return 0
}

func (c WindowsCmd) GetServerNameByPid(pid int) string {
	r, err := c.Exec(fmt.Sprintf(`tasklist | findstr "%d"`, pid))
	if r == "" || err != nil {
		return ""
	}

	rs, err := c.Exec(fmt.Sprintf(`for /f "tokens=1" %s in ("%s") do echo %s`, "%i", strings.Split(r, "\r\n")[0], "%i"))
	if err != nil {
		return ""
	}

	t := strings.Split(rs, "\r\n")
	if len(t) == 2 {
		return strings.TrimSpace(t[1])
	}

	return rs
}

func (c WindowsCmd) GetPidByServerName(server string) int {
	r, err := c.Exec(fmt.Sprintf(`wmic process where "commandline like '%s' and caption<>'WMIC.exe'" get processid /value
`, "%"+server+"%"))
	if err != nil {
		return 0
	}

	list := strings.Split(r, "\r\n")

	for _, v := range list {
		t := strings.TrimSpace(v)
		if strings.Contains(t, "ProcessId=") {
			p := strings.Split(t, "=")[1]

			if pid, err := strconv.ParseInt(strings.TrimSpace(p), 10, 64); err == nil {
				return int(pid)
			}
		}
	}

	return 0
}

func (c WindowsCmd) GetPidListByServerName(server string) []int {
	r, err := c.Exec(fmt.Sprintf(`wmic process where "commandline like '%s' and caption<>'WMIC.exe'" get processid /value
`, "%"+server+"%"))
	if err != nil {
		return nil
	}

	list := strings.Split(r, "\r\n")

	il := make([]int, 0)
	for _, v := range list {
		t := strings.TrimSpace(v)
		if strings.Contains(t, "ProcessId=") {
			p := strings.Split(t, "=")[1]

			if pid, err := strconv.ParseInt(strings.TrimSpace(p), 10, 64); err == nil {
				il = append(il, int(pid))
			}
		}
	}

	if len(il) > 0 {
		return il
	}

	return nil
}

func (c WindowsCmd) GetTmpPath() string {
	return "C:\\"
}

func (c WindowsCmd) GetSeparator() string {
	return "\\"
}

func (c WindowsCmd) KillPID(pid int) bool {
	_, err := c.Exec(fmt.Sprintf("taskkill /pid %d /f", pid))
	return err == nil
}
