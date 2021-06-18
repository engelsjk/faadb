package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/engelsjk/faadb/services/reserved/rpc"
)

func main() {

	addr := "http://localhost:8084" // reserved server

	client := rpc.NewReservedProtobufClient(addr, &http.Client{})

	GetAircraft(client, "138SS", "")
	// GetAircraft(client, "", "NATIONAL AERONAUTICS AND SPACE ADMINISTRATION")
}

func GetAircraft(client rpc.Reserved, nnumber, registrant string) {
	query := &rpc.Query{NNumber: nnumber, RegistrantName: registrant}
	aircraft, err := client.GetAircraft(context.Background(), query)
	if err != nil {
		fmt.Printf("reserved: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("reserved: %+v\n", aircraft)
}
