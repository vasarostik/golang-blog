package user

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/api/user/platform/pgsql"
	"github.com/vasarostik/go_blog/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, go_blog.User) (*go_blog.User, error)
	List(echo.Context, *go_blog.Pagination) ([]go_blog.User, error)
	View(echo.Context, int) (*go_blog.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*go_blog.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.NewUser(), rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, go_blog.User) (*go_blog.User, error)
	View(orm.DB, int) (*go_blog.User, error)
	List(orm.DB, *go_blog.ListQuery, *go_blog.Pagination) ([]go_blog.User, error)
	Update(orm.DB, *go_blog.User) error
	Delete(orm.DB, *go_blog.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *go_blog.AuthUser
	EnforceUser(echo.Context, int) error
}
