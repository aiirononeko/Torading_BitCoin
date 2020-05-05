package main

import (
	"fmt"

	"github.com/aiirononeko/gotrading/bitflyer"
	"github.com/aiirononeko/gotrading/config"
	"github.com/aiirononeko/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.SecretKey)
	fmt.Println(apiClient.GetBalance())
}
