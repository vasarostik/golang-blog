package pgsql

import (
	"strings"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)

// NewUser returns a new user database instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// Create creates a new user on database
func (u *User) Create(db orm.DB, usr go_blog.User) (*go_blog.User, error) {
	var user = new(go_blog.User)
	err := db.Model(user).Where("lower(username) = ? and deleted_at is null",
		strings.ToLower(usr.Username)).Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, go_blog.ErrUsernameAlreadyExists

	}

	if err := db.Insert(&usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

// View returns single user by ID
func (u *User) View(db orm.DB, id int) (*go_blog.User, error) {
	var user = new(go_blog.User)
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(user, sql, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update updates user's contact info
func (u *User) Update(db orm.DB, user *go_blog.User) error {
	_, err := db.Model(user).WherePK().UpdateNotZero()
	return err
}

// List returns list of all users
func (u *User) List(db orm.DB, qp *go_blog.ListQuery) ([]go_blog.User, error) {
	var users []go_blog.User
	q := db.Model(&users).Column("user.*", "Role").Order("user.id desc")
	if qp != nil {
		q.Where(qp.Query)
	}
	if err := q.Select(); err != nil {
		return nil, err
	}
	return users, nil
}

// Delete sets deleted_at for a user
func (u *User) Delete(db orm.DB, user *go_blog.User) error {
	return db.Delete(user)
}
