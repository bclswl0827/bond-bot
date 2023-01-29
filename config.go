package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token  string `json:"token"`
	ChatId int64  `json:"chat_id"`
	Hour   int    `json:"hour"`
	Minute int    `json:"minute"`
	Proxy  string `json:"proxy"`
}

func ReadConfig(path string) {
	if json.Unmarshal(func(path string) []byte {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalln("Failed to read config file")
		}
		return b
	}(path), &config) != nil {
		log.Fatalln("Failed to parse config file")
	}
}
