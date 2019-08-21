package transport

import (
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/api/chat"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"net/http"
)

// HTTP represents user http service
type HTTP struct {
	svc chat.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc chat.Service, e *echo.Echo, jwtURL echo.MiddlewareFunc, jwtHeader echo.MiddlewareFunc) {
	h := HTTP{svc}

	// swagger:route GET /ws chat generateChat
	// Creates new Web Socket connection.
	// responses:
	//  200:
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	e.GET("/ws",h.handleConnections,jwtURL)

	// swagger:route GET /mesages chat generateChat
	// Creates new Web Socket connection.
	// responses:
	//  200: messagesListResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	e.GET("/messages",h.handleMessages,jwtHeader)

}

func(h *HTTP) handleConnections(c echo.Context) error {

	err := h.svc.CreateWebSocket(c)

	if(err != nil){
		println(err)
	}

	return nil
}

func(h *HTTP) handleMessages(c echo.Context) error {

	messages,err := h.svc.Redis_GetMessages(c)

	if(err != nil){
		println(err)
	}

	return c.JSON(http.StatusOK, go_blog.MessagesList{messages.Messages})
}




