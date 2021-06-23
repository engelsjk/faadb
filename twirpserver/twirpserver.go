package twirpserver

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(port, name string, handler http.Handler) {

	addr := net.JoinHostPort("", port)

	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:        addr,
		Handler:     handler,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	log.Printf("%s : running server at %s\n", name, addr)
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("%s : listen error: %v", name, err)
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
	log.Printf("%s : shutting down...\n", name)

	go func() {
		<-signalChan
		log.Fatalf("%s : terminating...\n", name)
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("%s : shutdown error: %v\n", name, err)
		cancel()
		defer os.Exit(1)
		return
	}

	log.Printf("%s : gracefully stopped\n", name)

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}
