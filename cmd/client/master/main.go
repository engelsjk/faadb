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

	// GetAircraft(client, "614ar", "")
	GetAircraft(client, "*0dz", "")
	// GetAircraft(client, "", "BLUE HEN HELO LEASING CO")
}

func GetAircraft(client master.Master, nnumber, registrant string) {
	query := &master.Query{NNumber: nnumber, RegistrantName: registrant}
	aircraft, err := client.GetAircraft(context.Background(), query)
	if err != nil {
		fmt.Printf("master: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("mastermaster: %+v\n", aircraft)
}
