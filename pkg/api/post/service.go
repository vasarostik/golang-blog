package post

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/api/post/platform/pgsql"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, go_blog.Post) (*go_blog.Post, error)
	MyList(echo.Context, int, *go_blog.Pagination) ([]go_blog.Post, error)
	List(echo.Context, *go_blog.Pagination) ([]go_blog.Post, error)
	View(echo.Context, int) (*go_blog.Post, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*go_blog.Post, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *Post {
	return &Post{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Post {
	return New(db, pgsql.NewPost(), rbac, sec)
}

// User represents user application service
type Post struct {
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
	Create(orm.DB, go_blog.Post) (*go_blog.Post, error)
	View(orm.DB, int) (*go_blog.Post, error)
	MyList(orm.DB, int, *go_blog.Pagination) ([]go_blog.Post, error)
	List(orm.DB, *go_blog.Pagination) ([]go_blog.Post, error)
	Update(orm.DB, *go_blog.Post) error
	Delete(orm.DB, *go_blog.Post) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	EnforceUser(echo.Context, int) error
	IsLowerRole(echo.Context, go_blog.AccessRole) error
}
