package alipay

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"spider/config"
)

type ChromeObj struct {
	W           selenium.WebDriver
	Caps        selenium.Capabilities
	IsFirstPage bool
	LastTime    int64
	Err         error
}

func (this *ChromeObj) String() string {
	title, _ := this.W.Title()
	url, _ := this.W.CurrentURL()
	webStr := fmt.Sprintf("w.title:[%s],w.url:[%s]", title, url)
	capStr := fmt.Sprintf("w.caps:[%v]", this.Caps)
	lastTimeStr := fmt.Sprintf("w.lasttime:[%d]", this.LastTime)
	errStr := fmt.Sprintf("w.err:[%s]", this.Err.Error())

	return fmt.Sprintf("%s | %s | %s | %s", webStr, capStr, lastTimeStr, errStr)
}


func (this *ChromeObj) Init() {
	//链接本地的浏览器 chrome
	this.Caps = selenium.Capabilities{
		"browserName": "chrome",
	}

	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	this.Caps.AddChrome(chromeCaps)
}

func (this *ChromeObj) Start() {
	// 调起chrome浏览器
	w, err := selenium.NewRemote(this.Caps, fmt.Sprintf(config.ServiceURL, config.DriverPort))
	if err != nil {
		this.Err = err
		return
	}

	this.W = w
}

func (this *ChromeObj) Close() {
	if this.W != nil {
		this.W.Quit()
	}
}
