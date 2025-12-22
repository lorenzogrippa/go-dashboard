package main

import (
	"dashboard/configuration"
	"log"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	dsn := "root:_local_root_@tcp(localhost:3306)/breeders?parseTime=true&tls=false"
	db, err := initDB(dsn)
	if err != nil {
		log.Panic(err)
	}

	testApp = application{
		App: configuration.New(db),
	}

	os.Exit(m.Run())
}
