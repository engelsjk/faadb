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

var (
	timeout = 5 * time.Second
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

type AircraftResponse struct {
	Registered   AugmentedAircraft `json:"registered"`
	Reserved     AugmentedAircraft `json:"reserved"`
	Deregistered AugmentedAircraft `json:"deregistered"`
}

func (l LookupService) GetAircraftByNNumber(nNumber string) (*AircraftResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp := &AircraftResponse{}

	// todo: add error handling

	m, _ := l.master.GetAircraft(ctx, &master.Query{NNumber: nNumber})
	resp.Registered = l.Augment(m)

	r, _ := l.reserved.GetAircraft(ctx, &reserved.Query{NNumber: nNumber})
	resp.Reserved = l.Augment(r)

	d, _ := l.dereg.GetAircraft(ctx, &dereg.Query{NNumber: nNumber})
	resp.Deregistered = l.Augment(d)

	return resp, nil
}

func (l LookupService) GetAircraftBySerialNumber(serialNumber string) (*AircraftResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp := &AircraftResponse{}

	// todo: add error handling

	m, _ := l.master.GetAircraft(ctx, &master.Query{SerialNumber: serialNumber})
	resp.Registered = l.Augment(m)

	r, _ := l.reserved.GetAircraft(ctx, &reserved.Query{SerialNumber: serialNumber})
	resp.Reserved = l.Augment(r)

	d, _ := l.dereg.GetAircraft(ctx, &dereg.Query{SerialNumber: serialNumber})
	resp.Deregistered = l.Augment(d)

	return resp, nil
}

func (l LookupService) GetOtherAircraftWithSameSerialNumber(nnumber string) (*AircraftResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	a, err := l.master.GetAircraft(ctx, &master.Query{NNumber: nnumber})
	if err != nil {
		return nil, err
	}

	if len(a.A) != 1 {
		return nil, fmt.Errorf("expected one registered aircraft")
	}

	serialNumber := a.A[0].SerialNumber

	resp := &AircraftResponse{}

	m, _ := l.master.GetAircraft(ctx, &master.Query{SerialNumber: serialNumber})
	resp.Registered = l.Augment(m)

	r, _ := l.reserved.GetAircraft(ctx, &reserved.Query{SerialNumber: serialNumber})
	resp.Reserved = l.Augment(r)

	d, _ := l.dereg.GetAircraft(ctx, &dereg.Query{SerialNumber: serialNumber})
	resp.Deregistered = l.Augment(d)

	return resp, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Master/Reserved/Dereg

func (l LookupService) GetMasterByNNumber(nnumber string) (*master.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return l.master.GetAircraft(ctx, &master.Query{NNumber: nnumber})
}

func (l LookupService) GetDeregByNNumber(nnumber string) (*dereg.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return l.dereg.GetAircraft(ctx, &dereg.Query{NNumber: nnumber})
}

func (l LookupService) GetReservedByNNumber(nnumber string) (*reserved.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return l.reserved.GetAircraft(ctx, &reserved.Query{NNumber: nnumber})
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Aircraft & Engine Type

func (l LookupService) GetAircraftType(code string) (*aircraft.AircraftType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return l.aircraft.GetAircraftType(ctx, &aircraft.Query{ManufacturerModelSeries: code})
}

func (l LookupService) GetEngineType(code string) (*engine.EngineType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return l.engine.GetEngineType(ctx, &engine.Query{ManufacturerModel: code})
}

/////////////////////////////////////////////////////////////////////////////////
// Augment

func (l LookupService) Augment(a interface{}) AugmentedAircraft {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	switch v := a.(type) {
	case *master.A, *reserved.A, *dereg.A:
		augmentedAircraft := make(AugmentedAircraft, 1)
		augmentedAircraft[0] = l.augmenter.AugmentA(ctx, v)
		return augmentedAircraft
	case *master.Aircraft, *reserved.Aircraft, *dereg.Aircraft:
		return l.augmenter.AugmentAircraft(ctx, v)
	default:
		return nil
	}
}

func (l LookupService) AugmentToBytes(a interface{}) []byte {
	return ToBytes(l.Augment(a))
}

/////////////////////////////////////////////////////////////////////////////////
// Marshal

func ToBytes(a interface{}) []byte {
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
	case *AircraftResponse:
		b, err := json.Marshal(v)
		if err != nil {
			return nil
		}
		return b
	default:
		return nil
	}
}
