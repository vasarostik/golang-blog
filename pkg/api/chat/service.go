package chat

import (
	"github.com/labstack/echo"
)

// Service represents user application interface
type Service interface {
	CreateWebSocket(echo.Context) error
}

// Initialize initializes auth application service
func Initialize() *ServiceStruct {
	return &ServiceStruct{
	}
}


type ServiceStruct struct {
}


