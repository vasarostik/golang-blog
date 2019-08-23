package main

import (
	"flag"
	"github.com/vasarostik/go_blog/pkg/nats"
	"github.com/vasarostik/go_blog/pkg/utl/config"

)

func main() {

	cfgPath := flag.String("p", "./dockerfiles/configManager/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load_Manager(*cfgPath)

	checkErr(err)

	err = nats.Start(cfg)

	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}




