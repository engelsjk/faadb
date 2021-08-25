package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/reserved"
	"github.com/engelsjk/faadb/servers/reservedserver"
	"github.com/engelsjk/faadb/twirpserver"
)

func main() {

	var flagPort = flag.String("p", "8084", "port")
	var flagDataPath = flag.String("d", "", "data path")
	var flagDBPath = flag.String("b", "", "database path")
	var flagReloadDB = flag.Bool("r", false, "reload database")
	flag.Parse()

	r, err := reservedserver.NewReserved(*flagDataPath, *flagDBPath, *flagReloadDB)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := reserved.NewReservedServer(reservedserver.NewServer(r))

	twirpserver.Start(*flagPort, r.Name, twirpHandler)
}
