package reservedserver

import (
	"context"

	"github.com/engelsjk/faadb/rpc/reserved"
)

type Server struct {
	reserved *ReservedService
}

func NewServer(reserved *ReservedService) *Server {
	return &Server{reserved: reserved}
}

func (s *Server) GetAircraft(ctx context.Context, query *reserved.Query) (*reserved.Aircraft, error) {
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
	if query.SerialNumber != "" {
		bs, err = s.reserved.svc.List("serial_number", query.SerialNumber, "serial_number", true)
	}
	if query.RegistrantName != "" {
		bs, err = s.reserved.svc.List("registrant_name", query.RegistrantName, "registrant.name", true)
	}
	if query.RegistrantStreet1 != "" {
		bs, err = s.reserved.svc.List("registrant_street_1", query.RegistrantName, "registrant.street_1", true)
	}
	if query.RegistrantState != "" {
		bs, err = s.reserved.svc.List("registrant_state", query.RegistrantState, "registrant.state", true)
	}
	if err != nil {
		return nil, err
	}
	return bytesToAircraft(bs)
}

func bytesToA(b []byte) (*reserved.A, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &reserved.A{
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

func bytesToAircraft(bs [][]byte) (*reserved.Aircraft, error) {
	as := make([]*reserved.A, len(bs))
	for i, b := range bs {
		a, err := bytesToA(b)
		if err != nil {
			return nil, err
		}
		as[i] = a
	}
	return &reserved.Aircraft{A: as}, nil
}
