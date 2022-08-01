package main

import (
	"time"
)

// 协程中定时重置 Flag
func ResetFlag() {
	// 创建一个循环
	for {
		if time.Now().Hour() == config.Hour &&
			time.Now().Minute() == config.Minute &&
			time.Now().Second() == 0 {
			Sent = false
		} else {
			Sent = true
		}
		time.Sleep(time.Second)
	}
}
