package alipay

import (
	"math/rand"
	"spider/model"
	"time"
)

func (c *ChromeObj) AccessPage(page string) {
	c.W.Get(page)
}

func (c *ChromeObj) AccessPool(account *model.Account) {

	for i:= 0; i < account.PageAccessNum; i++ {
		list := getPageList()
		l := len(list)
		r := rand.Intn(l)

		c.AccessPage(list[r])
		time.Sleep(time.Duration(account.PageAccessWaitTime) * time.Second)
	}

}

func getPageList() []string {
	return []string{
		"https://uemprod.alipay.com/user/ihome.htm",
		"https://mbillexprod.alipay.com/enterprise/arBillQuery.htm",
	}
}
