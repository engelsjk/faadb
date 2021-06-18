package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/dereg"
	"github.com/engelsjk/faadb/servers/deregserver"
	server "github.com/engelsjk/faadb/twirp-web-server"
)

func main() {

	var flagPort = flag.String("p", "8085", "port")
	var flagDataPath = flag.String("dp", "DEREG.txt", "data path")
	var flagDBPath = flag.String("db", "dereg.db", "database path")
	flag.Parse()

	a, err := deregserver.NewDeregService(*flagDataPath, *flagDBPath)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := dereg.NewDeregServer(deregserver.NewServer(a))

	server.Start(*flagPort, a.Name, twirpHandler)
}
