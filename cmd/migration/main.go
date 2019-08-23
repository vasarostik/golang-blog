package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	config_service "github.com/vasarostik/go_blog/pkg/configManager/service"
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/vasarostik/go_blog/pkg/utl/configManager/configClient"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"github.com/vasarostik/go_blog/pkg/utl/postgres"
	"github.com/vasarostik/go_blog/pkg/utl/secure"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
)

func main() {

	cfgPath := flag.String("p", "./dockerfiles/configManager/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load_Manager(*cfgPath)

	checkErr(err)

	var configApi = new(config.API_ms)
	var req 	  = new(config_service.Request)

	//Manager client
	configManagerClient,err := configClient.New(cfg)
	checkErr(err)


	byteConf,err := configManagerClient.GetAPIConfig(context.Background(),req)
	checkErr(err)

	err = yaml.Unmarshal(byteConf.Data,configApi)
	checkErr(err)


	db, err := postgres.New(configApi.DB.PSN, configApi.DB.Timeout, configApi.DB.LogQueries)
	checkErr(err)

	u, err := pg.ParseURL(configApi.DB.PSN)

	if err != nil {
		println("Can`t parse connection string!")
	}

	db = pg.Connect(u)

	_, err = db.Exec("SELECT 1")
	checkErr(err)

	exists,err:= db.Model(&go_blog.Role{}).Exists()
	exists,err = db.Model(&go_blog.User{}).Exists()
	exists,err = db.Model(&go_blog.Post{}).Exists()

	fmt.Printf("DB tables exists: %t \n",exists)

	if (exists == false) {

		dbInsert := `
			INSERT INTO public.roles VALUES (100, 100, 'SUPER_ADMIN');
			INSERT INTO public.roles VALUES (110, 110, 'ADMIN');
			INSERT INTO public.roles VALUES (200, 200, 'USER');`
		queries := strings.Split(dbInsert, ";")

		createSchema(db, &go_blog.Role{}, &go_blog.User{}, &go_blog.Post{})

		for _, v := range queries[0 : len(queries)-1] {
			_, err := db.Exec(v)
			checkErr(err)
		}

		sec := secure.New(1, nil)

		userInsert := `INSERT INTO public.users (id, created_at, updated_at, first_name, last_name, username, password, active, role_id) VALUES (1, now(),now(),'Admin', 'Admin', 'admin', '%s', true, 100);`
		_, err = db.Exec(fmt.Sprintf(userInsert, sec.Hash("admin")))
		checkErr(err)
		println("Migration completed successfully!")
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createSchema(db *pg.DB, models ...interface{}) {
	for _, model := range models {
		checkErr(db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		}))
	}
}
