package monitors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// CoinGeckoClient предоставляет методы для работы с API CoinGecko
type CoinGeckoClient struct {
	BaseURL string
}

// PriceResponse представляет ответ от API CoinGecko
type PriceResponse map[string]map[string]interface{}

// NewCoinGeckoClient создает нового клиента для CoinGecko API
func NewCoinGeckoClient() *CoinGeckoClient {
	return &CoinGeckoClient{
		BaseURL: "https://api.coingecko.com/api/v3/simple/price",
	}
}

// GetPrice получает цену криптовалюты в указанной валюте
func (c *CoinGeckoClient) GetPrice(crypto string, currency string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("%s?ids=%s&vs_currencies=%s", c.BaseURL, crypto, currency))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var priceResp PriceResponse
	err = json.Unmarshal(body, &priceResp)
	if err != nil {
		return 0, err
	}

	rawPrice, ok := priceResp[crypto][currency]
	if !ok {
		return 0, fmt.Errorf("цена для %s в %s не найдена", crypto, currency)
	}

	switch price := rawPrice.(type) {
	case float64:
		return price, nil
	case string:
		return strconv.ParseFloat(price, 64)
	default:
		return 0, fmt.Errorf("неожиданный тип для цены %s в %s", crypto, currency)
	}
}
