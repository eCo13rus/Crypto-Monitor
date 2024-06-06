package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	_ "path/filepath"
)

type Config struct {
	RedisAddr           string
	SupportedCryptos    []string
	SupportedCurrencies []string
}

var AppConfig *Config

func LoadConfig() {
	// Проверяем, установлен ли путь к конфигурационному файлу через переменную окружения
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "." // путь по умолчанию для локальной разработки
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Не удалось прочитать конфигурационный файл: %v", err)
	}
	AppConfig = &Config{
		RedisAddr:           viper.GetString("redis_addr"),
		SupportedCryptos:    viper.GetStringSlice("supported_cryptos"),
		SupportedCurrencies: viper.GetStringSlice("supported_currencies"),
	}
	log.Printf("Конфигурация загружена: RedisAddr=%s", AppConfig.RedisAddr)
}
