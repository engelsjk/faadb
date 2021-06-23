package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/master"
	"github.com/engelsjk/faadb/servers/masterserver"
	"github.com/engelsjk/faadb/twirpserver"
)

func main() {

	var flagPort = flag.String("p", "8081", "port")
	var flagDataPath = flag.String("dp", "MASTER.txt", "data path")
	var flagDBPath = flag.String("db", "master.db", "database path")
	var flagReloadDB = flag.Bool("reload", false, "reload database")
	flag.Parse()

	m, err := masterserver.NewMasterService(*flagDataPath, *flagDBPath, *flagReloadDB)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := master.NewMasterServer(masterserver.NewServer(m))

	twirpserver.Start(*flagPort, m.Name, twirpHandler)
}
