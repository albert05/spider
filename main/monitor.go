package main

import (
	"fmt"
	"os"
	"spider/config"
	"spider/model"
	"spider/service"
	. "spider/util/cmd"
)

func main() {
	service.Init(false)

	al := model.GetAccessList()

	// start driver
	if len(al) > 0 {
		runScript(config.ScriptRunDriverName, "driver")
	}

	for _, item := range al {
		if item.IsRun == true {
			runScript(config.ScriptRunName, item.Id)
		}
	}
}

func runScript(runName, id string) {
	pid := IC.GetPidByServerName(fmt.Sprintf("%s  -id %s", runName, id))
	if pid == 0 {
		IC.ExecASync(makeCmd(runName, id))
	}
}

func makeCmd(runName, id string) string {
	path := fmt.Sprintf("%s\\src\\%s\\main", os.Getenv("GOPATH"), config.ProNAME)

	return fmt.Sprintf("%s\\%s -id %s", path, runName, id)
}
