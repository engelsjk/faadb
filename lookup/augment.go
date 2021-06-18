package lookup

import (
	"context"

	"github.com/engelsjk/faadb/rpc/aircraft"
	"github.com/engelsjk/faadb/rpc/engine"
	"github.com/engelsjk/faadb/rpc/master"
)

type AugmentedAircraft struct {
	*master.Aircraft
	AircraftType *aircraft.AircraftType `json:"AircraftType"`
	EngineType   *engine.EngineType     `json:"EngineType"`
}

type MultipleAugmentedAircraft []*AugmentedAircraft

type Augmenter struct {
	GetAircraftType func(context.Context, *aircraft.Query) (*aircraft.AircraftType, error)
	GetEngineType   func(context.Context, *engine.Query) (*engine.EngineType, error)
}

func (a Augmenter) AugmentAircraft(ctx context.Context, m *master.Aircraft) *AugmentedAircraft {
	aa := &AugmentedAircraft{m, nil, nil}
	if m == nil {
		return aa
	}
	var err error
	if aa.AircraftType, err = a.GetAircraftType(ctx, &aircraft.Query{
		ManufacturerModelSeries: m.ManufacturerAircraftModelCode,
	}); err != nil {
		aa.AircraftType = nil
	}
	if aa.EngineType, err = a.GetEngineType(ctx, &engine.Query{
		ManufacturerModel: m.ManufacturerEngineModelCode,
	}); err != nil {
		aa.EngineType = nil
	}
	return aa
}

func (a Augmenter) AugmentMultipleAircraft(ctx context.Context, m *master.MultipleAircraft) MultipleAugmentedAircraft {
	multipleAugmentedAircraft := make(MultipleAugmentedAircraft, len(m.Aircraft))
	for i, ac := range m.Aircraft {
		multipleAugmentedAircraft[i] = a.AugmentAircraft(ctx, ac)
	}
	return multipleAugmentedAircraft
}
