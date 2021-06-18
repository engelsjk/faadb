package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/engelsjk/faadb/rpc/engine"
)

func main() {

	addr := "http://localhost:8083" // engine server

	client := engine.NewEngineProtobufClient(addr, &http.Client{})

	engineType, err := client.GetEngineType(context.Background(), &engine.Query{ManufacturerModel: "52150"})
	if err != nil {
		fmt.Printf("engine: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("engine: %+v\n", engineType)
}
