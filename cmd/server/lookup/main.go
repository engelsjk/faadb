package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/engelsjk/faadb/lookup"
	"github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/rpc/dereg"
	"github.com/engelsjk/faadb/rpc/engine"
	"github.com/engelsjk/faadb/rpc/master"
	"github.com/engelsjk/faadb/rpc/reserved"
	"github.com/labstack/echo/v4"
)

type Lookup struct {
	Name     string
	master   master.Master
	aircraft aircraft.Aircraft
	engine   engine.Engine
	reserved reserved.Reserved
	dereg    dereg.Dereg
}

func main() {

	var flagPort = flag.String("p", "8080", "port")
	var flagMasterAddr = flag.String("master", "http://localhost:8081", "master service addr")
	var flagAircraftAddr = flag.String("aircraft", "http://localhost:8082", "aircraft service addr")
	var flagEngineAddr = flag.String("engine", "http://localhost:8083", "engine service addr")
	var flagReservedAddr = flag.String("reserved", "http://localhost:8084", "reserved service addr")
	var flagDeregAddr = flag.String("dereg", "http://localhost:8085", "dereg service addr")
	flag.Parse()

	lu := lookup.NewLookup(lookup.Options{
		MasterAddr:   *flagMasterAddr,
		AircraftAddr: *flagAircraftAddr,
		EngineAddr:   *flagEngineAddr,
		ReservedAddr: *flagReservedAddr,
		DeregAddr:    *flagDeregAddr,
	})

	///////////////////////////////////////////////////////////////////////
	// todo: define lookup server routes someplace else?
	// also clean up this code

	e := echo.New()
	e.GET("/aircraft", func(c echo.Context) error {
		nnumber := c.QueryParam("n")
		registrant := c.QueryParam("registrant")
		sameRegistrantName := c.QueryParam("sameRegistrantName")
		sameRegistrantStreet1 := c.QueryParam("sameRegistrantStreet1")
		partial := c.QueryParam("partial")

		// todo: clean up query param logic!
		if nnumber != "" {
			if sameRegistrantName == "true" {
				return c.JSONBlob(http.StatusOK, lu.AugmentToBytes(lu.GetOtherAircraftByRegistrant(nnumber, "name")))
			}
			if sameRegistrantStreet1 == "true" {
				return c.JSONBlob(http.StatusOK, lu.AugmentToBytes(lu.GetOtherAircraftByRegistrant(nnumber, "street1")))
			}
			if partial == "true" {
				// todo: maybe only allow a partial nnumber of > 3 chars to limit response?
				return c.JSONBlob(http.StatusOK, lookup.ToBytes(lu.GetMultipleAircraftByPartialNNumber(nnumber)))
			}
			return c.JSONBlob(http.StatusOK, lu.AugmentToBytes(lu.GetAircraftByNNumber(nnumber)))
		}
		if registrant != "" {
			return c.JSONBlob(http.StatusOK, lu.AugmentToBytes(lu.GetMultipleAircraftByRegistrantName(registrant)))
		}
		return c.JSON(http.StatusBadRequest, "request not valid")
	})

	///////////////////////////////////////////////////////////////////////

	addr := net.JoinHostPort("", *flagPort)

	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:        addr,
		Handler:     e,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	fmt.Printf("running %s server at %s\n", lu.Name, addr)
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}
