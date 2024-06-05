package monitors

import (
	"crypto-monitor/internal/cache"
	"crypto-monitor/internal/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"log"
	"time"
)

// Updater отвечает за обновление цен криптовалют
type Updater struct {
	client *CoinGeckoClient
}

// NewUpdater создает новый экземпляр Updater
func NewUpdater() *Updater {
	return &Updater{
		client: NewCoinGeckoClient(),
	}
}

// Start запускает процесс обновления цен с интервалом в 60 секунд
func (u *Updater) Start() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			u.updatePrices()
		}
	}
}

// updatePrices обновляет цены криптовалют и сохраняет их в Redis
func (u *Updater) updatePrices() {
	cryptos := []string{"bitcoin", "ethereum"}
	currencies := []string{"usd", "eur", "rub"}

	for _, crypto := range cryptos {
		// Получение цены на криптовалюту
		for _, currency := range currencies {
			priceValue, err := u.client.GetPrice(crypto, currency)
			if err != nil {
				log.Printf("Ошибка получения цены для %s в %s: %v", crypto, currency, err)
				continue
			}

			// Создание объекта модели цены
			price := models.Price{
				Crypto:    crypto,
				Price:     priceValue,
				Timestamp: time.Now(),
			}

			// Сериализация модели в JSON
			priceJSON, err := json.Marshal(price)
			if err != nil {
				log.Printf("Ошибка при сериализации цены для %s в %s: %v", crypto, currency, err)
				continue
			}

			// Запись цены в Redis
			err = cache.GetRedisClient().Set(context.Background(), fmt.Sprintf("%s_%s", crypto, currency), priceJSON, 0).Err()
			if err != nil {
				log.Printf("Ошибка записи цены для %s в %s в кэш: %v", crypto, currency, err)
				continue
			}

			log.Printf("Обновлена цена для %s: %.2f %s", crypto, priceValue, currency)
		}
	}
}
