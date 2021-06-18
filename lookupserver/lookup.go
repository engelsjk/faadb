package lookupserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/rpc/dereg"
	"github.com/engelsjk/faadb/rpc/engine"
	"github.com/engelsjk/faadb/rpc/master"
	"github.com/engelsjk/faadb/rpc/reserved"
	"google.golang.org/protobuf/encoding/protojson"
)

type LookupService struct {
	Name      string
	master    master.Master
	aircraft  aircraft.Aircraft
	engine    engine.Engine
	reserved  reserved.Reserved
	dereg     dereg.Dereg
	augmenter Augmenter
}

type Options struct {
	MasterAddr   string
	AircraftAddr string
	EngineAddr   string
	ReservedAddr string
	DeregAddr    string
}

func NewLookupService(opts Options) *LookupService {

	client := &http.Client{} // one client (??)

	m := master.NewMasterProtobufClient(opts.MasterAddr, client)
	a := aircraft.NewAircraftProtobufClient(opts.AircraftAddr, client)
	e := engine.NewEngineProtobufClient(opts.EngineAddr, client)
	r := reserved.NewReservedProtobufClient(opts.ReservedAddr, client)
	d := dereg.NewDeregProtobufClient(opts.DeregAddr, client)

	augmenter := Augmenter{
		GetAircraftType: a.GetAircraftType,
		GetEngineType:   e.GetEngineType,
	}

	return &LookupService{
		Name:      "lookup",
		master:    m,
		aircraft:  a,
		engine:    e,
		reserved:  r,
		dereg:     d,
		augmenter: augmenter,
	}
}

// func (l LookupService) GetAircraftByNNumber2(nnumber string) (*master.Aircraft, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	a, err := l.master.GetAircraft(ctx, &master.Query{NNumber: nnumber})
// 	r, err := l.reserved.GetAircraft(ctx, &reserved.Query{NNumber: nnumber})
// 	d, err := l.dereg.GetAircraft(ctx, &dereg.Query{NNumber: nnumber})

// 	return nil, nil
// }

// Master

func (l LookupService) GetAircraftByNNumber(nnumber string) (*master.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.master.GetAircraft(ctx, &master.Query{NNumber: nnumber})
}

func (l LookupService) GetAircraftByRegistrantName(registrantName string) (*master.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.master.GetAircraft(ctx, &master.Query{RegistrantName: registrantName})
}

func (l LookupService) GetAircraftByRegistrantStreet1(registrantStreet1 string) (*master.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.master.GetAircraft(ctx, &master.Query{RegistrantStreet1: registrantStreet1})
}

func (l LookupService) GetOtherAircraftByRegistrantName(nnumber string) (*master.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a, err := l.master.GetAircraft(ctx, &master.Query{NNumber: nnumber})
	if err != nil {
		return nil, err
	}
	return l.GetAircraftByRegistrantName(a.A[0].RegistrantName)
}

// Dereg

func (l LookupService) GetDeregByNNumber(nnumber string) (*dereg.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.dereg.GetAircraft(ctx, &dereg.Query{NNumber: nnumber})
}

func (l LookupService) GetDeregBySerialNumber(serialNumber string) (*dereg.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.dereg.GetAircraft(ctx, &dereg.Query{SerialNumber: serialNumber})
}

func (l LookupService) GetOtherDeregBySerialNumber(nnumber string) (*dereg.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a, err := l.dereg.GetAircraft(ctx, &dereg.Query{NNumber: nnumber})
	if err != nil {
		return nil, err
	}
	return l.GetDeregBySerialNumber(a.A[0].SerialNumber)
}

// Reserved

func (l LookupService) GetReservedByNNumber(nnumber string) (*dereg.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.dereg.GetAircraft(ctx, &dereg.Query{NNumber: nnumber})
}

/////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

// Aircraft Type

func (l LookupService) GetAircraftType(code string) (*aircraft.AircraftType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.aircraft.GetAircraftType(ctx, &aircraft.Query{ManufacturerModelSeries: code})
}

// Engine Type

func (l LookupService) GetEngineType(code string) (*engine.EngineType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.engine.GetEngineType(ctx, &engine.Query{ManufacturerModel: code})
}

/////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

func (l LookupService) Augment(a interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch v := a.(type) {
	case *master.A, *reserved.A, *dereg.A:
		return l.augmenter.AugmentA(ctx, v), nil
	case *master.Aircraft, *reserved.Aircraft, *dereg.Aircraft:
		return l.augmenter.AugmentAircraft(ctx, v), nil
	default:
		return nil, fmt.Errorf("unknown type '%s'", v)
	}
}

func (l LookupService) AugmentToBytes(a interface{}, err error) []byte {
	return ToBytes(l.Augment(a, err))
}

/////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

func ToBytes(a interface{}, err error) []byte {
	if err != nil {
		return nil
	}
	switch v := a.(type) {
	case *master.A:
		b, err := protojson.Marshal(v)
		if err != nil {
			return nil
		}
		return b
	case *reserved.A:
		b, err := protojson.Marshal(v)
		if err != nil {
			return nil
		}
		return b
	case *dereg.A:
		b, err := protojson.Marshal(v)
		if err != nil {
			return nil
		}
		return b
	case *master.Aircraft, *reserved.Aircraft, *dereg.Aircraft, *AugmentedAircraft, AugmentedAircraft:
		b, err := json.Marshal(v)
		if err != nil {
			return nil
		}
		return b
	default:
		return nil
	}
}
