package reserved

import (
	"context"

	pb "github.com/engelsjk/faadb/rpc/reserved"
)

type Server struct {
	reserved *Reserved
}

func NewServer(reserved *Reserved) *Server {
	return &Server{reserved: reserved}
}

func (s *Server) GetAircraft(ctx context.Context, query *pb.Query) (*pb.Aircraft, error) {
	bs, err := s.reserved.svc.Get(query.NNumber)
	if err != nil {
		return nil, err
	}
	return bytesToAircraft(bs[0])
}

func (s *Server) GetMultipleAircraftByRegistrant(ctx context.Context, query *pb.Query) (*pb.MultipleAircraft, error) {
	bs, err := s.reserved.svc.List("registrant_name", query.RegistrantName, "registrant.name", true)
	if err != nil {
		return nil, err
	}
	return bytesToMultipleAircraft(bs)
}

func bytesToAircraft(b []byte) (*pb.Aircraft, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &pb.Aircraft{
		NNumber:              record.NNumber,
		RegistrantName:       record.Registrant.Name,
		RegistrantStreet1:    record.Registrant.Street1,
		RegistrantStreet2:    record.Registrant.Street2,
		RegistrantCity:       record.Registrant.City,
		RegistrantState:      record.Registrant.State,
		RegistrantZipcode:    record.Registrant.ZipCode,
		ReserveDate:          record.ReserveDate,
		ReservationType:      record.ReservationType.Description,
		ExpirationNoticeDate: record.ExpirationNoticeDate,
		PurgeDate:            record.PurgeDate,
	}, nil
}

func bytesToMultipleAircraft(bs [][]byte) (*pb.MultipleAircraft, error) {
	aircraft := make([]*pb.Aircraft, len(bs))
	for i, b := range bs {
		a, err := bytesToAircraft(b)
		if err != nil {
			return nil, err
		}
		aircraft[i] = a
	}
	return &pb.MultipleAircraft{Aircraft: aircraft}, nil
}
