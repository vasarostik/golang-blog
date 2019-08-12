package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	nats_service "github.com/vasarostik/go_blog/pkg/nats"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/gorm"

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
	}

	log.Println("Started subscription on ", cfg.NATS_Subscriber.Subject)

	db, err := gorm.New()
	checkErr(err)

	err = nats_service.Start(cfg,natsClient,db)
	checkErr(err)


	r := mux.NewRouter()
	if err := http.ListenAndServe(cfg.NATS_Subscriber.Addr, r); err != nil {
		log.Fatal(err)
	}

	defer db.Close()


	//db.Where("id = ?",117).Find(&post)




}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}




