# Crypto Monitor

Crypto Monitor - это приложение для мониторинга цен криптовалют (Bitcoin и Ethereum) с использованием API CoinGecko и кэшированием результатов в Redis.

## Функциональность

- Периодическое обновление цен криптовалют с использованием CoinGecko API.
- Кэширование результатов в Redis для быстрого доступа.
- HTTP API для получения текущих цен криптовалют.

## Технологии

- Go
- Gin (Web Framework)
- Redis (Кэширование)
- CoinGecko API

## Установка

### Требования

- Go 1.16*
- Redis

### Шаги установки

1. **Клонирование репозитория:**

    ```bash
    git clone https://github.com/<твой_пользователь>/crypto-monitor.git
    cd crypto-monitor
    ```

2. **Установка зависимостей:**

    ```bash
    go mod tidy
    ```

3. **Настройка конфигурации:**

   Создайте файл `config.json` в корневой директории проекта:

    ```json
    {
        "redis_addr": "localhost:6379"
    }
    ```

4. **Создание файла `.env` (по желанию):**

   Создайте файл `.env` и добавьте следующие параметры:

    ```env
    GIN_MODE=release
    ```

5. **Сборка проекта:**

    ```bash
    go build -o crypto-monitor cmd/main.go
    ```

6. **Запуск приложения:**

    ```bash
    ./crypto-monitor
    ```

## Использование

### API

- **Получить цену криптовалюты:**

    ```
    GET /price/:crypto?currency=:currency
    ```

  Пример:

    ```bash
    curl http://localhost:8080/price/bitcoin?currency=usd
    ```

### Примеры запросов

- Получить цену Bitcoin в долларах США:

    ```bash
    curl http://localhost:8080/price/bitcoin?currency=usd
    ```

- Получить цену Ethereum в евро:

    ```bash
    curl http://localhost:8080/price/ethereum?currency=eur
    ```

## Разработчики

- Илья

## Лицензия

Проект доступен под лицензией MIT.