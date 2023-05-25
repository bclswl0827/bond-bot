package main

import (
	"time"
)

func ResetFlag(config Config, timzone *time.Location, flag *bool) {
	for {
		var (
			hour   = time.Now().In(timzone).Hour()
			minute = time.Now().In(timzone).Minute()
			second = time.Now().In(timzone).Second()
		)

		if hour == config.Hour &&
			minute == config.Minute &&
			second == 0 {
			*flag = false
		} else {
			*flag = true
		}

		time.Sleep(time.Second)
	}
}
