package util

import "time"

//获取当前的Unix时间戳
func GetCurrentTime() int64 {
	return time.Now().Unix()
}

//获取当前的毫秒级时间戳
func GetCurrentMilliTime() int64 {
	return time.Now().UnixNano() / 1000000
}
