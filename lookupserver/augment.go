package lookupserver

import (
	"context"

	"github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/rpc/dereg"
	"github.com/engelsjk/faadb/rpc/engine"
	"github.com/engelsjk/faadb/rpc/master"
	"github.com/engelsjk/faadb/rpc/reserved"
)

type Aircraft interface {
}

type AugmentedA struct {
	Aircraft     interface{}            `json:"Aircraft"`
	AircraftType *aircraft.AircraftType `json:"AircraftType"`
	EngineType   *engine.EngineType     `json:"EngineType"`
}

type AugmentedAircraft []*AugmentedA

type Augmenter struct {
	GetAircraftType func(context.Context, *aircraft.Query) (*aircraft.AircraftType, error)
	GetEngineType   func(context.Context, *engine.Query) (*engine.EngineType, error)
}

func (g Augmenter) AugmentA(ctx context.Context, a interface{}) *AugmentedA {

	var aircraftCode, engineCode string
	aa := &AugmentedA{Aircraft: a}
	switch v := a.(type) {
	case nil:
		return aa
	case *master.A:
		aircraftCode, engineCode = v.ManufacturerAircraftModelCode, v.ManufacturerEngineModelCode
	case *reserved.A:
		aircraftCode, engineCode = "", ""
	case *dereg.A:
		aircraftCode, engineCode = v.ManufacturerAircraftModelCode, v.ManufacturerEngineModelCode
	default:
		return aa
	}

	var err error
	if aa.AircraftType, err = g.GetAircraftType(ctx, &aircraft.Query{
		ManufacturerModelSeries: aircraftCode,
	}); err != nil {
		aa.AircraftType = nil
	}
	if aa.EngineType, err = g.GetEngineType(ctx, &engine.Query{
		ManufacturerModel: engineCode,
	}); err != nil {
		aa.EngineType = nil
	}
	return aa
}

func (g Augmenter) AugmentAircraft(ctx context.Context, a interface{}) AugmentedAircraft {
	switch v := a.(type) {
	case nil:
		return nil
	case *master.Aircraft:
		augmentedAircraft := make(AugmentedAircraft, len(v.A))
		for i, ac := range v.A {
			augmentedAircraft[i] = g.AugmentA(ctx, ac)
		}
		return augmentedAircraft
	case *reserved.Aircraft:
		augmentedAircraft := make(AugmentedAircraft, len(v.A))
		for i, ac := range v.A {
			augmentedAircraft[i] = g.AugmentA(ctx, ac)
		}
		return augmentedAircraft
	case *dereg.Aircraft:
		augmentedAircraft := make(AugmentedAircraft, len(v.A))
		for i, ac := range v.A {
			augmentedAircraft[i] = g.AugmentA(ctx, ac)
		}
		return augmentedAircraft
	default:
		return nil
	}
}
