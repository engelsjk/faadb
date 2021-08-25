package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/servers/aircraftserver"
	"github.com/engelsjk/faadb/twirpserver"
)

func main() {

	var flagPort = flag.String("p", "8082", "port")
	var flagDataPath = flag.String("d", "", "data path")
	var flagDBPath = flag.String("b", "", "database path")
	var flagReloadDB = flag.Bool("r", false, "reload database")

	flag.Parse()

	a, err := aircraftserver.NewAircraftService(*flagDataPath, *flagDBPath, *flagReloadDB)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := aircraft.NewAircraftServer(aircraftserver.NewServer(a))

	twirpserver.Start(*flagPort, a.Name, twirpHandler)
}
