package main

import (
	"fmt"
	"spider/config"
	. "spider/util/cmd"
)

func main() {
	fmt.Println(IC.GetPidByServerName(fmt.Sprintf("%s  -id %s", config.ScriptRunName, "t1")))
}

