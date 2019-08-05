package user

import (
	"github.com/vasarostik/go_blog/pkg/api/post"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"time"

	"github.com/labstack/echo"
)

// New creates new user logging service
func New(svc post.Service, logger go_blog.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents post logging service
type LogService struct {
	post.Service
	logger go_blog.Logger
}

const name = "post"

// Create logging
func (ls *LogService) Create(c echo.Context, req go_blog.Post) (resp *go_blog.Post, err error) {
	defer func(begin time.Time) {

		ls.logger.Log(
			c,
			name, "Create post`s request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}


// List logging
func (ls *LogService) MyList(c echo.Context, id int, req *go_blog.Pagination) (resp []go_blog.Post, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "MyList post`s request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.MyList(c,id, req)
}

// List logging
func (ls *LogService) List(c echo.Context, req *go_blog.Pagination) (resp []go_blog.Post, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List post`s request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c, req)
}

// View logging
func (ls *LogService) View(c echo.Context, req int) (resp *go_blog.Post, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View post`s request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, req)
}

// Delete logging
func (ls *LogService) Delete(c echo.Context, req int) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete post`s request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, req)
}

// Update logging
func (ls *LogService) Update(c echo.Context, req *post.Update) (resp *go_blog.Post, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Update post`s request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, req)
}
