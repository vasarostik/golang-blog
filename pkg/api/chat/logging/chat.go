package auth

import (
	"github.com/vasarostik/go_blog/pkg/api/chat"
	"time"

	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/utl/model"
)

// New creates new chat logging service
func New(svc chat.Service, logger go_blog.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents chat logging service
type LogService struct {
	chat.Service
	logger go_blog.Logger
}


const name = "chat"

// Chat logging
func (ls *LogService) CreateChat(c echo.Context) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Chat request", err,
			map[string]interface{}{
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return nil
}

