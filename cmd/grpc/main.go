package main

import (
	"flag"
	"github.com/vasarostik/go_blog/pkg/grpc/service"

	"github.com/vasarostik/go_blog/pkg/utl/config"
	gc "github.com/vasarostik/go_blog/pkg/utl/grpc/server"
	//"fmt"
	//"github.com/vasarostik/go_blog/pkg/grpc"
	//"os"
	//"github.com/vasarostik/go_blog/pkg/api"
	"github.com/vasarostik/go_blog/pkg/utl/redis"
)

func main() {

	cfgPath := flag.String("p", "./dockerfiles/grpc/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	dbRedisClient, err := redis.New(cfg.Redis)

	grpcService := service.New(dbRedisClient)

	gc.Start(grpcService, cfg.GRPC)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
