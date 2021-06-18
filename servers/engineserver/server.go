package engineserver

import (
	"context"

	"github.com/engelsjk/faadb/rpc/engine"
)

type Server struct {
	engine *EngineService
}

func NewServer(engine *EngineService) *Server {
	return &Server{engine: engine}
}

func (s *Server) GetEngineType(ctx context.Context, query *engine.Query) (engineType *engine.EngineType, err error) {
	bs, err := s.engine.svc.Get(query.ManufacturerModel)
	if err != nil {
		return nil, err
	}
	return bytesToEngineType(bs[0])
}

func bytesToEngineType(b []byte) (*engine.EngineType, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &engine.EngineType{
		ManufacturerModelCode: record.ManufacturerModelCode,
		ManufacturerName:      record.ManufacturerName,
		ModelName:             record.ModelName,
		EngineType:            record.EngineType.Description,
		Horsepower:            record.Horsepower,
		PoundsOfThrust:        record.PoundsOfThrust,
	}, nil
}
