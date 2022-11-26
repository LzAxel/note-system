package main

import (
	"log"
	"note-system/internal/app"
	"note-system/internal/config"
	"note-system/internal/handler"
	"note-system/internal/service"
	"note-system/internal/storage"
	"note-system/internal/storage/psql"
	"note-system/pkg/jwt"
	"note-system/pkg/logging"
	"time"
)

func main() {
	log.Println("config init")
	cfg := config.GetConfig()

	log.Println("logger init")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	db, err := psql.NewPostgresStorage(psql.Config(cfg.DBConfig))
	if err != nil {
		logger.Fatalf("failed to connect to db: %v", err)
	}

	jwtManager := jwt.NewJWTManager(cfg.JWTSecret, time.Hour*24)
	storage := storage.NewStorage(logger, db)
	service := service.NewService(logger, storage, jwtManager)
	handler := handler.NewHandler(logger, service)

	a, err := app.NewApp(cfg, logger, handler)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("running app")
	a.Run()
}
