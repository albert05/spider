package main

import (
	"fmt"
	"spider/util/logger"
	"spider/config"
	"spider/service"
)

func main() {
	logger.Info(fmt.Sprintf("driver is runing"))
	logger.Info(fmt.Sprintf("bind port %d", config.DriverPort))

	service.InitDriver()
}
