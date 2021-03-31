package ws

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// room id

type Client struct {
	Manager *ClientManager
	Conn    *websocket.Conn
	Send    chan []byte
	Id      string // 客户端过来创建的id 可能是时间戳 也可以是一个结构体
}

//
func (c *Client) Read() {
	defer func() {
		c.Manager.Unregister <- c
		_ = c.Conn.Close()
	}()
	//
	//c.conn.SetReadLimit(maxMessageSize)
	//c.conn.SetReadDeadline(time.Now().Add(pongWait))
	//c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		// msgType
		// media
		_, msg, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		c.Manager.Broadcast <- msg

	}
}

//
func (c *Client) Write() {
	for msg := range c.Send {
		_ = c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
	defer c.Conn.Close()
}

type ClientManager struct {
	Clients    map[*Client]bool // client address
	Broadcast  chan []byte      // 群发
	Register   chan *Client
	Unregister chan *Client
}

//https://github.com/zeromicro/zero-examples/blob/main/chat/internal/client.go
// http://www.topgoer.com/%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B/WebSocket%E7%BC%96%E7%A8%8B.html
func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (m *ClientManager) Run() {
	for {
		select {
		// register
		case c := <-m.Register:
			m.Clients[c] = true
		case c := <-m.Unregister:
			if _, ok := m.Clients[c]; ok {
				delete(m.Clients, c)
				close(c.Send) // close channel
			}
		case data := <-m.Broadcast:
			for client, _ := range m.Clients {
				select {
				case client.Send <- data:
				default:
					// error delete client
					delete(m.Clients, client)
					close(client.Send)
				}
			}

		}
	}
}
