package password

import (
	"github.com/labstack/echo"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)

// Change changes user's password
func (p *Password) Change(c echo.Context, userID int, oldPass, newPass string) error {
	if err := p.rbac.EnforceUser(c, userID); err != nil {
		return err
	}

	u, err := p.udb.View(p.db, userID)
	if err != nil {
		return err
	}

	if !p.sec.HashMatchesPassword(u.Password, oldPass) {
		return go_blog.ErrIncorrectPassword
	}

	if !p.sec.Password(newPass, u.FirstName, u.LastName, u.Username) {
		return go_blog.ErrInsecurePassword
	}

	u.ChangePassword(p.sec.Hash(newPass))

	return p.udb.Update(p.db, u)
}
