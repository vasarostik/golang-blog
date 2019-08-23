package main

import (
	"flag"
	config_service "github.com/vasarostik/go_blog/pkg/configManager/service"
	"github.com/vasarostik/go_blog/pkg/utl/azure"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/configManager/configServer"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {
	var conf = new(config.Configuration)

	cfgPath := flag.String("p", "./dockerfiles/configManager/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load_Manager(*cfgPath)

	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")

	configFile,err := azure.Load(accountName,accountKey)
	checkErr(err)

	println(configFile.String())

	err = yaml.Unmarshal(configFile.Bytes(), &conf)
	checkErr(err)



	grpcService := config_service.New(conf)

	configServer.Start(grpcService, cfg)

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
