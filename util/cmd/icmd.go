package cmd

type ICmd interface {
	Exec(cmdStr string) (string, error)
	ExecASync(cmdStr string) error
	GetPidByPort(port int) int
	GetServerNameByPid(pid int) string
	GetPidByServerName(server string) int
	GetPidListByServerName(server string) []int
	GetTmpPath() string
	GetSeparator() string
	KillPID(pid int) bool
}
