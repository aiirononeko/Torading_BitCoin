package main

import (
	"fmt"
	"log"

	"github.com/aiirononeko/gotrading/config"
	"github.com/aiirononeko/gotrading/utils"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.SecretKey)

	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
