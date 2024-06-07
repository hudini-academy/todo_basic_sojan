package main

import (
	"flag"
	"log"
	"net/http"
	"todomysql/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	todos    *mysql.TodosModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@/todo?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog, errorLog := initializeLogFiles()

	// Call the openDB function to initialize a connection pool.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
		log.Println(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		todos:    &mysql.TodosModel{DB: db},
	}

	// initilaize a new http.sevrer struct, setting Addr and handler the same old but custom errorlog
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	err = srv.ListenAndServe() //calling listenAndServe() method on new http.Server struct
	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Printf("Starting server on %s", srv.Addr)
}
