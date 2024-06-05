package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// Hub хранит активные подключения и управляет передачей сообщений между ними.
type Hub struct {
	Clients    map[*Client]bool // Клиенты, подключенные к хабу
	Broadcast  chan []byte      // Канал для рассылки сообщений
	register   chan *Client     // Канал для регистрации новых клиентов
	unregister chan *Client     // Канал для удаления клиентов
}

// Новый экземпляр Hub
func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),      // Создание канала для широковещательной рассылки сообщений
		register:   make(chan *Client),     // Создание канала для регистрации новых клиентов
		unregister: make(chan *Client),     // Создание канала для удаления клиентов
		Clients:    make(map[*Client]bool), // Создание карты для хранения зарегистрированных клиентов
	}
}

// Метод Run запускает хаб и обрабатывает входящие запросы
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: // Регистрация нового клиента
			h.Clients[client] = true
		case client := <-h.unregister: // Удаление клиента
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast: // Рассылка сообщения всем клиентам
			for client := range h.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

// Объект для обновления соединений WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Структура Client представляет подключение WebSocket клиента и содержит поля
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

// ServeWs обрабатывает HTTP запрос на обновление до WebSocket и создает Client
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Ошибка при обновлении соединения: %v", err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
