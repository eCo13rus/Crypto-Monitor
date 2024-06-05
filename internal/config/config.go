package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	RedisAddr string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Не удалось прочитать конфигурационный файл: %v", err)
	}

	AppConfig = &Config{
		RedisAddr: viper.GetString("redis_addr"),
	}

	log.Printf("Конфигурация загружена: RedisAddr=%s", AppConfig.RedisAddr)
}
