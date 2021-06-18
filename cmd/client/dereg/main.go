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

	GetAircraft(client, "10021", "")
	fmt.Println("***")
	GetAircraft(client, "*0021", "")
	fmt.Println("***")
	GetAircraft(client, "", "BLUE HEN HELO LEASING CO")
}

func GetAircraft(client dereg.Dereg, nnumber, registrantName string) {
	query := &dereg.Query{NNumber: nnumber, RegistrantName: registrantName}
	aircraft, err := client.GetAircraft(context.Background(), query)
	if err != nil {
		fmt.Printf("dereg: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("dereg: %+v\n", aircraft)
}
