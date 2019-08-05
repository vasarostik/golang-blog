package pgsql

import (
	"github.com/go-pg/pg/orm"
	"github.com/vasarostik/go_blog/pkg/utl/model"
)

// NewUser returns a new user database instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u *User) View(db orm.DB, id int) (*go_blog.User, error) {
	user := &go_blog.User{Base: go_blog.Base{ID: id}}
	err := db.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates user's info
func (u *User) Update(db orm.DB, user *go_blog.User) error {
	return db.Update(user)
}
