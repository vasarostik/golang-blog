package nats

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	config_service "github.com/vasarostik/go_blog/pkg/configManager/service"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/configManager/configClient"
	"github.com/vasarostik/go_blog/pkg/utl/gorm"
	"github.com/vasarostik/go_blog/pkg/utl/logfile"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
)

func Start(cfg *config.ConfigManager) error {

	var configNATS = new(config.NATS_ms)
	var req 	   = new(config_service.Request)

	//Manager client
	configManagerClient,err := configClient.New(cfg)
	checkErr(err)

	byteConf,err := configManagerClient.GetNATSConfig(context.Background(),req)
	checkErr(err)

	err = yaml.Unmarshal(byteConf.Data,configNATS)
	checkErr(err)


	natsClient, err := nats.Connect(configNATS.NATS_Server.Addr)

	if err != nil {
		log.Fatalf("Can't connect to %s: %v\n", configNATS.NATS_Server.Addr, err)
	} else {
		log.Println("Started subscription on ", configNATS.NATS_Subscriber.Subject)
	}

	db, err := gorm.New(configNATS.NATS_Subscriber.PSN)
	checkErr(err)
	defer db.Close()

	file, err := logfile.New(configNATS.NATS_Subscriber)
	checkErr(err)
	defer file.Close()

	err = StartService(configNATS.NATS_Subscriber, natsClient, db, file)
	checkErr(err)

	r := mux.NewRouter()

	log.Printf("This is Nats Subscriber listening on endpoint "+ configNATS.NATS_Subscriber.Addr)

	if err := http.ListenAndServe(configNATS.NATS_Subscriber.Addr, r); err != nil {
		log.Fatal(err)
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		println(err.Error())
	}
}


