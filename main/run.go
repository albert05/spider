package main

import (
	"encoding/json"
	"fmt"
	"spider/exception"
	"spider/service/alipay"
	"spider/util/dates"
	"spider/util/logger"
	"time"
	"errors"
	"spider/service"
	. "spider/model"
)


func main() {
	service.Init(true)
	logger.Info(fmt.Sprintf("alipay is starting"))

	chrome := alipay.ChromeObj{}
	defer func() {
		exception.Handle(true)
		chrome.Close()
	}()

	chrome.Init()
	chrome.Start()

	chrome.LastTime = dates.NowTime() - CurAccount.FirstSpiderTime

	for {
		if !chrome.MLogin(CurAccount) {
			panic(errors.New(fmt.Sprintf("aipay login failed, %s", chrome.String())))
		}

		b, _ := json.Marshal(chrome.GetBill())
		logger.Info(fmt.Sprintf("bill data:[%s]", string(b)))

		//  TODO send data

		chrome.AccessPool(CurAccount)

		chrome.LastTime -= CurAccount.IntervalTime
		time.Sleep(time.Duration(CurAccount.SleepTime) * time.Second)
	}
}
