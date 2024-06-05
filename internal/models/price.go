package models

import (
	"time"
)

// Price - структура для хранения информации о цене криптовалюты
type Price struct {
	Crypto    string    `json:"crypto"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}
