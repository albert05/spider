package alipay

import (
	"spider/util/dates"
	"strings"
	"time"
	"spider/model"
)

const loginURL = "https://auth.alipay.com/login/index.htm" //"https://www.alipay.com"
const TryLoginNum = 3
const TryLoginTime = 5

func (c *ChromeObj) Login(username, password string) bool {
	c.W.Get(loginURL)

	c.Clicks("#J-loginMethod-tabs li", 1)

	c.Click("#J-login-btn")

	c.InputSingleStr("#J-input-user", username)

	c.InputSingleStr("#password_container input", password)

	c.Click("#J-login-btn")

	dates.SleepSecond(5)

	if !c.IsLogin() {
		return false
	}

	return true
}

func (c *ChromeObj) MLogin(account *model.Account) bool {
	for i := 0; !c.IsLogin() && i < TryLoginNum; i++ {
		c.Login(account.UserName, account.Password)

		time.Sleep(TryLoginTime * time.Second)
	}

	return c.IsLogin()
}

func (c *ChromeObj) IsLogin() bool {
	url, err := c.W.CurrentURL(); if err != nil || strings.Contains(url, "login/index.htm") || url == "data:," {
		return false
	}

	return true
}
