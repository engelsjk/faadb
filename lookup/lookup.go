package lookup

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

type Lookup struct {
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
	ReservedAddr string
	DeregAddr    string
	AircraftAddr string
	EngineAddr   string
}

func NewLookup(opts Options) *Lookup {

	client := &http.Client{}

	m := master.NewMasterProtobufClient(opts.MasterAddr, client)
	r := reserved.NewReservedProtobufClient(opts.ReservedAddr, client)
	d := dereg.NewDeregProtobufClient(opts.ReservedAddr, client)
	a := aircraft.NewAircraftProtobufClient(opts.AircraftAddr, client)
	e := engine.NewEngineProtobufClient(opts.EngineAddr, client)

	augmenter := Augmenter{
		GetAircraftType: a.GetAircraftType,
		GetEngineType:   e.GetEngineType,
	}

	return &Lookup{
		Name:      "lookup",
		master:    m,
		aircraft:  a,
		engine:    e,
		reserved:  r,
		dereg:     d,
		augmenter: augmenter,
	}
}

func (l Lookup) GetAircraftByNNumber(nnumber string) (*master.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.master.GetAircraft(ctx, &master.Query{NNumber: nnumber})
}

func (l Lookup) GetMultipleAircraftByPartialNNumber(nnumber string) (*master.MultipleAircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.master.GetPossibleAircraft(ctx, &master.Query{NNumber: nnumber})
}

func (l Lookup) GetMultipleAircraftByRegistrantName(registrantName string) (*master.MultipleAircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.master.GetMultipleAircraftByRegistrantName(ctx, &master.Query{RegistrantName: registrantName})
}

func (l Lookup) GetMultipleAircraftByRegistrantStreet1(registrantStreet1 string) (*master.MultipleAircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.master.GetMultipleAircraftByRegistrantStreet1(ctx, &master.Query{RegistrantStreet1: registrantStreet1})
}

func (l Lookup) GetOtherAircraftByRegistrant(nnumber, what string) (*master.MultipleAircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a, err := l.master.GetAircraft(ctx, &master.Query{NNumber: nnumber})
	if err != nil {
		return nil, err
	}
	switch what {
	case "name":
		return l.GetMultipleAircraftByRegistrantName(a.RegistrantName)
	case "street1":
		return l.GetMultipleAircraftByRegistrantStreet1(a.RegistrantStreet1)
	default:
		return nil, fmt.Errorf("unknown registrant query '%s'", what)
	}
}

func (l Lookup) GetAircraftType(code string) (*aircraft.AircraftType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.aircraft.GetAircraftType(ctx, &aircraft.Query{ManufacturerModelSeries: code})
}

func (l Lookup) GetEngineType(code string) (*engine.EngineType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return l.engine.GetEngineType(ctx, &engine.Query{ManufacturerModel: code})
}

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

func (l Lookup) Augment(a interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch v := a.(type) {
	case *master.Aircraft:
		return l.augmenter.AugmentAircraft(ctx, v), nil
	case *master.MultipleAircraft:
		return l.augmenter.AugmentMultipleAircraft(ctx, v), nil
	default:
		return nil, fmt.Errorf("unknown type '%s'", v)
	}
}

func (l Lookup) AugmentToBytes(a interface{}, err error) []byte {
	return ToBytes(l.Augment(a, err))
}

func ToBytes(a interface{}, err error) []byte {
	if err != nil {
		return nil
	}
	switch v := a.(type) {
	case *master.Aircraft:
		b, err := protojson.Marshal(v)
		if err != nil {
			return nil
		}
		return b
	case *master.MultipleAircraft, *AugmentedAircraft, MultipleAugmentedAircraft:
		b, err := json.Marshal(v)
		if err != nil {
			return nil
		}
		return b
	default:
		return nil
	}
}
