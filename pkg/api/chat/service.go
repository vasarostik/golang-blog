package chat

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)

const ChatKey = "Web-Socket-Chat"

// Service represents user application interface
type Service interface {
	CreateWebSocket(echo.Context) error
	Redis_AddMessage(echo.Context, go_blog.ChatMessage) (error)
	Redis_GetMessages(echo.Context)(go_blog.MessagesList,error)
}

// Initialize initializes auth application service
func Initialize(db *redis.Client) *ServiceStruct {
	return &ServiceStruct{
		key:ChatKey,
		db:db,
	}

}


type ServiceStruct struct
{
	key string
	db *redis.Client
}




