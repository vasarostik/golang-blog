package transport

import (
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/api/chat"
)

// HTTP represents user http service
type HTTP struct {
	svc chat.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc chat.Service, e *echo.Echo, jwtURL echo.MiddlewareFunc) {
	h := HTTP{svc}

	e.GET("/ws",h.handleConnections,jwtURL)
}


func(h *HTTP) handleConnections(c echo.Context) error {


	err := h.svc.CreateWebSocket(c)

	if(err != nil){
		println(err)
	}

	return nil
}


