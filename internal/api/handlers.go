package api

import (
	"context"
	"crypto-monitor/internal/cache"
	"crypto-monitor/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPrice - обработчик для получения цены криптовалюты
func GetPrice(c *gin.Context) {
	crypto := c.Param("crypto")
	currency := c.DefaultQuery("currency", "usd")

	client := cache.GetRedisClient()
	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis клиент не инициализирован"})
		return
	}

	priceString, err := client.Get(context.Background(), fmt.Sprintf("%s_%s", crypto, currency)).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Не удалось получить цену из кэша: %v", err)})
		return
	}

	var price models.Price
	err = json.Unmarshal([]byte(priceString), &price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при десериализации данных цены: %v", err)})
		return
	}

	c.JSON(http.StatusOK, price)
}
