package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/engine"
	"github.com/engelsjk/faadb/servers/engineserver"
	"github.com/engelsjk/faadb/twirpserver"
)

func main() {

	var flagPort = flag.String("p", "8083", "port")
	var flagDataPath = flag.String("d", "", "data path")
	var flagDBPath = flag.String("b", "", "database path")
	var flagReloadDB = flag.Bool("r", false, "reload database")
	flag.Parse()

	e, err := engineserver.NewEngineService(*flagDataPath, *flagDBPath, *flagReloadDB)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := engine.NewEngineServer(engineserver.NewServer(e))

	twirpserver.Start(*flagPort, e.Name, twirpHandler)
}
