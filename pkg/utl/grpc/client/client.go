package client

import (
	"github.com/vasarostik/go_blog/pkg/grpc/service"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"google.golang.org/grpc"
	"log"
)

func New(cfg *config.GRPC) (service.CreatePostZClient, error) {

	conn, err := grpc.Dial(cfg.Addr, grpc.WithInsecure())
	if err != nil {
		return nil,err
	}

	client := service.NewCreatePostZClient(conn)
	log.Printf("Connected to: gRPC listening on endpoint "+cfg.Addr)

	return client,nil
}