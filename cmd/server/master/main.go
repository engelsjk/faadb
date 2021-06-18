package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/services/master"
	"github.com/engelsjk/faadb/services/master/rpc"
	server "github.com/engelsjk/faadb/twirp-web-server"
)

func main() {

	var flagPort = flag.String("p", "8081", "port")
	var flagDataPath = flag.String("dp", "MASTER.txt", "data path")
	var flagDBPath = flag.String("db", "master.db", "database path")
	flag.Parse()

	m, err := master.NewMasterService(*flagDataPath, *flagDBPath)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := rpc.NewMasterServer(master.NewServer(m))

	server.Start(*flagPort, m.Name, twirpHandler)
}
