package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	// All incoming API requests go through this channel to avoid having to lock
	// access to the state object
	incoming chan []byte
}

//NewClientManager inits an empty ClientManager
func NewClientManager() *ClientManager {
	return &ClientManager{
		state:      &model.State{},
		clients:    map[*Client]bool{},
		register:   make(chan *Client),
		unregister: make(chan *Client),
		incoming:   make(chan []byte),
	}
}

//UpdateClientsState sends the latest version of the state object to every client
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

//Run handles register and unregister events as well as incoming API requests
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

// readPump processes incoming messages for each websocket until it errors
func (c *Client) readPump(manager *ClientManager) {
	defer func() {
		manager.unregister <- c
		c.conn.Close()
	}()
	// TODO: Check if we want to change any of the default config values
	for {
		fmt.Println("Waiting for message")
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("unexpected websocket read error: %v\n", err)
			}
			break
		}
		fmt.Printf("got message: %v\n", message)
		manager.incoming <- message
	}
}

// writePump processes each client's outgoing channel and sends these messages
// over the websocket
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

//AddClient upgrades a connection to a websocket and then initializes it
func addClient(w http.ResponseWriter, r *http.Request, manager *ClientManager) {
	fmt.Println("connection attempted")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("error upgrading conn: %v\n", err)
		return
	}
	client := &Client{conn: conn, outgoing: make(chan []byte)}
	fmt.Println("made client")
	manager.register <- client
	fmt.Println("registered client")
	go client.readPump(manager)
	fmt.Println("reading client")
	go client.writePump(manager)
	fmt.Println("writing client")
}

func Setup(r *mux.Router) {
	clientManager := NewClientManager()
	go clientManager.Run()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		addClient(w, r, clientManager)
	})
}
