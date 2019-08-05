package query

import (
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/utl/model"
)

// List prepares data for list queries
func List(u *go_blog.AuthUser) (*go_blog.ListQuery, error) {
	switch true {
	case u.Role <= go_blog.AdminRole: // user is SuperAdmin or Admin
		return &go_blog.ListQuery{Query: "deleted_at is null"}, nil
	default:
		return nil, echo.ErrForbidden
	}
}

