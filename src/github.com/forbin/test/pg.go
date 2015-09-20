package main

import (
	"github.com/forbin"
	"os"
	"log"
	"time"
)

func main() {
	var conninfo string = "dbname=template1 sslmode=disable"
	srv := forbin.NewPGCommand("postgres", conninfo)
	if srv == nil {
		log.Fatalf("Err NewPGCommand")
		os.Exit(-1)
	}

	err := srv.InitDatabase()
	if err != nil {
		log.Fatalf("Err InitDatabase: %s", err.Error())
		os.Exit(-1)
	}
	log.Println("InitDatabase OK")

	err = srv.Startup()
	if err != nil {
		log.Fatalf("Err Startup %s", err.Error())
		os.Exit(-1)
	}
	log.Println("Startup DB OK")
	time.Sleep(3 * time.Second)

	err = srv.HealCheck()
	if err != nil {
		log.Fatalf("Err HealCheck: %s", err.Error())
		os.Exit(-1)
	}
	log.Println("HealCheck DB OK")

	status, err := srv.LogStatus()
	if err != nil {
		log.Fatalf("Err LogStatus fail: %s", err.Error())
		os.Exit(-1)
	}
	log.Printf("LogStatus: %s", status)

	err = srv.Graceful()
	if err != nil {
		log.Fatalf("Err Graceful: %s", err.Error())
		os.Exit(-1)
	}
	log.Println("Graceful DB OK")

	err = srv.Shutdown()
	if err != nil {
		log.Fatalf("Err Shutdown: %s", err.Error())
		os.Exit(-1)
	}
	log.Println("Shutdown DB OK")

	time.Sleep(1 * time.Second)
	log.Println("End")
}
