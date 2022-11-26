package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug   bool   `env:"IS_DEBUG" env-default:"false"`
	JWTSecret string `env:"JWT_SECRET" env-required:"true"`
	Listen    struct {
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"8080"`
	}
	AppConfig struct {
		LogLevel  string `env:"LOG_LEVEL" env-default:"info"`
		AdminUser struct {
			Login    string `env:"ADMIN_LOGIN" env-required:"true"`
			Password string `env:"ADMIN_PASS" env-required:"true"`
		}
	}
	DBConfig struct {
		Host     string `env:"DB_HOST" env-required:"true"`
		Port     string `env:"DB_PORT" env-required:"true"`
		Username string `env:"DB_USERNAME" env-required:"true"`
		Password string `env:"DB_PASS" env-required:"true"`
		DBName   string `env:"DB_NAME" env-required:"true"`
		SSLMode  string `env:"DB_SSLMODE" env-required:"true"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("collecting config")

		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			configHeaderText := "Note System"
			helpText, _ := cleanenv.GetDescription(instance, &configHeaderText)
			log.Println(helpText)
			log.Fatal(err)
		}
	})
	return instance
}
