package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppConfig struct {
		LogLevel string `env:"LOG-LEVEL" env-default:"trace"`
	}
	HttpConfig struct {
		UserPort     string `env:"HTTP-USER-PORT" env-required:"true"`
		AdminPort    string `env:"HTTP-ADMIN-PORT" env-required:"true"`
		Host         string `env:"HTTP-HOST" env-required:"true"`
		ReadTimeout  int    `env:"HTTP-READTIMEOUT" env-required:"true"`
		WriteTimeout int    `env:"HTTP-WRITETIMEOUT" env-required:"true"`
	}
	Posgres struct {
		Host     string `env:"POSTGRES-HOST" env-required:"true"`
		Port     string `env:"POSTGRES-PORT" env-required:"true"`
		User     string `env:"POSTGRES-USER" env-required:"true"`
		DbName   string `env:"POSTGRES-DBNAME" env-required:"true"`
		Password string `env:"POSTGRES-PASSWORD" env-required:"true"`
		SSLMode  string `env:"POSTGRES-SSLMODE" env-required:"true"`
	}
}

var (
	config *Config
	once   sync.Once
)

func InitConfig() *Config {
	log.Println("init config")
	once.Do(func() {
		config = &Config{}

		err := cleanenv.ReadEnv(config)
		if err != nil {
			helpText := "Simple MarketPlace HTTP Server API"
			help, _ := cleanenv.GetDescription(config, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return config
}
