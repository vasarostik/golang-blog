package grpc

import (
	"context"
	config_service "github.com/vasarostik/go_blog/pkg/configManager/service"
	"github.com/vasarostik/go_blog/pkg/grpc/service"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/configManager/configClient"
	gc "github.com/vasarostik/go_blog/pkg/utl/grpc/server"
	"github.com/vasarostik/go_blog/pkg/utl/redis"
	"gopkg.in/yaml.v2"
)

func Start(cfg *config.ConfigManager) error {
	var configGRPC = new(config.GRPC_ms)
	var req 	  = new(config_service.Request)

	//Manager client
	configManagerClient,err := configClient.New(cfg)
	checkErr(err)

	byteConf,err := configManagerClient.GetAPIConfig(context.Background(),req)
	checkErr(err)

	err = yaml.Unmarshal(byteConf.Data,configGRPC)
	checkErr(err)


	dbRedisClient, err := redis.New(configGRPC.Redis)

	grpcService := service.New(dbRedisClient)

	gc.Start(grpcService, configGRPC.GRPC)

	return nil
}

func checkErr(err error) {
	if err != nil {
		println(err.Error())
	}
}
