package chat

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"log"
	"net/http"
	)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan go_blog.ChatMessage)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *ServiceStruct) CreateWebSocket(c echo.Context) (error) {

	go handleMessages()
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()
	clients[ws] = true

	for {
		var msg go_blog.ChatMessage
		// Read in a new message as JSON and map it to a Message object

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		// send the new message to the broadcast channel
		msg.Username = c.Get("username").(string)
		broadcast <- msg
	}

	return nil
}


func handleMessages() {
	for {
		// grab next message from the broadcast channel
		msg := <-broadcast

		// send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				_ = client.Close()
				delete(clients, client)
			}
		}
	}
}

