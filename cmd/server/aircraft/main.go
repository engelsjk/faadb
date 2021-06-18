package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/servers/aircraftserver"
	server "github.com/engelsjk/faadb/twirp-web-server"
)

func main() {

	var flagPort = flag.String("p", "8082", "port")
	var flagDataPath = flag.String("dp", "ACFTREF.txt", "data path")
	var flagDBPath = flag.String("db", "aircraft.db", "database path")
	flag.Parse()

	a, err := aircraftserver.NewAircraftService(*flagDataPath, *flagDBPath)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := aircraft.NewAircraftServer(aircraftserver.NewServer(a))

	server.Start(*flagPort, a.Name, twirpHandler)
}
