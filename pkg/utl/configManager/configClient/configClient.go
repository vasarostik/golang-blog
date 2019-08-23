package configClient

import (
	config_service "github.com/vasarostik/go_blog/pkg/configManager/service"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"google.golang.org/grpc"
	"log"
	"os"
)

func New(cfg *config.ConfigManager) (config_service.SendClient, error) {

	conn, err := grpc.Dial(os.Getenv("MANAGER_HOST")+cfg.Addr, grpc.WithInsecure())
	if err != nil {
		return nil,err
	}

	client := config_service.NewSendClient(conn)
	log.Printf("Connected to: Config Manager on endpoint "+os.Getenv("MANAGER_HOST")+cfg.Addr)

	return client,nil
}