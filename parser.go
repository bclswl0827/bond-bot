package main

import (
	"bytes"
	"encoding/json"
	"time"
)

func BondParser(data []byte) any {
	var bond any
	json.NewDecoder(&JsonpWrapper{
		Underlying: bytes.NewBuffer(data),
		Prefix:     "_",
	}).Decode(&bond)

	return bond
}

func BondFilter(data any) string {
	type Bonds struct {
		Result struct {
			Data []struct {
				Name   string `json:"SECURITY_NAME_ABBR"`
				Code   string `json:"SECURITY_CODE"`
				Date   string `json:"VALUE_DATE"`
				Rating string `json:"RATING"`
			} `json:"data"`
		} `json:"result"`
	}
	var (
		message string = ""
		bonds   Bonds
	)
	json.Unmarshal(func(data any) []byte {
		b, _ := json.Marshal(data)
		return b
	}(data), &bonds)

	for _, v := range bonds.Result.Data {
		// 匹配今天
		if v.Date == time.Now().Format("2006-01-02")+" 00:00:00" {
			message += "·" + v.Name + "（" + v.Code + " / " + v.Rating + "）\n"
		}
		// 匹配明天
		if v.Date == time.Now().Add(time.Hour*24).Format("2006-01-02")+" 00:00:00" {
			message += "·" + v.Name + "（" + v.Code + " / " + v.Rating + " / 预约）\n"
		}
		// 匹配后天
		if v.Date == time.Now().Add(time.Hour*24*2).Format("2006-01-02")+" 00:00:00" {
			message += "·" + v.Name + "（" + v.Code + " / " + v.Rating + " / 预约）\n"
		}
	}

	if len(message) == 0 {
		message = "今天没有可转债供申购或预约"
	}
	return message
}
