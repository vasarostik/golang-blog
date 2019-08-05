package password

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/api/password/platform/pgsql"
	"github.com/vasarostik/go_blog/pkg/utl/model"
)

// Service represents password application interface
type Service interface {
	Change(echo.Context, int, string, string) error
}

// New creates new password application service
func New(db *pg.DB, udb UserDB, rbac RBAC, sec Securer) *Password {
	return &Password{
		db:   db,
		udb:  udb,
		rbac: rbac,
		sec:  sec,
	}
}

// Initialize initalizes password application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Password {
	return New(db, pgsql.NewUser(), rbac, sec)
}

// Password represents password application service
type Password struct {
	db   *pg.DB
	udb  UserDB
	rbac RBAC
	sec  Securer
}

// UserDB represents user repository interface
type UserDB interface {
	View(orm.DB, int) (*go_blog.User, error)
	Update(orm.DB, *go_blog.User) error
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
	HashMatchesPassword(string, string) bool
	Password(string, ...string) bool
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	EnforceUser(echo.Context, int) error
}
