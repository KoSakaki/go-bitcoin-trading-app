package controllers

import (

	"gotrading/app/models"
	"gotrading/bitflyer"
	"gotrading/config"
	"log"
)

func StreamIngestionData() {
	var tickerChannel = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	go func() {
		for ticker := range tickerChannel {
			log.Printf("action=StreanIgestionDate, %v", ticker)
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated {
					// TODO
					// fmt.Println("Save")
				}
			}
		}
	}()

}
