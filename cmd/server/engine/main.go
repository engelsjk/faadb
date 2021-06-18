package main

import (
	"flag"
	"log"

	"github.com/engelsjk/faadb/rpc/engine"
	"github.com/engelsjk/faadb/servers/engineserver"
	server "github.com/engelsjk/faadb/twirp-web-server"
)

func main() {

	var flagPort = flag.String("p", "8083", "port")
	var flagDataPath = flag.String("dp", "ENGINE.txt", "data path")
	var flagDBPath = flag.String("db", "engine.db", "database path")
	flag.Parse()

	e, err := engineserver.NewEngineService(*flagDataPath, *flagDBPath)
	if err != nil {
		log.Fatal(err)
	}

	twirpHandler := engine.NewEngineServer(engineserver.NewServer(e))

	server.Start(*flagPort, e.Name, twirpHandler)
}
