package server

import (
	"github.com/vasarostik/go_blog/pkg/grpc/service"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start(v1API service.CreatePostZServer, cfg *config.GRPC) {
	lis, err := net.Listen("tcp", cfg.Addr)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	service.RegisterCreatePostZServer(grpcServer,v1API)
	log.Printf("This is gRPC listening on endpoint "+cfg.Addr)
	_=grpcServer.Serve(lis)
}