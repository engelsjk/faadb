package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/engelsjk/faadb/rpc/master"
)

func main() {

	addr := "http://localhost:8081" // master server

	client := master.NewMasterProtobufClient(addr, &http.Client{})

	// GetAircraft(client, "614ar")
	GetPossibleAircraft(client, "0dz")
	// GetMultipleAircraftByRegistrant(client, "BLUE HEN HELO LEASING CO")
}

func GetAircraft(client master.Master, nnumber string) {
	aircraft, err := client.GetAircraft(context.Background(), &master.Query{NNumber: nnumber})
	if err != nil {
		fmt.Printf("master: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("aircraft: %+v\n", aircraft)
}

func GetPossibleAircraft(client master.Master, nnumber string) {
	aircraft, err := client.GetPossibleAircraft(context.Background(), &master.Query{NNumber: nnumber})
	if err != nil {
		fmt.Printf("master: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("aircraft: %+v\n", aircraft)
}

func GetMultipleAircraftByRegistrant(client master.Master, registrant string) {
	aircraft, err := client.GetMultipleAircraftByRegistrantName(context.Background(), &master.Query{RegistrantName: registrant})
	if err != nil {
		fmt.Printf("master: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("aircraft: %+v\n", aircraft)
}
