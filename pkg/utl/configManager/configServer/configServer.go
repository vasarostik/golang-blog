package configServer

import (
	config_service "github.com/vasarostik/go_blog/pkg/configManager/service"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func Start(v1API config_service.SendServer, cfg *config.ConfigManager) {
	lis, err := net.Listen("tcp", os.Getenv("MANAGER_HOST")+cfg.Addr)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	config_service.RegisterSendServer(grpcServer,v1API)
	log.Printf("This is ConfigServer listening on endpoint "+ os.Getenv("MANAGER_HOST")+ cfg.Addr)
	_=grpcServer.Serve(lis)
}

