package alipay

import (
	"fmt"
	"github.com/tebeka/selenium"
	"spider/common"
	"spider/util/dates"
)

func (c *ChromeObj) InputSingleStr(element, content string) {
	ele, err := c.W.FindElement(selenium.ByCSSSelector, element)
	if err != nil {
		return
	}

	l := common.SplitStr(content)

	for _, s := range l {
		ele.SendKeys(s)
		dates.SleepSecond(1)
	}

	dates.SleepSecond(5)
}

func (c *ChromeObj) InputStr(element, content string) {
	ele, err := c.W.FindElement(selenium.ByCSSSelector, element)
	if err != nil {
		return
	}

	ele.SendKeys(content)
	dates.SleepSecond(1)
}

func (c *ChromeObj) InputsStr(element, content string, index int) {
	eles, err := c.W.FindElements(selenium.ByCSSSelector, element)
	if err != nil {
		return
	}

	eles[index].SendKeys(content)
	dates.SleepSecond(1)
}

func (c *ChromeObj) Click(element string) {
	ele, err := c.W.FindElement(selenium.ByCSSSelector, element)
	if err != nil {
		return
	}

	ele.Click()
	dates.SleepSecond(3)
}

func (c *ChromeObj) Clicks(element string, index int) {
	ele, err := c.W.FindElements(selenium.ByCSSSelector, element)
	if err != nil {
		return
	}

	ele[index].Click()
	dates.SleepSecond(3)
}

func (c *ChromeObj) CheckIsExists(v string) bool {
	ele, err := c.W.FindElement(selenium.ByCSSSelector, v)
	if err != nil {
		return false
	}

	if ok, err := ele.IsDisplayed(); !ok || err != nil {
		return false
	}

	return true
}

func (c *ChromeObj) ExecJS(v string) {
	c.W.ExecuteScript(v, nil)
}

func (c *ChromeObj) Text(v string) string {
	ele, err := c.W.FindElement(selenium.ByCSSSelector, v)
	if err != nil {
		return ""
	}

	s, err := ele.Text()
	if err != nil {
		return ""
	}

	return s
}

func (c *ChromeObj) TextE(e selenium.WebElement) string {
	s, err := e.Text()
	if err != nil {
		return ""
	}

	return s
}

func (c *ChromeObj) GetAttrE(e selenium.WebElement, attr string) string {
	s, err := e.GetAttribute(attr)
	if err != nil {
		return ""
	}

	return s
}

func (c *ChromeObj) GetAttr(v, attr string) string {
	ele, err := c.W.FindElement(selenium.ByCSSSelector, v)
	if err != nil {
		return ""
	}

	s, err := ele.GetAttribute(attr)
	if err != nil {
		return ""
	}

	return s
}

func (c *ChromeObj) setVisiable(y int) {
	c.ExecJS(fmt.Sprintf("window.scrollTo(document.body.scrollWidth,%d);", y))
}

func (c *ChromeObj) setBottom() {
	c.ExecJS("window.scrollTo(document.body.scrollWidth,document.body.scrollHeight);")
}
