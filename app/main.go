package main

import (
	"database/sql"
	"flag"
	"net/http"

	h "github.com/ccdb-api/app/controller"
	"github.com/ccdb-api/app/service"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db         *sql.DB
	dbName     = flag.String("db_name", "db2", "Database name; Must have an existing ccdb_dev database!!")
	dbUser     = flag.String("db_user", "root", "Name of user to use during test")
	dbPassword = flag.String("db_pass", "root", "Password for user")
)

func main() {
	var err error
	flag.Parse()
	db, err = sql.Open("mysql", *dbUser+":"+*dbPassword+"@tcp(127.0.0.1:3306)/"+*dbName+"?charset=utf8&parseTime=true&sql_mode=ansi")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	app := &service.AppServer{db}
	http.ListenAndServe(":8080", h.Handler(app))
}
