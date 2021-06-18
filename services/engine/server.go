package engine

import (
	"context"

	"github.com/engelsjk/faadb/services/engine/rpc"
)

type Server struct {
	engine *EngineService
}

func NewServer(engine *EngineService) *Server {
	return &Server{engine: engine}
}

func (s *Server) GetEngineType(ctx context.Context, query *rpc.Query) (engineType *rpc.EngineType, err error) {
	bs, err := s.engine.svc.Get(query.ManufacturerModel)
	if err != nil {
		return nil, err
	}
	return bytesToEngineType(bs[0])
}

func bytesToEngineType(b []byte) (*rpc.EngineType, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &rpc.EngineType{
		ManufacturerModelCode: record.ManufacturerModelCode,
		ManufacturerName:      record.ManufacturerName,
		ModelName:             record.ModelName,
		EngineType:            record.EngineType.Description,
		Horsepower:            record.Horsepower,
		PoundsOfThrust:        record.PoundsOfThrust,
	}, nil
}
