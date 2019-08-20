package chat

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"log"
	"net/http"
	"time"
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

		err = s.Redis_AddMessage(c,msg)
		if err != nil {
			println(err)
		}

		broadcast <- msg
	}

	return nil
}


func (s *ServiceStruct) Redis_AddMessage(ctx echo.Context, in go_blog.ChatMessage) (error) {
		msg,err := json.Marshal(in)

		if err != nil{
			println(err)
		} else {
			_ = s.db.ZAdd(s.key,
				redis.Z{Score:float64(time.Now().Unix()),Member:msg})
		}

	return nil
}


func (s *ServiceStruct) Redis_GetMessages(ctx echo.Context) (go_blog.MessagesList,error) {
	var messages = new(go_blog.MessagesList)

	messageslist, err := s.db.ZRange(s.key, 0, -1).Result()

	if err != nil {
		panic(err)
	}

	messages.Messages = messageslist

	return *messages,nil
}






//go routine
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

