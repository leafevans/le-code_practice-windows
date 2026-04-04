package models

import (
	"crypto/md5"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func UnixToTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		zap.L().Error("time parse error", zap.Error(err))
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDay() string {
	return time.Now().Format("2006-01-02")
}

func Md5(str string) string {
	return fmt.Sprintf("%x\n", md5.Sum([]byte(str)))
}

func Hello(in string) string {
	return in + "world"
}
