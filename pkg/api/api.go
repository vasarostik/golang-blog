package api

import (
	"crypto/sha1"
	"github.com/nats-io/go-nats"
	"github.com/vasarostik/go_blog/pkg/api/chat"
	"github.com/vasarostik/go_blog/pkg/utl/zlog"

	"github.com/vasarostik/go_blog/pkg/api/auth"
	al "github.com/vasarostik/go_blog/pkg/api/auth/logging"
	at "github.com/vasarostik/go_blog/pkg/api/auth/transport"
	csl "github.com/vasarostik/go_blog/pkg/api/chat/logging"
	ct "github.com/vasarostik/go_blog/pkg/api/chat/transport"
	"github.com/vasarostik/go_blog/pkg/api/password"
	pl "github.com/vasarostik/go_blog/pkg/api/password/logging"
	pt "github.com/vasarostik/go_blog/pkg/api/password/transport"
	"github.com/vasarostik/go_blog/pkg/api/post"
	psl "github.com/vasarostik/go_blog/pkg/api/post/logging"
	pst "github.com/vasarostik/go_blog/pkg/api/post/transport"
	"github.com/vasarostik/go_blog/pkg/api/user"
	ul "github.com/vasarostik/go_blog/pkg/api/user/logging"
	ut "github.com/vasarostik/go_blog/pkg/api/user/transport"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/grpc/client"
	"github.com/vasarostik/go_blog/pkg/utl/middleware/jwt"
	"github.com/vasarostik/go_blog/pkg/utl/postgres"
	"github.com/vasarostik/go_blog/pkg/utl/rbac"
	"github.com/vasarostik/go_blog/pkg/utl/secure"
	"github.com/vasarostik/go_blog/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	db, err := postgres.New(cfg.DB.PSN, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.New()
	jwt := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	log := zlog.New()
	e := server.New()

	//GRPC client
	GRPCclient,err := client.New(cfg.GRPC)

	if err != nil {
		return err
	}

	//NATS client
	natsClient, err := nats.Connect(cfg.NATS_Server.Addr)

	if err != nil {
		println(err)
	}


	//e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, jwt.MWFunc())

	v1 := e.Group("/v1")
	v1.Use(jwt.MWFunc())


	ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1, e)
	pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)
	pst.NewHTTP(psl.New(post.Initialize(db, rbac, sec, GRPCclient, natsClient), log), v1)
	ct.NewHTTP(csl.New(chat.Initialize(),log),e,jwt.MWFuncURL())

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
