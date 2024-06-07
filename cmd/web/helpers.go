package main

import (
	"database/sql"
	"log"
	"os"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func initializeLogFiles() (*log.Logger, *log.Logger) {
	//to log your output to standard streams and redirect the output to a file at runtime.
	f, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//defer f.Close()

	//log.New() to create a logger for writing information messages.
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	el, err := os.OpenFile("./error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//defer el.Close()

	//logger for writing error messages in the same way,
	errorLog := log.New(el, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog

}
