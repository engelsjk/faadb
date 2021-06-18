package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/engelsjk/faadb/services/aircraft/rpc"
)

func main() {

	addr := "http://localhost:8082" // aircraft server

	client := rpc.NewAircraftProtobufClient(addr, &http.Client{})

	aircraftType, err := client.GetAircraftType(context.Background(), &rpc.Query{ManufacturerModelSeries: "2802630"})
	if err != nil {
		fmt.Printf("aircraft: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("aircraft: %+v\n", aircraftType)
}
