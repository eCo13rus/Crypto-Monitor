package cache

import (
	"context"
	"crypto-monitor/internal/config"
	"github.com/go-redis/redis/v8"
	"log"
)

// Объявление переменной глобального клиента Redis
var redisClient *redis.Client

// InitializeRedis инициализирует подключение к Redis
func InitializeRedis() {
	log.Println("Инициализация Redis клиента...")

	// Конфигурация Redis клиента на основе конфигурации приложения
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.AppConfig.RedisAddr,
	})

	// Пинг к Redis для проверки подключения
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	} else {
		log.Println("Подключение к Redis установлено успешно.")
	}
}

// GetRedisClient возвращает текущий экземпляр Redis клиента
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		log.Println("Redis клиент не инициализирован.")
	}
	return redisClient
}
