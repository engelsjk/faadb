package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/engelsjk/faadb/rpc/reserved"
)

func main() {

	addr := "http://localhost:8084" // reserved server

	client := reserved.NewReservedProtobufClient(addr, &http.Client{})

	// GetReservation(client, "138SS")
	GetMultipleReservationsByRegistrant(client, "NATIONAL AERONAUTICS AND SPACE ADMINISTRATION")
}

func GetReservation(client reserved.Reserved, nnumber string) {
	reservation, err := client.GetAircraft(context.Background(), &reserved.Query{NNumber: nnumber})
	if err != nil {
		fmt.Printf("reserved: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("reservation: %+v\n", reservation)
}

func GetMultipleReservationsByRegistrant(client reserved.Reserved, registrant string) {
	reservations, err := client.GetMultipleAircraftByRegistrant(context.Background(), &reserved.Query{RegistrantName: registrant})
	if err != nil {
		fmt.Printf("reserved: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("reservations: %+v\n", reservations)
}
