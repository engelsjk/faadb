package main

import (
	"flag"

	lookup "github.com/engelsjk/faadb/lookupserver"
	"github.com/engelsjk/faadb/rpc/active"
	"github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/rpc/engine"
	"github.com/engelsjk/faadb/rpc/reserved"
)

type Lookup struct {
	Name     string
	active   active.Active
	aircraft aircraft.Aircraft
	engine   engine.Engine
	reserved reserved.Reserved
}

func main() {

	var flagPort = flag.String("p", "8080", "port")
	var flagActiveAddr = flag.String("active", "http://localhost:8081", "active service addr")
	var flagAircraftAddr = flag.String("aircraft", "http://localhost:8082", "aircraft service addr")
	var flagEngineAddr = flag.String("engine", "http://localhost:8083", "engine service addr")
	var flagReservedAddr = flag.String("reserved", "http://localhost:8084", "reserved service addr")
	var flagDeregAddr = flag.String("dereg", "http://localhost:8085", "dereg service addr")
	flag.Parse()

	svc := lookup.NewLookupService(lookup.Options{
		ActiveAddr:   *flagActiveAddr,
		AircraftAddr: *flagAircraftAddr,
		EngineAddr:   *flagEngineAddr,
		ReservedAddr: *flagReservedAddr,
		DeregAddr:    *flagDeregAddr,
	})

	server := lookup.NewServer(svc)
	server.Start(*flagPort)

	///////////////////////////////////////////////////////////////////////

}
