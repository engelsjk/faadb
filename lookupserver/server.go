package lookupserver

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	lookup *LookupService
}

func NewServer(lookup *LookupService) *Server {
	return &Server{lookup: lookup}
}
func (s Server) Start(port string) {

	e := echo.New()

	addRoutes(e, s.lookup)

	addr := net.JoinHostPort("", port)

	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:        addr,
		Handler:     e,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	log.Printf("%s : running server at %s\n", s.lookup.Name, addr)
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
	log.Printf("%s : shutting down...\n", s.lookup.Name)

	go func() {
		<-signalChan
		log.Fatalf("%s : terminating...\n", s.lookup.Name)
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("%s : shutdown error: %v\n", s.lookup.Name, err)
		cancel()
		defer os.Exit(1)
		return
	}

	log.Printf("%s : gracefully stopped\n", s.lookup.Name)

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}
