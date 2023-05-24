package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func BondData() []byte {
	client := http.Client{Timeout: 10 * time.Second}
	log.Println("Getting convertible bonds data")
	for {
		resp, err := client.Get("https://datacenter-web.eastmoney.com/api/data/v1/get?callback=_&sortColumns=PUBLIC_START_DATE&sortTypes=-1&pageNumber=1&quoteType=0&reportName=RPT_BOND_CB_LIST&columns=ALL&quoteColumns=f2~01~CONVERT_STOCK_CODE~CONVERT_STOCK_PRICE,f235~10~SECURITY_CODE~TRANSFER_PRICE,f236~10~SECURITY_CODE~TRANSFER_VALUE,f2~10~SECURITY_CODE~CURRENT_BOND_PRICE,f237~10~SECURITY_CODE~TRANSFER_PREMIUM_RATIO,f239~10~SECURITY_CODE~RESALE_TRIG_PRICE,f240~10~SECURITY_CODE~REDEEM_TRIG_PRICE,f23~01~CONVERT_STOCK_CODE~PBV_RATIO")
		if err != nil {
			log.Println("Retry due to network failure")
			continue
		}
		defer resp.Body.Close()

		b, _ := ioutil.ReadAll(resp.Body)
		return b
	}
}
