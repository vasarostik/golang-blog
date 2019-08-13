package logfile

import (
	"fmt"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"os"
)

func New(cfg *config.Configuration) (*os.File, error) {

	f, err := os.OpenFile(cfg.NATS_Subscriber.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return f,nil
}
