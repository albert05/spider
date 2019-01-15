package common

import (
	"fmt"
	"spider/util/dates"
	"strings"
	"time"
)

const DefaultSleepTIME = time.Millisecond * 10

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if pos > len(runes) {
		return ""
	}
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func Wait(timePoint float64) {
	currTime := dates.TimeInt2float(dates.CurrentMicro())
	fmt.Println(currTime, timePoint)

	for currTime < timePoint {
		time.Sleep(DefaultSleepTIME)

		currTime = dates.TimeInt2float(dates.CurrentMicro())
	}
}

func SplitStr(s string) []string {
	return strings.SplitN(s, "", len(s))
}
