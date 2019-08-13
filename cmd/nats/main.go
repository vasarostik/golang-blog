package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	nats_service "github.com/vasarostik/go_blog/pkg/nats"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/gorm"
	"github.com/vasarostik/go_blog/pkg/utl/logfile"

	"log"
	"net/http"
)

func main() {

	cfgPath := flag.String("p", "./cmd/nats/conf.local.yaml", "Path to config file")

	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	natsClient, err := nats.Connect(cfg.NATS_Server.Addr)

	if err != nil {
		log.Fatalf("Can't connect to %s: %v\n", cfg.NATS_Server.Addr, err)
	} else {
		log.Println("Started subscription on ", cfg.NATS_Subscriber.Subject)
	}

	db, err := gorm.New()
	checkErr(err)
	defer db.Close()


	file,err := logfile.New(cfg)
	checkErr(err)
	defer file.Close()

	err = nats_service.Start(cfg,natsClient,db,file)
	checkErr(err)

	r := mux.NewRouter()
	if err := http.ListenAndServe(cfg.NATS_Subscriber.Addr, r); err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}




