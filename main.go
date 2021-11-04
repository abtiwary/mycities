package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mycitiesdb "github.com/abtiwary/mycities/db"
	mycitiesserver "github.com/abtiwary/mycities/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	endCh := make(chan os.Signal, 1)
	signal.Notify(endCh, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	citiesDb, err := mycitiesdb.NewCitiesDatabase(
		"localhost",
		5432,
		"mycities",
		"postgres",
		"postgres",
	)
	defer citiesDb.DB.Close()
	if err != nil {
		panic("error creating a database connection")
	}

	del, _ := citiesDb.GetGeometryByCityName("delhi")
	fmt.Println(string(del))

	gServer, err := mycitiesserver.NewServer(
		"0.0.0.0",
		8615,
		citiesDb,
	)
	gServer.StartHTTPServer()
	defer gServer.StopHTTPServer()

	for {
		select {
		case <-endCh:
			log.Debugf("terminating...")
			os.Exit(-1)
		}
	}

}
