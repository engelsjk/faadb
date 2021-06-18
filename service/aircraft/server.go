package aircraft

import (
	"context"

	pb "github.com/engelsjk/faadb/rpc/aircraft"
)

type Server struct {
	aircraft *Aircraft
}

func NewServer(aircraft *Aircraft) *Server {
	return &Server{aircraft: aircraft}
}

func (s *Server) GetAircraftType(ctx context.Context, query *pb.Query) (aircraftType *pb.AircraftType, err error) {
	bs, err := s.aircraft.svc.Get(query.ManufacturerModelSeries)
	if err != nil {
		return nil, err
	}
	return bytesToAircraftType(bs[0])
}

func bytesToAircraftType(b []byte) (*pb.AircraftType, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &pb.AircraftType{
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
