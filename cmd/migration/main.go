package main

import (
	"fmt"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"log"
	"strings"
	"github.com/vasarostik/go_blog/pkg/utl/secure"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func main() {
	dbInsert := `
	INSERT INTO public.roles VALUES (100, 100, 'SUPER_ADMIN');
	INSERT INTO public.roles VALUES (110, 110, 'ADMIN');
	INSERT INTO public.roles VALUES (120, 120, 'COMPANY_ADMIN');
	INSERT INTO public.roles VALUES (130, 130, 'LOCATION_ADMIN');
	INSERT INTO public.roles VALUES (200, 200, 'USER');`
	queries := strings.Split(dbInsert, ";")


	db := pg.Connect(&pg.Options{
		User: "postgres",
		Password: "example",
		Database: "postgres",
		Addr: "db:5432",
	})



	_, err := db.Exec("SELECT 1")
	checkErr(err)
	createSchema(db, &go_blog.Role{}, &go_blog.User{}, &go_blog.Post{})

	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		checkErr(err)
	}

	sec := secure.New(1, nil)

	userInsert := `INSERT INTO public.users (id, created_at, updated_at, first_name, last_name, username, password, active, role_id) VALUES (1, now(),now(),'Admin', 'Admin', 'admin', '%s', true, 100);`
	_, err = db.Exec(fmt.Sprintf(userInsert, sec.Hash("admin")))
	checkErr(err)
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
