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
	var flagDataPath = flag.String("dp", "", "data path")
	var flagDBPath = flag.String("db", "", "database path")
	var flagReloadDB = flag.Bool("reload", false, "reload database")
	flag.Parse()

	e, err := engineserver.NewEngineService(*flagDataPath, *flagDBPath, *flagReloadDB)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := engine.NewEngineServer(engineserver.NewServer(e))

	twirpserver.Start(*flagPort, e.Name, twirpHandler)
}
