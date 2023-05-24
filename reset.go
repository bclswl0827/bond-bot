package main

import (
	"time"
)

func ResetFlag(config Config, flag *bool) {
	for {
		if time.Now().Hour() == config.Hour &&
			time.Now().Minute() == config.Minute &&
			time.Now().Second() == 0 {
			*flag = false
		} else {
			*flag = true
		}

		time.Sleep(time.Second)
	}
}
