package main

import (
	"github.com/aiirononeko/gotrading/app/controllers"
	"github.com/aiirononeko/gotrading/config"
	"github.com/aiirononeko/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
