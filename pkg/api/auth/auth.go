package auth

import (
	"github.com/vasarostik/go_blog/pkg/utl/model"

	"github.com/labstack/echo"
)

// Authenticate tries to authenticate the user provided by username and password
func (a *Auth) Authenticate(c echo.Context, user, pass string) (*go_blog.AuthToken, error) {
	u, err := a.udb.FindByUsername(a.db, user)
	if err != nil {
		return nil, err
	}

	if !a.sec.HashMatchesPassword(u.Password, pass) {
		return nil, go_blog.ErrInvalidCredentials
	}

	token, expire, err := a.tg.GenerateToken(u)
	if err != nil {
		return nil, go_blog.ErrUnauthorized
	}

	u.UpdateLastLogin(a.sec.Token(token))

	if err := a.udb.Update(a.db, u); err != nil {
		return nil, err
	}

	return &go_blog.AuthToken{Token: token, Expires: expire, RefreshToken: u.Token}, nil
}

// Refresh refreshes jwt token and puts new claims inside
func (a *Auth) Refresh(c echo.Context, token string) (*go_blog.RefreshToken, error) {
	user, err := a.udb.FindByToken(a.db, token)
	if err != nil {
		return nil, err
	}
	token, expire, err := a.tg.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &go_blog.RefreshToken{Token: token, Expires: expire}, nil
}

// Me returns info about currently logged user
func (a *Auth) Me(c echo.Context) (*go_blog.User, error) {
	au := a.rbac.User(c)
	return a.udb.View(a.db, au.ID)
}
