package lookupserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

type QuerySet struct {
	master   *master.Query
	reserved *reserved.Query
	dereg    *dereg.Query
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

func (l LookupService) GetAircraft(query *Query, filter *Filter) (*AircraftResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp := &AircraftResponse{}

	// todo: add error handling

	querySet := queryToQuerySet(query, filter)

	m, err := l.master.GetAircraft(ctx, querySet.master)
	if err != nil {
		log.Println(err.Error())
	}
	resp.Registered = l.Augment(m)

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
	querySet.master.NNumber = query.NNumber

	a, err := l.master.GetAircraft(ctx, querySet.master)
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
		querySet.master.SerialNumber = serialNumber
		querySet.reserved.SerialNumber = serialNumber
		querySet.dereg.SerialNumber = serialNumber
	}

	if query.RegistrantName == "same" {
		registrantName := a.A[0].RegistrantName
		querySet.master.RegistrantName = registrantName
		querySet.reserved.RegistrantName = registrantName
		querySet.dereg.RegistrantName = registrantName
	}

	if query.RegistrantStreet1 == "same" {
		registrantStreet1 := a.A[0].RegistrantStreet1
		querySet.master.RegistrantStreet1 = registrantStreet1
		querySet.reserved.RegistrantStreet1 = registrantStreet1
		querySet.dereg.RegistrantStreet1 = registrantStreet1
	}

	resp := &AircraftResponse{}

	m, _ := l.master.GetAircraft(ctx, querySet.master)
	resp.Registered = l.Augment(m)

	r, _ := l.reserved.GetAircraft(ctx, querySet.reserved)
	resp.Reserved = l.Augment(r)

	d, _ := l.dereg.GetAircraft(ctx, querySet.dereg)
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
// Query/Filter

func queryToQuerySet(query *Query, filter *Filter) QuerySet {

	querySet := QuerySet{
		master:   &master.Query{},
		reserved: &reserved.Query{},
		dereg:    &dereg.Query{},
	}

	if filter != nil {
		querySet.master.RegistrantState = filter.RegistrantState
		querySet.master.AircraftModelCode = filter.AircraftModelCode
		querySet.master.AirworthinessClassificationCode = filter.AirworthinessClassificationCode
		querySet.master.ApprovedOperationCode = filter.ApprovedOperationCode

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

	querySet.master.NNumber = query.NNumber
	querySet.master.SerialNumber = query.SerialNumber
	querySet.master.ModeSCodeHex = query.ModeSCodeHex
	querySet.master.RegistrantName = query.RegistrantName
	querySet.master.RegistrantStreet1 = query.RegistrantStreet1

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
