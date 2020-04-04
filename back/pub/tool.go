package pub

import (
	"time"
	"strconv"
	"crypto/md5"
	"io"
	"fmt"
	"os"
	"OrderMeal/data/log"
)

// 获取当前秒数
func GetSecond() string {
	sec := strconv.FormatInt(time.Now().Unix(), 10)
	return sec
}

// 获取当前月份
func GetMonth() string {
	now := time.Now()
	year, month, _ := now.Date()
	return fmt.Sprintf("%d%d", year, month)
}

// 获取当前格式化日期
func GetDate() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%d%d%d", year, month, day)
}

// 获取格式化时间
func GetTime() string {
	now := time.Now()
	date := now.Format("2006-01-02 15:07:02")
	return date
}

// md5加密字符串
func GetMd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 数字字符串 => 64位数字
func StrToNum(str string) int64 {
	num, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		log.Debug("数字字符串转64位数字失败：", err)
		num = 0
	}
	return num
}

// 判断目录是否存在，不存在则创建目录
func DirExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	return nil
}

// 创建新目录
func CreateDir(path string) {
	err := os.Mkdir(path, os.ModePerm);
	if err != nil {
		log.Error("创建目录失败：", err)
	}
}
