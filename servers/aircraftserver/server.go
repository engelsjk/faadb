package aircraftserver

import (
	"context"

	"github.com/engelsjk/faadb/rpc/aircraft"
)

type Server struct {
	aircraft *AircraftService
}

func NewServer(aircraft *AircraftService) *Server {
	return &Server{aircraft: aircraft}
}

func (s *Server) GetAircraftType(ctx context.Context, query *aircraft.Query) (aircraftType *aircraft.AircraftType, err error) {
	bs, err := s.aircraft.svc.Get(query.ManufacturerModelSeries)
	if err != nil {
		return nil, err
	}
	return bytesToAircraftType(bs[0])
}

func bytesToAircraftType(b []byte) (*aircraft.AircraftType, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &aircraft.AircraftType{
		ManufacturerModelSeriesCode: record.ManufacturerModelSeriesCode,
		ManufacturerName:            record.ManufacturerName,
		ModelName:                   record.ModelName,
		AircraftType:                record.AircraftType.Description,
		EngineType:                  record.EngineType.Description,
		AircraftCategoryCode:        record.AircraftCategoryCode.Description,
		BuilderCertificationCode:    record.BuilderCertificationCode.Description,
		NumberOfEngines:             record.NumberOfEngines,
		NumberOfSeats:               record.NumberOfSeats,
		AircraftWeight:              record.AircraftWeight.Description,
		AircraftCruisingSpeed:       record.AircraftCruisingSpeed,
		TCDataSheet:                 record.TCDataSheet,
		TCDataHolder:                record.TCDataHolder,
	}, nil
}
