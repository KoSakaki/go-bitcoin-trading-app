package main

import (
	"fmt"
	"log"
	"time"

	"example.com/bitflyer"
	"example.com/config"
	"example.com/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)

	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	tickerChannel := make(chan bitflyer.Ticker)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)

	for ticker := range tickerChannel {
		fmt.Println(ticker)
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Second))
		fmt.Println(ticker.TruncateDateTime(time.Minute))
		fmt.Println(ticker.TruncateDateTime(time.Hour))
	}

	order := &bitflyer.Order{
		ProductCode:     config.Config.ProductCode,
		ChildOrderType:  "MARKET",
		Side:            "BUY",
		Size:            0.001,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}
	res, err := apiClient.SendOrder(order)
	if err != nil {
		log.Println("Error: ", err)
	}
	if res.Status != 200 {
		log.Printf("status : %v, Error: %s", res.Status, res.ErrorMessage)

	}

	i := res.ChildOrderAcceptanceID
	params := map[string]string{
		"product_code":              config.Config.ProductCode,
		"child_order_acceptance_id": i,
	}
	r, _ := apiClient.ListOrder(params)
	fmt.Println(r)

}
