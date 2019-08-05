// Package user contains user application services
package user

import (
	"github.com/labstack/echo"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"github.com/vasarostik/go_blog/pkg/utl/query"
)

// Create creates a new user account
func (u *User) Create(c echo.Context, req go_blog.User) (*go_blog.User, error) {
	if err := u.rbac.AccountCreate(c, req.RoleID); err != nil {
		return nil, err
	}
	req.Password = u.sec.Hash(req.Password)
	return u.udb.Create(u.db, req)
}

// List returns list of users
func (u *User) List(c echo.Context, p *go_blog.Pagination) ([]go_blog.User, error) {
	au := u.rbac.User(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return u.udb.List(u.db, q, p)
}

// View returns single user
func (u *User) View(c echo.Context, id int) (*go_blog.User, error) {
	if err := u.rbac.EnforceUser(c, id); err != nil {
		return nil, err
	}
	return u.udb.View(u.db, id)
}

// Delete deletes a user
func (u *User) Delete(c echo.Context, id int) error {
	user, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}
	if err := u.rbac.IsLowerRole(c, user.Role.AccessLevel); err != nil {
		return err
	}
	return u.udb.Delete(u.db, user)
}

// Update contains user's information used for updating
type Update struct {
	ID        int
	FirstName string
	LastName  string
}

// Update updates user's contact information
func (u *User) Update(c echo.Context, r *Update) (*go_blog.User, error) {
	if err := u.rbac.EnforceUser(c, r.ID); err != nil {
		return nil, err
	}

	if err := u.udb.Update(u.db, &go_blog.User{
		Base:      go_blog.Base{ID: r.ID},
		FirstName: r.FirstName,
		LastName:  r.LastName,
	}); err != nil {
		return nil, err
	}

	return u.udb.View(u.db, r.ID)
}
