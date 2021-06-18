package reserved

import (
	"context"

	"github.com/engelsjk/faadb/services/reserved/rpc"
)

type Server struct {
	reserved *ReservedService
}

func NewServer(reserved *ReservedService) *Server {
	return &Server{reserved: reserved}
}

func (s *Server) GetAircraft(ctx context.Context, query *rpc.Query) (*rpc.Aircraft, error) {
	var (
		bs  [][]byte
		err error
	)
	if query.NNumber != "" {
		nnumber := query.NNumber
		exact := true
		r := []rune(query.NNumber)
		if string(r[0]) == "*" {
			nnumber = string(r[1:])
			exact = false
		}
		bs, err = s.reserved.svc.List("nnumber", nnumber, "nnumber", exact)
	}
	if query.RegistrantName != "" {
		bs, err = s.reserved.svc.List("registrant_name", query.RegistrantName, "registrant.name", true)
	}
	if err != nil {
		return nil, err
	}
	return bytesToAircraft(bs)
}

func bytesToA(b []byte) (*rpc.A, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &rpc.A{
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

func bytesToAircraft(bs [][]byte) (*rpc.Aircraft, error) {
	as := make([]*rpc.A, len(bs))
	for i, b := range bs {
		a, err := bytesToA(b)
		if err != nil {
			return nil, err
		}
		as[i] = a
	}
	return &rpc.Aircraft{A: as}, nil
}
