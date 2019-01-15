package service

import (
	"github.com/tebeka/selenium"
	"spider/config"
	"spider/util/cmd"
)

type ChromeDriver struct {
	Service *selenium.Service
	Pid     int
}

func isRestart() bool {
	// check port occupy
	pid := cmd.IC.GetPidByPort(config.DriverPort)
	if pid == 0 {
		return true
	}

	server := cmd.IC.GetServerNameByPid(pid)
	if server == config.DriverName {
		return false
	}

	// kill other process
	cmd.IC.KillPID(pid)

	return true
}

func (this *ChromeDriver) init() {
	//如果seleniumServer没有启动，就启动一个seleniumServer所需要的参数，可以为空，示例请参见https://github.com/tebeka/selenium/blob/master/example_test.go
	opts := []selenium.ServiceOption{}
	//opts := []selenium.ServiceOption{
	//    selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
	//    selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
	//}

	//selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(config.DriverURL, config.DriverPort, opts...)
	if nil != err {
		panic(err)
	}

	this.Service = service
}

func (this *ChromeDriver) Clear() {
	if this.Service != nil {
		this.Service.Stop()
	}
}
