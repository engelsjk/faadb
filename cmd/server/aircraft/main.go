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

	aircraftrpc "github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/service/aircraft"
)

func main() {

	var flagPort = flag.String("p", "8082", "port")
	var flagDataPath = flag.String("dp", "ACFTREF.txt", "data path")
	var flagDBPath = flag.String("db", "aircraft.db", "database path")
	flag.Parse()

	a, err := aircraft.NewAircraft(*flagDataPath, *flagDBPath)
	if err != nil {
		log.Fatal(err)
	}

	server := aircraft.NewServer(a)
	twirpHandler := aircraftrpc.NewAircraftServer(server)

	addr := net.JoinHostPort("", *flagPort)

	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:        addr,
		Handler:     twirpHandler,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	fmt.Printf("running %s server at %s\n", a.Name, addr)
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
