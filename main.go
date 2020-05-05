package main

import (
	"fmt"
	"time"

	"github.com/aiirononeko/gotrading/bitflyer"
	"github.com/aiirononeko/gotrading/config"
	"github.com/aiirononeko/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.SecretKey)

	tickerChanel := make(chan bitflyer.Ticker)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChanel)
	for ticker := range tickerChanel {
		fmt.Println(ticker)
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Hour))
		fmt.Println(ticker.TruncateDateTime(time.Minute))
		fmt.Println(ticker.TruncateDateTime(time.Second))
	}
}
