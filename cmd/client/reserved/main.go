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

	GetAircraft(client, "138SS", "")
	// GetAircraft(client, "", "NATIONAL AERONAUTICS AND SPACE ADMINISTRATION")
}

func GetAircraft(client reserved.Reserved, nnumber, registrant string) {
	query := &reserved.Query{NNumber: nnumber, RegistrantName: registrant}
	aircraft, err := client.GetAircraft(context.Background(), query)
	if err != nil {
		fmt.Printf("reserved: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("reserved: %+v\n", aircraft)
}
