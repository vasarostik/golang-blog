package config_service

import (
	"context"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"gopkg.in/yaml.v2"
)

type PostResp struct {

}

func New(cfg *config.Configuration) *Server {
	return &Server{
		cfg: cfg,
	}
}

type Server struct {
	cfg *config.Configuration
}

func (s *Server) GetAPIConfig(ctx context.Context, request *Request) (*ConfStruct,error) {
	var apiConf = new(ConfStruct)

	temp, err := yaml.Marshal(s.cfg.API_ms)
	if err != nil {
		panic(err)
	}

	apiConf.Data = temp

	return apiConf,nil
}

func (s *Server) GetGRPCConfig(ctx context.Context, request *Request) (*ConfStruct,error) {
	var grpcConf = new(ConfStruct)

	temp, err := yaml.Marshal(s.cfg.GRPC_ms)
	if err != nil {
		panic(err)
	}

	grpcConf.Data = temp

	return grpcConf,nil
}

func (s *Server) GetNATSConfig(ctx context.Context, request *Request) (*ConfStruct,error) {
	var natsConf = new(ConfStruct)

	temp, err := yaml.Marshal(s.cfg.NATS_ms)
	if err != nil {
		panic(err)
	}

	natsConf.Data = temp

	return natsConf,nil
}



