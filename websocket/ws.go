package websocket

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Websocketmethods interface {
	Read()
	Write(message string)
}

type Wesocketclients struct {
	Id   string
	Conn *websocket.Conn
	send chan string
}

var upgreader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (c *Wesocketclients) Read() {
	defer func() {
		log.Println("Closing WebSocket connection for client:", c.Id)
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break // Exit the loop cleanly
		}
		log.Printf("Received message from client [%s]: %s", c.Id, string(msg))
	}
}

func (c *Wesocketclients) Write(message string) {
	c.send <- message
}

func Websockethandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	clientid := parts[2]
	conn, err := upgreader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Connection not establised", http.StatusBadRequest)
	}

	fmt.Println("New client connected ", clientid)

	client := Wesocketclients{
		Id:   clientid,
		Conn: conn,
		send: make(chan string),
	}
	go client.Read()
	go func() {
		defer func() {
			log.Println("Closing sender for client:", client.Id)
			client.Conn.Close()
		}()

		for message := range client.send {
			log.Printf("Sending message to client [%s]: %s", client.Id, message)
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Write error:", err)
				break // Exit loop on error
			}
		}
	}()

	client.Write("welcome" + clientid)

}
