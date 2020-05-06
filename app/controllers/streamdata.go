package controllers

import (
	"log"

	"github.com/aiirononeko/gotrading/app/models"
	"github.com/aiirononeko/gotrading/bitflyer"
	"github.com/aiirononeko/gotrading/config"
)

func StreamIngestionData() {
	var tickerChannl = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.SecretKey)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)
	go func() {
		for ticker := range tickerChannl {
			log.Printf("action=StreamIngestionData, %v", ticker)
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated == true && duration == config.Config.TradeDurations {
					// TODO
				}
			}
		}
	}()
}
