package rbac

import (
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/utl/model"
)

// New creates new RBAC service
func New() *Service {
	return &Service{}
}

// Service is RBAC application service
type Service struct{}

func checkBool(b bool) error {
	if b {
		return nil
	}
	return echo.ErrForbidden
}

// User returns user data stored in jwt token
func (s *Service) User(c echo.Context) *go_blog.AuthUser {
	id := c.Get("id").(int)
	user := c.Get("username").(string)
	role := c.Get("role").(go_blog.AccessRole)
	return &go_blog.AuthUser{
		ID:         id,
		Username:   user,
		Role:       role,
	}
}

// EnforceUser checks whether the request to change user data is done by the same user
func (s *Service) EnforceUser(c echo.Context, ID int) error {
	// TODO: Implement querying db and checking the requested user's company_id/location_id
	if s.isAdmin(c) {
		return nil
	}
	return checkBool(c.Get("id").(int) == ID)
}

func (s *Service) isAdmin(c echo.Context) bool {
	return !(c.Get("role").(go_blog.AccessRole) > go_blog.AdminRole)
}


// AccountCreate performs auth check when creating a new account
// Location admin cannot create accounts, needs to be fixed on EnforceLocation function
func (s *Service) AccountCreate(c echo.Context, roleID go_blog.AccessRole) error {
	return s.IsLowerRole(c, roleID)
}

// IsLowerRole checks whether the requesting user has higher role than the user it wants to change
// Used for account creation/deletion
func (s *Service) IsLowerRole(c echo.Context, r go_blog.AccessRole) error {
	return checkBool(c.Get("role").(go_blog.AccessRole) < r)
}
