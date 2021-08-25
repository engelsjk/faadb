package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/dereg"
	"github.com/engelsjk/faadb/servers/deregserver"
	"github.com/engelsjk/faadb/twirpserver"
)

func main() {

	var flagPort = flag.String("p", "8085", "port")
	var flagDataPath = flag.String("dp", "", "data path")
	var flagDBPath = flag.String("db", "", "database path")
	var flagReloadDB = flag.Bool("reload", false, "reload database")
	flag.Parse()

	a, err := deregserver.NewDeregService(*flagDataPath, *flagDBPath, *flagReloadDB)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := dereg.NewDeregServer(deregserver.NewServer(a))

	twirpserver.Start(*flagPort, a.Name, twirpHandler)
}
