package go_blog

import "github.com/labstack/echo"

// AuthToken holds authentication token details with refresh token
type AuthToken struct {
	Token        string `json:"token"`
	Expires      string `json:"expires"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshToken holds authentication token details
type RefreshToken struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

// RBACService represents role-based access control service interface
type RBACService interface {
	User(echo.Context) *AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, AccessRole, int, int) error
	IsLowerRole(echo.Context, AccessRole) error
}