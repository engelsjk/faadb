package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/engelsjk/faadb/rpc/active"
)

func main() {

	addr := "http://localhost:8081" // active server

	client := active.NewActiveProtobufClient(addr, &http.Client{})

	// GetAircraft(client, "614ar", "")
	GetAircraft(client, "*0dz", "")
	// GetAircraft(client, "", "BLUE HEN HELO LEASING CO")
}

func GetAircraft(client active.Active, nnumber, registrant string) {
	query := &active.Query{NNumber: nnumber, RegistrantName: registrant}
	aircraft, err := client.GetAircraft(context.Background(), query)
	if err != nil {
		fmt.Printf("active: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("active: %+v\n", aircraft)
}
