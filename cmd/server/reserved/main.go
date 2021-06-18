package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/services/reserved"
	"github.com/engelsjk/faadb/services/reserved/rpc"
	server "github.com/engelsjk/faadb/twirp-web-server"
)

func main() {

	var flagPort = flag.String("p", "8084", "port")
	var flagDataPath = flag.String("dp", "RESERVED.txt", "data path")
	var flagDBPath = flag.String("db", "reserved.db", "database path")
	flag.Parse()

	r, err := reserved.NewReserved(*flagDataPath, *flagDBPath)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := rpc.NewReservedServer(reserved.NewServer(r))

	server.Start(*flagPort, r.Name, twirpHandler)
}
