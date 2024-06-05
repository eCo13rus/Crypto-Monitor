package main

import (
	"crypto-monitor/internal/api"
	"crypto-monitor/internal/cache"
	"crypto-monitor/internal/config"
	"crypto-monitor/internal/monitors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	DefaultPort  = "8080"
	FallbackPort = "8081"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// Определяем режим работы Gin
	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	// Загрузка конфигурации
	config.LoadConfig()
	log.Println("Конфигурация успешно загружена.")

	// Инициализация Redis
	cache.InitializeRedis()
	log.Println("Redis клиент успешно инициализирован.")

	// Запуск обновляющего компонента для обновления цен криптовалют
	updater := monitors.NewUpdater()
	go updater.Start()

	// Инициализация маршрутов
	router := gin.Default()
	api.RegisterRoutes(router)

	log.Println("Сервер запущен на порту", DefaultPort)

	// Запуск HTTP сервера
	err = http.ListenAndServe(fmt.Sprintf(":%s", DefaultPort), router)
	if err != nil {
		log.Printf("Ошибка запуска на порту %s: %v", DefaultPort, err)
		log.Println("Попытка запуска на порту", FallbackPort)
		err = http.ListenAndServe(fmt.Sprintf(":%s", FallbackPort), router)
		if err != nil {
			log.Fatalf("Ошибка при запуске сервера на резервном порту %s: %v", FallbackPort, err)
		}
	}
}
