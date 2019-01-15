package common

import (
	"fmt"
	"os"
	"spider/config"
	"spider/util/cmd"
)

func Lock() bool {
	path := cmd.IC.GetTmpPath()
	err := MakeDir(path)
	if err != nil {
		panic(err)
	}

	fileName := getLockName()
	if IsExist(fileName) {
		fmt.Println(path + " is running")
		return false
	}

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer f.Close()
	return true
}

func UnLock() bool {
	fileName := getLockName()
	if !IsExist(fileName) {
		fmt.Println(fileName + " is not exists")
		return false
	}

	if err := os.RemoveAll(fileName); err != nil {
		fmt.Println(fileName + " unlock failed")
		return false
	}

	return true
}

func getLockName() string {
	return cmd.IC.GetTmpPath() + fmt.Sprintf(os.Args[0]+"_%d.lock", config.AccessID)
}
