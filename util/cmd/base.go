package cmd

type BaseCmd struct {
}

func (c BaseCmd) Exec(cmdStr string) (string, error) {
	return "", nil
}

func (c BaseCmd) ExecASync(cmdStr string) error {
	return nil
}

func (c BaseCmd) GetPidByPort(port int) int {
	return 0
}

func (c BaseCmd) GetServerNameByPid(pid int) string {
	return ""
}

func (c BaseCmd) GetPidByServerName(server string) int {
	return 0
}

func (c BaseCmd) GetPidListByServerName(server string) []int {
	return nil
}

func (c BaseCmd) GetTmpPath() string {
	return ""
}

func (c BaseCmd) GetSeparator() string {
	return ""
}

func (c BaseCmd) KillPID(pid int) bool {
	return true
}
