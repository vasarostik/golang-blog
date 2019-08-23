package api

import (
	"context"
	"crypto/sha1"
	"github.com/nats-io/go-nats"
	"github.com/vasarostik/go_blog/pkg/api/chat"
	config_service "github.com/vasarostik/go_blog/pkg/configManager/service"
	"github.com/vasarostik/go_blog/pkg/utl/configManager/configClient"
	"github.com/vasarostik/go_blog/pkg/utl/redis"
	"github.com/vasarostik/go_blog/pkg/utl/zlog"
	"gopkg.in/yaml.v2"
	logging "log"

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
func Start(cfg *config.ConfigManager) error {
	var configApi = new(config.API_ms)
	var req 	  = new(config_service.Request)

	//Manager client
	configManagerClient,err := configClient.New(cfg)
	checkErr(err)


	byteConf,err := configManagerClient.GetAPIConfig(context.Background(),req)
	checkErr(err)

	err = yaml.Unmarshal(byteConf.Data,configApi)
	checkErr(err)


	db, err := postgres.New(configApi.DB.PSN, configApi.DB.Timeout, configApi.DB.LogQueries)
	checkErr(err)


	sec := secure.New(configApi.App.MinPasswordStr, sha1.New())
	rbac := rbac.New()
	jwt := jwt.New(configApi.JWT.Secret, configApi.JWT.SigningAlgorithm,configApi.JWT.Duration)
	log := zlog.New()
	e := server.New()

	//GRPC client
	GRPCclient,err := client.New(configApi.GRPC)
	checkErr(err)


	//NATS client
	natsClient, err := nats.Connect(configApi.NATS_Server.Addr)
	checkErr(err)
	logging.Printf("Connected to: Nats server listening on endpoint "+ configApi.NATS_Server.Addr)


	//Redis client
	dbRedisClient, err := redis.New(configApi.Redis)
	checkErr(err)
	logging.Printf("Connected to: Redis server listening on endpoint "+ configApi.Redis.Addr)

	//Swagger UI
	e.Static("/swagger",configApi.App.SwaggerUIPath)

	at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, jwt.MWFunc())

	v1 := e.Group("/v1")
	v1.Use(jwt.MWFunc())


	ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1, e)
	pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)
	pst.NewHTTP(psl.New(post.Initialize(db, rbac, sec, GRPCclient, natsClient), log), v1)
	ct.NewHTTP(csl.New(chat.Initialize(dbRedisClient),log),e,jwt.MWFuncURL(),jwt.MWFunc())
	logging.Printf("This is Main API server listening on endpoint "+ configApi.Server.Port)

	server.Start(e, &server.Config{
		Port:                configApi.Server.Port,
		ReadTimeoutSeconds:  configApi.Server.ReadTimeout,
		WriteTimeoutSeconds: configApi.Server.WriteTimeout,
		Debug:               configApi.Server.Debug,
	})

	return nil
}

func checkErr(err error) {
	if err != nil {
		println(err.Error())
	}
}
