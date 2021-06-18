package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/engelsjk/faadb/rpc/dereg"
)

func main() {

	addr := "http://localhost:8085" // dereg server

	client := dereg.NewDeregProtobufClient(addr, &http.Client{})

	GetMultipleAircraft(client, "10021")
	// GetMultipleAircraftByRegistrant(client, "BLUE HEN HELO LEASING CO")
}

func GetMultipleAircraft(client dereg.Dereg, nnumber string) {
	aircraft, err := client.GetMultipleAircraft(context.Background(), &dereg.Query{NNumber: nnumber})
	if err != nil {
		fmt.Printf("master: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("aircraft: %+v\n", aircraft)
}

func GetMultipleAircraftByRegistrant(client dereg.Dereg, registrant string) {
	aircraft, err := client.GetMultipleAircraftByRegistrant(context.Background(), &dereg.Query{RegistrantName: registrant})
	if err != nil {
		fmt.Printf("master: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("aircraft: %+v\n", aircraft)
}
