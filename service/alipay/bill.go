package alipay

import (
	"fmt"
	"github.com/tebeka/selenium"
	"spider/util/dates"
	"strings"
	"spider/util/logger"
)

const BillURL = "https://mbillexprod.alipay.com/enterprise/accountDetail.htm"

func (c *ChromeObj) GetBill() []map[string]string {
	c.InitBillSearchCond()

	isRun := true
	billList := make([]map[string]string, 0)
	for isRun && c.NextPage() {
		body, err := c.W.FindElements(selenium.ByCSSSelector, ".ant-table-body table tbody tr")
		if err != nil {
			return nil
		}

		l := len(body)

		for i := 0; i < l; i++ {
			tds, err := body[i].FindElements(selenium.ByCSSSelector, "td")
			if err != nil {
				return nil
			}

			alipayTime := strings.Replace(c.TextE(tds[0]), "\n", " ", -1)

			t := dates.Str2Int(alipayTime)
			logger.Info(fmt.Sprintf("%s,%d", alipayTime, t))

			if t < c.LastTime {
				isRun = false
				break
			}

			sp, err := tds[1].FindElement(selenium.ByCSSSelector, "div div span span")
			if err != nil {
				return nil
			}

			alipayRecord := c.GetAttrE(sp, "title")
			logger.Info(fmt.Sprintf("%s", alipayRecord))

			alipayMoney := c.TextE(tds[5])
			logger.Info(fmt.Sprintf("%s", alipayMoney))

			accs, err := tds[3].FindElements(selenium.ByCSSSelector, "div div span")
			if err != nil {
				return nil
			}

			alipayAccount := c.TextE(accs[0])
			alipayName := c.TextE(accs[1])
			if alipayName == "--" {
				continue
			}

			logger.Info(fmt.Sprintf("%s,%s", alipayAccount, alipayName))

			remark, err := tds[7].FindElement(selenium.ByCSSSelector, "span")
			if err != nil {
				return nil
			}

			alipayRemark := c.TextE(remark)
			logger.Info(fmt.Sprintf("%s", alipayRemark))

			billList = append(billList, map[string]string{
				"alipayTime":    alipayTime,
				"alipayRecord":  alipayRecord,
				"alipayMoney":   alipayMoney,
				"alipayName":    alipayName,
				"alipayAccount": alipayAccount,
				"alipayRemark":  alipayRemark,
			})

			c.IsFirstPage = false
		}
	}

	return billList
}

func (c *ChromeObj) InitBillSearchCond() {
	c.W.Get(BillURL)

	c.Clicks(".quickTimeItem___r5kmW", 2)
	c.setVisiable(0)

	if !c.CheckIsExists(".ant-btn-icon-only") {
		return
	}

	c.Clicks(".ant-btn-icon-only", 2)
	c.Clicks(".ant-checkbox-input", 1)
	c.Clicks(".ant-btn-primary", 1)

	tx := c.Text(".ant-table-body table tbody")
	if tx == "" {
		return
	}

	c.IsFirstPage = true
}

func (c *ChromeObj) NextPage() bool {
	if !c.IsFirstPage {
		if c.GetAttr(".ant-pagination-next", "aria-disabled") == "true" {
			return false
		}

		c.setBottom()
		c.Click(".ant-pagination-next")
	}

	return true
}
