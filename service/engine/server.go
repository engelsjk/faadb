package engine

import (
	"context"

	pb "github.com/engelsjk/faadb/rpc/engine"
)

type Server struct {
	engine *Engine
}

func NewServer(engine *Engine) *Server {
	return &Server{engine: engine}
}

func (s *Server) GetEngineType(ctx context.Context, query *pb.Query) (engineType *pb.EngineType, err error) {
	bs, err := s.engine.svc.Get(query.ManufacturerModel)
	if err != nil {
		return nil, err
	}
	return bytesToEngineType(bs[0])
}

func bytesToEngineType(b []byte) (*pb.EngineType, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &pb.EngineType{
		ManufacturerModelCode: record.ManufacturerModelCode,
		ManufacturerName:      record.ManufacturerName,
		ModelName:             record.ModelName,
		EngineType:            record.EngineType.Description,
		Horsepower:            record.Horsepower,
		PoundsOfThrust:        record.PoundsOfThrust,
	}, nil
}
