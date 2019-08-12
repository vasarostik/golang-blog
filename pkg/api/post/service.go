package post

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/nats-io/go-nats"
	"github.com/vasarostik/go_blog/pkg/api/post/platform/pgsql"
	"github.com/vasarostik/go_blog/pkg/grpc/service"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, go_blog.Post) (*go_blog.Post, error)
	MyList(echo.Context, int) ([]go_blog.Post, error)
	MyListGRPC(echo.Context, int) ([]go_blog.Post, error)
	List(echo.Context) ([]go_blog.Post, error)
	View(echo.Context, int) (*go_blog.Post, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*go_blog.Post, error)
	PublishMessage(string, go_blog.PublishPostMessage) error
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer, grpcClient service.CreatePostZClient, natsClient *nats.Conn) *Post {
	return &Post{db: db, udb: udb, rbac: rbac, sec: sec, grpcClient: grpcClient, natsClient:natsClient}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer, grpcClient service.CreatePostZClient,natsClient *nats.Conn) *Post {
	return New(db, pgsql.NewPost(), rbac, sec, grpcClient, natsClient)
}

// User represents user application service
type Post struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
	grpcClient service.CreatePostZClient
	natsClient *nats.Conn
}



// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, go_blog.Post) (*go_blog.Post, error)
	View(orm.DB, int) (*go_blog.Post, error)
	MyList(orm.DB, int) ([]go_blog.Post, error)
	List(orm.DB) ([]go_blog.Post, error)
	Update(orm.DB, *go_blog.Post) error
	Delete(orm.DB, *go_blog.Post) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	EnforceUser(echo.Context, int) error
}
