package common

import (
	"fmt"
	"os"
	"path/filepath"
	"spider/config"
	"spider/util/cmd"
	"spider/util/dates"
	"time"
)

// 检查文件或目录是否存在
// param string filename  文件或目录
// return bool
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 创建目录
// param string dirPath 目录
// return error
func MakeDir(dirPath string) error {
	if !IsExist(dirPath) {
		return os.MkdirAll(dirPath, 0755)
	}

	return nil
}

// 获取文件目录
// param string file
// return string
func GetFilePath(file string) string {
	if stat, err := os.Stat(file); err == nil && stat.IsDir() {
		return file
	}

	return filepath.Dir(file)
}

// 获取目录指定匹配模式的所有文件
// param string dir
// param string pattern 匹配模式
// return slice
func GetAllFileByPattern(dir string, pattern string) []string {
	filePattern := GetFilePath(dir) + cmd.IC.GetSeparator() + pattern
	files, err := filepath.Glob(filePattern)
	if err != nil {
		return nil
	}

	return files
}

func GetFileModTime(file string) int64 {
	f, err := os.Open(file)
	if err != nil {
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

func GetLogPath() (string, string) {
	access := config.AccessID
	if access == "" {
		access = "monitor"
	}

	logDir := fmt.Sprintf("%s/%s/%s", config.LogPATH, access, dates.NowYearMonthStr())
	logFile := dates.NowDayStr() + ".log"
	return logDir, logFile
}
