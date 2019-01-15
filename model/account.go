package model

type Account struct {
	Id       			string	`ID`
	Name     			string	`名称`
	UserName 			string	`alipay用户名`
	Password 			string	`Alipay密码`
	FirstSpiderTime 	int64	`第一次启动爬取前多久数据`
	IntervalTime		int64	`每次爬取时间间隔`
	SleepTime			int64	`每次爬取休眠时间`
	PageAccessNum		int		`随机页面访问次数`
	PageAccessWaitTime	int64	`随机页面访问等待时间`
	Status 				int		`爬取状态`
	IsRun  				bool	`是否爬取`
}

var AccessList []Account
var CurAccount *Account

func init() {
	AccessList = GetAccessList()
}

func GetAccessList() []Account {
	return []Account{
		{
			Id:       "t1",
			Name:     "test1",
			UserName: "test1",
			Password: "test1",
			FirstSpiderTime: 3*86400,
			IntervalTime: 3600,
			SleepTime: 30,
			PageAccessNum: 2,
			PageAccessWaitTime: 5,
			IsRun: true,
		},
		{
			Id:       "t2",
			Name:     "test2",
			UserName: "test2",
			Password: "test2",
			FirstSpiderTime: 2*86400,
			IntervalTime: 3600,
			SleepTime: 30,
			PageAccessNum: 2,
			PageAccessWaitTime: 5,
			IsRun: false,
		},
	}
}

func CheckAccessIDExists(id string) bool {
	for _, item := range AccessList {
		if item.Id == id {
			CurAccount = &item
			return true
		}
	}

	return false
}
