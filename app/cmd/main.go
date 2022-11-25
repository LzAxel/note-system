package main

import (
	"log"
	"note-system/internal/app"
	"note-system/internal/config"
	"note-system/pkg/logging"
)

func main() {
	log.Println("config init")
	cfg := config.GetConfig()

	log.Println("logger init")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	logger.Info("sds")
	a, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("running app")
	a.Run()
}
