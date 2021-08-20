package lookupserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/engelsjk/faadb/rpc/active"
	"github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/rpc/dereg"
	"github.com/engelsjk/faadb/rpc/engine"
	"github.com/engelsjk/faadb/rpc/reserved"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	timeout = 5 * time.Second
)

type LookupService struct {
	Name      string
	active    active.Active
	aircraft  aircraft.Aircraft
	engine    engine.Engine
	reserved  reserved.Reserved
	dereg     dereg.Dereg
	augmenter Augmenter
}

type Options struct {
	ActiveAddr   string
	AircraftAddr string
	EngineAddr   string
	ReservedAddr string
	DeregAddr    string
}

type QuerySet struct {
	active   *active.Query
	reserved *reserved.Query
	dereg    *dereg.Query
}

func NewLookupService(opts Options) *LookupService {

	client := &http.Client{} // one client (??)

	m := active.NewActiveProtobufClient(opts.ActiveAddr, client)
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
		active:    m,
		aircraft:  a,
		engine:    e,
		reserved:  r,
		dereg:     d,
		augmenter: augmenter,
	}
}

type AircraftResponse struct {
	Active       AugmentedAircraft `json:"active"`
	Reserved     AugmentedAircraft `json:"reserved"`
	Deregistered AugmentedAircraft `json:"deregistered"`
}

func (l LookupService) GetAircraft(query *Query, filter *Filter) (*AircraftResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp := &AircraftResponse{}

	// todo: add error handling

	querySet := queryToQuerySet(query, filter)

	m, err := l.active.GetAircraft(ctx, querySet.active)
	if err != nil {
		log.Println(err.Error())
	}
	resp.Active = l.Augment(m)

	r, err := l.reserved.GetAircraft(ctx, querySet.reserved)
	if err != nil {
		log.Println(err.Error())
	}
	resp.Reserved = l.Augment(r)

	d, err := l.dereg.GetAircraft(ctx, querySet.dereg)
	if err != nil {
		log.Println(err.Error())
	}
	resp.Deregistered = l.Augment(d)

	return resp, nil
}

func (l LookupService) GetOtherAircraft(query *Query, filter *Filter) (*AircraftResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	querySet := queryToQuerySet(nil, nil)
	querySet.active.NNumber = query.NNumber

	a, err := l.active.GetAircraft(ctx, querySet.active)
	if err != nil {
		return nil, err
	}

	if len(a.A) != 1 {
		return nil, fmt.Errorf("expected one registered aircraft")
	}

	// filters

	if filter == nil {
		filter = &Filter{}
	}

	if filter.RegistrantState == "same" {
		filter.RegistrantState = a.A[0].RegistrantState
	}

	if filter.AircraftModelCode == "same" {
		filter.AircraftModelCode = a.A[0].ManufacturerAircraftModelCode
	}

	if filter.AirworthinessClassificationCode == "same" {
		filter.AirworthinessClassificationCode = a.A[0].CertificationAirworthinessClassificationCode
	}

	if filter.ApprovedOperationCode == "same" {
		filter.ApprovedOperationCode = a.A[0].CertificationApprovedOperationCode
	}

	querySet = queryToQuerySet(nil, filter)

	if query.SerialNumber == "same" {
		serialNumber := a.A[0].SerialNumber
		querySet.active.SerialNumber = serialNumber
		querySet.reserved.SerialNumber = serialNumber
		querySet.dereg.SerialNumber = serialNumber
	}

	if query.RegistrantName == "same" {
		registrantName := a.A[0].RegistrantName
		querySet.active.RegistrantName = registrantName
		querySet.reserved.RegistrantName = registrantName
		querySet.dereg.RegistrantName = registrantName
	}

	if query.RegistrantStreet1 == "same" {
		registrantStreet1 := a.A[0].RegistrantStreet1
		querySet.active.RegistrantStreet1 = registrantStreet1
		querySet.reserved.RegistrantStreet1 = registrantStreet1
		querySet.dereg.RegistrantStreet1 = registrantStreet1
	}

	resp := &AircraftResponse{}

	m, _ := l.active.GetAircraft(ctx, querySet.active)
	resp.Active = l.Augment(m)

	r, _ := l.reserved.GetAircraft(ctx, querySet.reserved)
	resp.Reserved = l.Augment(r)

	d, _ := l.dereg.GetAircraft(ctx, querySet.dereg)
	resp.Deregistered = l.Augment(d)

	return resp, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Active/Reserved/Dereg

func (l LookupService) GetActiveByNNumber(nnumber string) (*active.Aircraft, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return l.active.GetAircraft(ctx, &active.Query{NNumber: nnumber})
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
	case *active.A, *reserved.A, *dereg.A:
		augmentedAircraft := make(AugmentedAircraft, 1)
		augmentedAircraft[0] = l.augmenter.AugmentA(ctx, v)
		return augmentedAircraft
	case *active.Aircraft, *reserved.Aircraft, *dereg.Aircraft:
		return l.augmenter.AugmentAircraft(ctx, v)
	default:
		return nil
	}
}

func (l LookupService) AugmentToBytes(a interface{}) []byte {
	return ToBytes(l.Augment(a))
}

/////////////////////////////////////////////////////////////////////////////////
// Query/Filter

func queryToQuerySet(query *Query, filter *Filter) QuerySet {

	querySet := QuerySet{
		active:   &active.Query{},
		reserved: &reserved.Query{},
		dereg:    &dereg.Query{},
	}

	if filter != nil {
		querySet.active.RegistrantState = filter.RegistrantState
		querySet.active.AircraftModelCode = filter.AircraftModelCode
		querySet.active.AirworthinessClassificationCode = filter.AirworthinessClassificationCode
		querySet.active.ApprovedOperationCode = filter.ApprovedOperationCode

		querySet.reserved.RegistrantState = filter.RegistrantState
		querySet.reserved.AircraftModelCode = filter.AircraftModelCode
		querySet.reserved.AirworthinessClassificationCode = filter.AirworthinessClassificationCode
		querySet.reserved.ApprovedOperationCode = filter.ApprovedOperationCode

		querySet.dereg.RegistrantState = filter.RegistrantState
		querySet.dereg.AircraftModelCode = filter.AircraftModelCode
		querySet.dereg.AirworthinessClassificationCode = filter.AirworthinessClassificationCode
		querySet.dereg.ApprovedOperationCode = filter.ApprovedOperationCode
	}

	if query == nil {
		return querySet
	}

	querySet.active.NNumber = query.NNumber
	querySet.active.SerialNumber = query.SerialNumber
	querySet.active.ModeSCodeHex = query.ModeSCodeHex
	querySet.active.RegistrantName = query.RegistrantName
	querySet.active.RegistrantStreet1 = query.RegistrantStreet1

	querySet.reserved.NNumber = query.NNumber
	querySet.reserved.SerialNumber = query.SerialNumber
	querySet.reserved.ModeSCodeHex = query.ModeSCodeHex
	querySet.reserved.RegistrantName = query.RegistrantName
	querySet.reserved.RegistrantStreet1 = query.RegistrantStreet1

	querySet.dereg.NNumber = query.NNumber
	querySet.dereg.SerialNumber = query.SerialNumber
	querySet.dereg.ModeSCodeHex = query.ModeSCodeHex
	querySet.dereg.RegistrantName = query.RegistrantName
	querySet.dereg.RegistrantStreet1 = query.RegistrantStreet1

	return querySet
}

/////////////////////////////////////////////////////////////////////////////////
// Marshal

func ToBytes(a interface{}) []byte {
	switch v := a.(type) {
	case *active.A:
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
	case *active.Aircraft, *reserved.Aircraft, *dereg.Aircraft, *AugmentedAircraft, AugmentedAircraft:
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
