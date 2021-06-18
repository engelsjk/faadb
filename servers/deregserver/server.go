package deregserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/engelsjk/faadb/rpc/dereg"
)

type Server struct {
	dereg *DeregService
}

func NewServer(dereg *DeregService) *Server {
	return &Server{dereg: dereg}
}

func (s *Server) GetAircraft(ctx context.Context, query *dereg.Query) (*dereg.Aircraft, error) {
	var (
		bs  [][]byte
		err error
	)
	if query.NNumber != "" {
		nnumber := query.NNumber
		exact := true
		r := []rune(query.NNumber)
		if string(r[0]) == "*" {
			nnumber = string(r[1:])
			exact = false
		}
		bs, err = s.dereg.svc.List("nnumber", nnumber, "nnumber", exact)
	}
	if query.RegistrantName != "" {
		bs, err = s.dereg.svc.List("registrant_name", query.RegistrantName, "registrant.name", true)
	}
	if err != nil {
		return nil, err
	}
	return bytesToAircraft(bs)
}

func bytesToA(b []byte) (*dereg.A, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &dereg.A{
		NNumber:                                  record.NNumber,
		SerialNumber:                             record.SerialNumber,
		ManufacturerAircraftModelCode:            record.Manufacturer.AircraftModelCode,
		ManufacturerEngineModelCode:              record.Manufacturer.EngineModelCode,
		ManufacturerYear:                         record.Manufacturer.Year,
		Status:                                   record.StatusCode.Description,
		RegistrantType:                           record.Registrant.Type.Description,
		RegistrantName:                           record.Registrant.Name,
		RegistrantStreet1:                        record.Registrant.Street1,
		RegistrantStreet2:                        record.Registrant.Street2,
		RegistrantCity:                           record.Registrant.City,
		RegistrantState:                          record.Registrant.State,
		RegistrantZipCode:                        record.Registrant.ZipCode,
		RegistrantRegion:                         record.Registrant.Region.Description,
		RegistrantCounty:                         record.Registrant.County,
		RegistrantCountry:                        record.Registrant.Country,
		RegistrantPhysicalAddress:                record.Registrant.PhysicalAddress,
		RegistrantPhysicalAddress2:               record.Registrant.PhysicalAddress2,
		RegistrantPhysicalCity:                   record.Registrant.PhysicalCity,
		RegistrantPhysicalState:                  record.Registrant.PhysicalState,
		RegistrantPhysicalZipCode:                record.Registrant.PhysicalZipCode,
		RegistrantPhysicalCounty:                 record.Registrant.PhysicalCounty,
		RegistrantPhysicalCountry:                record.Registrant.PhysicalCountry,
		CertificationAirworthinessClassification: record.Certification.AirworthinessClassification.Description,
		CertificationApprovedOperations:          record.Certification.ApprovedOperation.Description,
		AirworthinessDate:                        record.AirworthinessDate,
		CancelDate:                               record.CancelDate,
		ExportCountry:                            record.ExportCountry,
		LastActivityDate:                         record.LastActivityDate,
		CertificationIssueDate:                   record.CertificateIssueDate,
		OwnershipOtherName1:                      record.Ownership.OtherName1,
		OwnershipOtherName2:                      record.Ownership.OtherName2,
		OwnershipOtherName3:                      record.Ownership.OtherName3,
		OwnershipOtherName4:                      record.Ownership.OtherName4,
		OwnershipOtherName5:                      record.Ownership.OtherName5,
		KitManufacturerName:                      record.Kit.ManufacturerName,
		KitModelName:                             record.Kit.ModelName,
		ModeSCode:                                record.ModeS.Code,
		ModeSCodeHex:                             record.ModeS.CodeHex,
	}, nil
}

func bytesToAircraft(bs [][]byte) (*dereg.Aircraft, error) {
	as := make([]*dereg.A, len(bs))
	for i, b := range bs {
		a, err := bytesToA(b)
		if err != nil {
			return nil, err
		}
		as[i] = a
	}
	return &dereg.Aircraft{A: as}, nil
}

func (s *Server) Start(port string) {

	twirpHandler := dereg.NewDeregServer(s)

	addr := net.JoinHostPort("", port)

	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:        addr,
		Handler:     twirpHandler,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	fmt.Printf("running %s server at %s\n", s.dereg.Name, addr)
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}
