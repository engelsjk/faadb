package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/engelsjk/faadb/rpc/aircraft"
)

func main() {

	addr := "http://localhost:8082" // aircraft server

	client := aircraft.NewAircraftProtobufClient(addr, &http.Client{})

	aircraftType, err := client.GetAircraftType(context.Background(), &aircraft.Query{ManufacturerModelSeries: "2802630"})
	if err != nil {
		fmt.Printf("aircraft: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("aircraft_type: %+v\n", aircraftType)
}
