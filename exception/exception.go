package exception

import (
	"fmt"
	"os"
	"runtime"
	"spider/config"
	"spider/util/logger"
	"spider/util/mail"
)

func Handle(isExit bool) {
	err := recover()
	if err != nil {
		buf := make([]byte, 2048)
		runtime.Stack(buf, true)
		errMsg := fmt.Sprintf("\n%s", buf)

		logger.Info(err)
		logger.Info(errMsg)

		var sErr string
		switch errObj := err.(type) {
		case error:
			sErr = errObj.Error()
		case string:
			sErr = errObj
		default:
			sErr = errMsg
		}

		mail.SendSingle(config.AdminMailer, config.ProNAME+" SYSTEM NOTICE", sErr)

		if isExit {
			//common.UnLock()
			os.Exit(1)
		}
	}

	logger.Info(fmt.Sprintf("spider is end"))
}
