package main

import (
	"flag"
	"github.com/vasarostik/go_blog/pkg/api"
	"github.com/vasarostik/go_blog/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "./dockerfiles/configManager/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load_Manager(*cfgPath)

	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
