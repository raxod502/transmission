package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/raxod502/transmission/backend/src/api"
	"github.com/raxod502/transmission/backend/src/model"
)

type Client struct {
	conn     *websocket.Conn
	outgoing chan []byte
	incoming chan []byte
}

type ClientManager struct {
	state *model.State
	// Registered clients.
	clients map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
	incoming   chan []byte
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients:    map[*Client]bool{},
		register:   make(chan *Client),
		unregister: make(chan *Client),
		incoming:   make(chan []byte),
	}
}

func (c *ClientManager) UpdateClientsState(state *model.State) error {
	payload, err := json.Marshal(state)
	if err != nil {
		return err
	}
	for client := range c.clients {
		client.outgoing <- payload
	}
	return nil
}

func (c *ClientManager) Run() {
	for {
		select {
		case client := <-c.register:
			c.clients[client] = true
		case client := <-c.unregister:
			if _, ok := c.clients[client]; ok {
				delete(c.clients, client)
				close(client.outgoing)
			}
		case message := <-c.incoming:
			err := api.HandleEvent(message, c.state)
			if err != nil {
				fmt.Printf("error handling api event %v\n", err)
				continue
			}
			err = c.UpdateClientsState(c.state)
			if err != nil {
				fmt.Printf("error updating the state of clients: %v\n", err)
				continue
			}
		}
	}
}

func (c *Client) readPump(manager *ClientManager) {
	defer func() {
		manager.unregister <- c
		c.conn.Close()
	}()
	// TODO: Check if we want to change any of the default config values
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("unexpected websocket read error: %v\n", err)
			}
			break
		}
		manager.incoming <- message
	}
}

func (c *Client) writePump(manager *ClientManager) {
	defer c.conn.Close()
	for {
		select {
		case message, ok := <-c.outgoing:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Printf("websocket nextwriter err: %v\n", err)
				return
			}
			w.Write(message)
		}
	}
}

var upgrader = websocket.Upgrader{}

func addClient(w http.ResponseWriter, r *http.Request, manager *ClientManager) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("error upgrading conn: %v\n", err)
		return
	}
	client := &Client{conn: conn, incoming: make(chan []byte), outgoing: make(chan []byte)}
	manager.register <- client
	go client.readPump(manager)
	go client.writePump(manager)
}
