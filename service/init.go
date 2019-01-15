package service

import (
	"errors"
	"flag"
	"fmt"
	"spider/common"
	"spider/config"
	"spider/model"
	"spider/util/logger"
)

func InitArgs() {
	flag.StringVar(&config.AccessID, "id", "", "alipay access id")
	flag.Parse()

	if config.AccessID != "driver" && !model.CheckAccessIDExists(config.AccessID) {
		panic(errors.New(fmt.Sprintf("access id[%s] is not exists", config.AccessID)))
	}
}

func InitLogs() {
	logger.SetConsole(false)
	logger.SetLevel(logger.INFO)
	logger.SetRollingDaily(common.GetLogPath)
}

func InitDriver() {
	if isRestart() {
		driver := &ChromeDriver{}
		driver.init()
	}
}

func Init(isArgs bool) {
	if isArgs {
		InitArgs()
	}

	//InitDriver()
	InitLogs()
}
