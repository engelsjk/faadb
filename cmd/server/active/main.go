package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/active"
	"github.com/engelsjk/faadb/servers/activeserver"
	"github.com/engelsjk/faadb/twirpserver"
)

func main() {

	var flagPort = flag.String("p", "8081", "port")
	var flagDataPath = flag.String("d", "", "data path")
	var flagDBPath = flag.String("b", "", "database path")
	var flagReloadDB = flag.Bool("r", false, "reload database")
	flag.Parse()

	m, err := activeserver.NewActiveService(*flagDataPath, *flagDBPath, *flagReloadDB)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := active.NewActiveServer(activeserver.NewServer(m))

	twirpserver.Start(*flagPort, m.Name, twirpHandler)
}
